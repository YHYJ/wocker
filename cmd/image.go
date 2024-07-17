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
	Short: "Manage docker images",
	Long:  `Specify or interactively manage docker images.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 解析参数
		listFlag, _ := cmd.Flags().GetBool("list")
		saveFlag, _ := cmd.Flags().GetBool("save")
		loadFlag, _ := cmd.Flags().GetBool("load")

		if listFlag {
			cli.ListImage()
		}

		if saveFlag {
			cli.SaveImages(args...)
		}

		if loadFlag {
			cli.LoadImage(args...)
		}
	},
}

func init() {
	imageCmd.Flags().Bool("list", false, "List all local images")
	imageCmd.Flags().Bool("save", false, "Save one or more images with TAG and ID to a tar archive, for example: '--save image1 image2:tag' or '--save all'")
	imageCmd.Flags().Bool("load", false, "Load an image from a tar archive, for example: '--load image1_archive image2_archive'")

	imageCmd.Flags().BoolP("help", "h", false, "help for image command")
	rootCmd.AddCommand(imageCmd)
}
