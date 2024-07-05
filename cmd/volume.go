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
	Short: "Backup or restore docker volumes",
	Long:  `Create a volume backup file with a timestamp, or use a backup file to restore the volume.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 解析参数
		listFlag, _ := cmd.Flags().GetBool("list")
		backupFlag, _ := cmd.Flags().GetBool("backup")
		restoreFlag, _ := cmd.Flags().GetBool("restore")

		if listFlag {
			cli.ListVolume()
		}

		if backupFlag {
			cli.BackupVolume(args)
		}

		if restoreFlag {
			cli.RestoreVolume(args[0])
		}
	},
}

func init() {
	volumeCmd.Flags().Bool("list", false, "List volumes")
	volumeCmd.Flags().Bool("backup", false, "Backup volumes, for example: '--backup volume1 volume2' or '--backup all'")
	volumeCmd.Flags().Bool("restore", false, "Restore volumes, for example: '--restore volume_backfile1 volume_backfile2'")

	volumeCmd.Flags().BoolP("help", "h", false, "help for volume command")
	rootCmd.AddCommand(volumeCmd)
}
