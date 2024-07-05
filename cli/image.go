/*
File: image.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-06-30 13:52:44

Description: 子命令 'image' 的实现
*/

package cli

import (
	"context"

	"github.com/docker/docker/api/types/image"
	"github.com/gookit/color"
	"github.com/yhyj/wocker/general"
)

func ListImage() {
	docker := general.DockerClient()

	images, err := docker.ImageList(context.Background(), image.ListOptions{All: true})
	if err != nil {
		fileName, lineNo := general.GetCallerInfo()
		color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
		return
	}

	for _, image := range images {
		color.Printf("%s %s\n", image.ID[:10], image.RepoTags[0])
	}
}

func SaveImage(name []string) {}

func LoadImage(file string) {}
