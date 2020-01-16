package cmd

import (
	"github.com/Smilefish0/gener/helpers"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

var doctorCmd = &cobra.Command{
	Use:     "doctor",
	Aliases: []string{"d"},
	Short:   "检查命令运行环境",
	Long:    `检查运行环境是否满足需求`,

	Run: func(cmd *cobra.Command, args []string) {

		// check $GOPATH
		var goPath = os.Getenv("GOPATH")
		if goPath == "" {
			color.Red("GOPATH 环境变量未设置")
			return
		}
		color.Blue("$GOPATH: %s\n", goPath)

		// check $GOPATH/src
		var goSrcPath = filepath.Join(goPath, "src")
		if !helpers.Exists(goSrcPath) {
			color.Red("$GOPATH/src 不存在")
			return
		}
		color.Blue("$GOPATH/src: %s\n", goSrcPath)

		// 获取当前所在目录
		pwdPath, err := os.Getwd()
		if err != nil {
			color.Red("获取当前目录失败")
			return
		}
		color.Blue("pwdPath: %s\n", pwdPath)

		// 检查配置文件
		envFilePath := filepath.Join(pwdPath, ".env")
		if !helpers.Exists(envFilePath) {
			color.Red(".env文件未找到，请先配置.env文件中的数据库连接参数")
			return
		}
		color.Blue(".env File: %s\n", envFilePath)

		// 检查xo生成器命令
		xoBinPath, err := exec.LookPath("xo")
		if err != nil {
			color.Red("xo命令未找到，请先运行: go get -u -v github.com/xo/xo")
			return
		}
		color.Blue("xo: %s\n", xoBinPath)

		// 检查protoc生成器命令
		protocBinPath, err := exec.LookPath("protoc")
		if err != nil {
			color.Red("protoc命令未找到，请先运行: go get -u -v github.com/golang/protobuf/{proto,protoc-gen-go}")
			return
		}
		color.Blue("protoc: %s\n", protocBinPath)

		// 检查proteus生成器命令
		proteusBinPath, err := exec.LookPath("proteus")
		if err != nil {
			color.Red("proteus命令未找到，请先运行: go get -u -v gopkg.in/src-d/proteus.v1/cli/proteus")
			return
		}
		color.Blue("proteus: %s\n", proteusBinPath)

		color.Green("所有运行环境检测成功")
	},
}
