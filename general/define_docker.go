/*
File: define_docker.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-06-30 13:39:57

Description: 与 docker 交互
*/

package general

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/client"
	"github.com/gookit/color"
)

// DockerClient 创建 docker 客户端
//
// 返回：
//   - docker 客户端
func DockerClient() *client.Client {
	docker, err := client.NewClientWithOpts(client.FromEnv)
	defer docker.Close()

	if err != nil {
		fileName, lineNo := GetCallerInfo()
		color.Printf("%s %s %s\n", DangerText(ErrorInfoFlag), SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
	}

	return docker
}

// SaveImage 将指定 image 保存到 tar 存档文件
//
// 参数：
//   - docker: docker 客户端
//   - imageID: image 的 ID
//   - filename: tar 存档文件名
//
// 返回：
//   - 错误信息
func SaveImage(docker *client.Client, imageID string, filename string) error {
	// 检索指定 image 为 io.ReadCloser
	reader, err := docker.ImageSave(context.Background(), []string{imageID})
	if err != nil {
		return err
	}
	defer reader.Close()

	// 创建文件
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将镜像数据写入文件
	_, err = io.Copy(file, reader)
	if err != nil {
		fileName, lineNo := GetCallerInfo()
		color.Printf("%s %s %s\n", DangerText(ErrorInfoFlag), SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
		return err
	}

	return nil
}
