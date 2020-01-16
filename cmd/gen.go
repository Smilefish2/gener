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
	Short:   "生成model和proto",
	Long:    `使用多个组合命令生成model和proto文件并复制到当前目录models目录下`,

	Run: func(cmd *cobra.Command, args []string) {

		// 一些公共路径
		var goPath = os.Getenv("GOPATH")
		var goSrcPath = filepath.Join(goPath, "src")
		var pwdPath, _ = os.Getwd()
		var modelPath = filepath.Join(pwdPath, "models")
		var protoPath = filepath.Join(pwdPath, "protos")
		var templatePath = filepath.Join(goSrcPath, "github.com/Smilefish0/gener/templates")

		// 检查protoPath目录
		if !helpers.Exists(modelPath) {
			err := os.Mkdir(modelPath, os.ModePerm)
			if err != nil {
				fmt.Printf("mkdir failed![%v]\n", err)
				return
			}
		}

		// 检查protoPath目录
		if !helpers.Exists(protoPath) {
			err := os.Mkdir(protoPath, os.ModePerm)
			if err != nil {
				fmt.Printf("mkdir failed![%v]\n", err)
				return
			}
		}

		// 检查xo生成器命令
		dsn := helpers.GetDatabaseDSN()
		xoCmd := "xo"
		xoArgs := []string{"\"" + dsn + "\"", "-o", "models", "--template-path", templatePath}
		xoExec := exec.Command(xoCmd, xoArgs...)
		_, err := xoExec.Output()
		if err != nil {
			fmt.Fprintf(os.Stderr, "The command failed to perform: %s (Command: %s, Arguments: %s)", err, xoCmd, xoArgs)
			return
		}

		// 复制models到$GOPATH/src
		cpCmd := "cp"
		cpArgs := []string{"-r", modelPath, goSrcPath}
		cpExec := exec.Command(cpCmd, cpArgs...)
		_, err = cpExec.Output()
		if err != nil {
			fmt.Fprintf(os.Stderr, "The command failed to perform: %s (Command: %s, Arguments: %s)", err, cpCmd, cpArgs)
			return
		}

		// 切换目录到$GOPATH/src并执行一组命令
		cmdStr := fmt.Sprintf("cd %s && proteus -p models -f models && cp models/models/generated.proto %s && rm -rf models", goSrcPath, protoPath)
		cmdMany := exec.Command("sh", "-c", cmdStr)
		if _, err := cmdMany.CombinedOutput(); err != nil {
			fmt.Errorf("Error: %v\n", err)
		} else {
			fmt.Printf("Success: \nmodel生成文件所在目录：%s\nproto生成文件所在路径:%s\n", modelPath, filepath.Join(protoPath, "generated.proto"))
		}
	},
}
