/*
File: version.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-06-30 13:46:23

Description: 执行子命令 'version'
*/

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yhyj/wocker/cli"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print program version",
	Long:  `Print program version and exit.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 解析参数
		onlyFlag, _ := cmd.Flags().GetBool("only")

		// 打印版本信息
		cli.PrintVersionInfo(onlyFlag)
	},
}

func init() {
	versionCmd.Flags().Bool("only", false, "Only print the version number, like 'v0.0.1'")

	versionCmd.Flags().BoolP("help", "h", false, "help for version command")
	rootCmd.AddCommand(versionCmd)
}
