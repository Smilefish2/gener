package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gener",
	Short: "数据库字段Go模型结构体和Proto原型文件生成器",
	Long:  `基于xo、proteus等开源项目实现的数据库字段Go模型结构体和Proto原型文件生成器，用于快速开发业务。`,
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {

	rootCmd.AddCommand(genCmd)
	rootCmd.AddCommand(doctorCmd)
}
