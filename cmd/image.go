/*
File: image.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-06-30 16:06:45

Description: 执行子命令 'image'
*/

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yhyj/wocker/cli"
)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Backup or restore docker image",
	Long:  `Create an image backup file with TAG and ID, or use a backup file to restore the image.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 解析参数
		listFlag, _ := cmd.Flags().GetBool("list")
		backupFlag, _ := cmd.Flags().GetBool("backup")
		restoreFlag, _ := cmd.Flags().GetBool("restore")
		if listFlag {

			cli.ListImage()
		}

		if backupFlag {
			cli.BackupImage(args)
		}

		if restoreFlag {
			cli.RestoreImage(args[0])
		}
	},
}

func init() {
	imageCmd.Flags().Bool("list", false, "List images")
	imageCmd.Flags().Bool("backup", false, "Backup images, for example: '--backup image1 image2' or '--backup all'")
	imageCmd.Flags().Bool("restore", false, "Restore images, for example: '--restore image_backfile1 image_backfile2'")

	imageCmd.Flags().BoolP("help", "h", false, "help for image command")
	rootCmd.AddCommand(imageCmd)
}
