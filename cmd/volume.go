/*
File: volume.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-06-30 16:06:45

Description: 执行子命令 'volume'
*/

package cmd

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// volumeCmd represents the volume command
var volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "Backup or restore docker volumes",
	Long:  `Create a volume backup file with a timestamp, or use a backup file to restore the volume.`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Println("volume called")
	},
}

func init() {
	rootCmd.AddCommand(volumeCmd)
}
