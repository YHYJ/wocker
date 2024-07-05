/*
File: volume.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-06-30 16:06:45

Description: 执行子命令 'volume'
*/

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yhyj/wocker/cli"
)

// volumeCmd represents the volume command
var volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "Manage docker volumes",
	Long:  `Specify or interactively manage daocker volumes.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 解析参数
		listFlag, _ := cmd.Flags().GetBool("list")
		saveFlag, _ := cmd.Flags().GetBool("save")
		loadFlag, _ := cmd.Flags().GetBool("load")

		if listFlag {
			cli.ListVolume()
		}

		if saveFlag {
			cli.SaveVolume(args)
		}

		if loadFlag {
			cli.LoadVolume(args[0])
		}
	},
}

func init() {
	volumeCmd.Flags().Bool("list", false, "List all volumes")
	volumeCmd.Flags().Bool("save", false, "Save one or more volumes with timestamp to a tar archive and compress it, for example: '--save volume1 volume2' or '--save all'")
	volumeCmd.Flags().Bool("load", false, "Load a volume from a tar archive, for example: '--load volume_backfile1 volume_backfile2'")

	volumeCmd.Flags().BoolP("help", "h", false, "help for volume command")
	rootCmd.AddCommand(volumeCmd)
}
