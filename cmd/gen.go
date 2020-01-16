package cmd

import (
	"fmt"
	"github.com/Smilefish0/gener/helpers"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

var genCmd = &cobra.Command{
	Use:     "gen",
	Aliases: []string{"g"},
	Short:   "生成models和proto",
	Long:    `生成models和proto`,

	Run: func(cmd *cobra.Command, args []string) {

		// 一些公共路径
		var goPath = os.Getenv("GOPATH")
		var goSrcPath = filepath.Join(goPath, "src")
		var pwdPath, _ = os.Getwd()
		var protoPath = filepath.Join(pwdPath, "protos")

		// 检查xo生成器命令
		dsn := helpers.GetDatabaseDSN()
		xoCmd := "xo"
		xoArgs := []string{dsn, "-o", "models", "--template-path", "templates"}
		xoExec := exec.Command(xoCmd, xoArgs...)
		_, err := xoExec.Output()
		if err != nil {
			fmt.Println("xo命令未找到，请先运行: go get -u github.com/xo/xo")
			//fmt.Fprintf(os.Stderr, "The command failed to perform: %s (Command: %s, Arguments: %s)", err, xoCmd, xoArgs)
			return
		}

		// 复制models到$GOPATH/src
		cpCmd := "cp"
		cpArgs := []string{"-r", filepath.Join(pwdPath, "models"), filepath.Join(goSrcPath, "models")}
		cpExec := exec.Command(cpCmd, cpArgs...)
		_, err = cpExec.Output()
		if err != nil {
			//fmt.Println("xo命令未找到，请先运行: go get -u github.com/xo/xo")
			fmt.Fprintf(os.Stderr, "The command failed to perform: %s (Command: %s, Arguments: %s)", err, cpCmd, cpArgs)
			return
		}

		// 切换目录到$GOPATH/src
		cmdStr := fmt.Sprintf("cd %s && proteus -p models -f models && cp models/models/* %s && rm -rf models", goSrcPath, protoPath)
		cmdMany := exec.Command("sh", "-c", cmdStr)
		if out, err := cmdMany.CombinedOutput(); err != nil {
			fmt.Errorf("Error: %v\n", err)
		} else {
			fmt.Printf("Success: %s\n%s\n", cmdStr, out)
		}
	},
}
