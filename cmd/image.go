/*
File: image.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-06-30 16:06:45

Description: 执行子命令 'image'
*/

package cmd

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Backup or restore docker image",
	Long:  `Create an image backup file with TAG and ID, or use a backup file to restore the image.`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Println("image called")
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)
}
