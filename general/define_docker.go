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
	"encoding/json"
	"io"
	"os"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"github.com/gookit/color"
)

var (
	ctx    = context.Background()
	docker = dockerClient()
)

// ListImages 列出所有 image
//
// 返回：
//   - image 列表
//   - 错误信息
func ListImages() ([]image.Summary, error) {
	return docker.ImageList(ctx, image.ListOptions{All: true})
}

// ListVolumes 列出所有 volume
//
// 返回：
//   - volume 列表
//   - 错误信息
func ListVolumes() (volume.ListResponse, error) {
	return docker.VolumeList(ctx, volume.ListOptions{})
}

// SaveImage 将指定 image 保存到 tar 存档文件
//
// 参数：
//   - imageID: image 的 ID
//   - filePath: tar 存档文件
//
// 返回：
//   - 错误信息
func SaveImage(imageID string, filePath string) error {
	// 检索指定 image 为 io.ReadCloser
	reader, err := docker.ImageSave(ctx, []string{imageID})
	if err != nil {
		return err
	}
	defer reader.Close()

	// 创建文件
	file, err := ReCreateFile(filePath)
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

type LoadResponse struct {
	Stream string `json:"stream"`
	Error  string `json:"error"`
}

// LoadImage 从 tar 存档文件加载 image
//
// 参数：
//   - filePath: tar 存档文件
//
// 返回：
//   - 错误信息
func LoadImage(filePath string) (bool, []string, error) {
	var (
		result  bool     = false
		message []string = make([]string, 0)
	)
	// 打开 tar 存档 文件
	file, err := os.Open(filePath)
	if err != nil {
		return result, message, err
	}
	defer file.Close()

	// 从 tar 存档文件加载 image
	response, err := docker.ImageLoad(ctx, file, true)
	if err != nil {
		return result, message, err
	}
	defer response.Body.Close()

	// 解析 docker service 返回的数据
	if response.Body != nil && response.JSON {
		decoder := json.NewDecoder(response.Body)
		for {
			var loadResponse LoadResponse
			if err := decoder.Decode(&loadResponse); err == io.EOF {
				break
			} else if err != nil {
				return result, message, err
			}

			if loadResponse.Error != "" {
				message = append(message, loadResponse.Error)
			} else {
				result = true
				message = append(message, loadResponse.Stream)
			}
		}
	}

	return result, message, nil
}

// dockerClient 创建 docker 客户端
//
// 返回：
//   - docker 客户端
func dockerClient() *client.Client {
	docker, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	defer docker.Close()

	if err != nil {
		fileName, lineNo := GetCallerInfo()
		color.Printf("%s %s %s\n", DangerText(ErrorInfoFlag), SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
	}

	return docker
}
