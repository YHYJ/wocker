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

	"github.com/docker/docker/api/types/container"
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
//   - 功能与命令 `docker images` 一样，但显示效果不一样
//
// 返回：
//   - image 列表
//   - 错误信息
func ListImages() ([]image.Summary, error) {
	return docker.ImageList(ctx, image.ListOptions{All: true})
}

// ListVolumes 列出所有 volume
//
//   - 功能与命令 `docker volume ls` 一样，但显示效果不一样
//
// 返回：
//   - volume 列表
//   - 错误信息
func ListVolumes() (volume.ListResponse, error) {
	return docker.VolumeList(ctx, volume.ListOptions{})
}

// SaveImage 将指定 image 保存到存档文件
//
//   - 功能与命令 `docker save <imageName> -o <archiveFile>` 一样
//
// 参数：
//   - imageName: image 的 Repository(:Tag) 或 ID
//   - archiveFile: 存档文件
//
// 返回：
//   - 错误信息
func SaveImage(imageName string, archiveFile string) error {
	// 检索指定 image 为 io.ReadCloser
	reader, err := docker.ImageSave(ctx, []string{imageName})
	if err != nil {
		return err
	}
	defer reader.Close()

	// 创建文件
	file, err := ReCreateFile(archiveFile)
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

// SaveVolume 将指定 volume 保存到存档文件
//
//   - 功能与命令 `docker run --rm -v <volumeName>:/volume -v <filePath>:/backup busybox tar czf /backup/<archiveFile> -C /volume .` 一样
//
// 参数：
//   - volumeName: volume 名
//   - filePath: 存档文件路径
//   - archiveFile: 存档文件名
//
// 返回：
//   - 错误信息
func SaveVolume(volumeName string, filePath string, archiveFile string) error {
	const (
		volumePathInContainer = "/volume" // volume 在容器中的挂载路径
		backupPathInContainer = "/backup" // 备份文件夹在容器中的挂载路径
	)
	backupFileInContainer := color.Sprintf("%s/%s", backupPathInContainer, archiveFile)

	// 创建一个临时容器并挂载 volume
	containerConfig := &container.Config{
		// 基于 busybox 镜像创建容器
		Image: "busybox",
		// 使用 tar 打包 volume 中的文件
		Cmd: []string{"tar", "czf", backupFileInContainer, "-C", volumePathInContainer, "."},
	}
	hostConfig := &container.HostConfig{
		// 自动删除容器
		AutoRemove: true,
		// 设置挂载点
		Binds: []string{
			color.Sprintf("%s:%s", volumeName, volumePathInContainer),
			color.Sprintf("%s:%s", filePath, backupPathInContainer),
		},
	}

	// 创建容器，容器名称留空使其随机生成
	resp, err := docker.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "")
	if err != nil {
		return err
	}
	containerID := resp.ID

	// 启动容器
	if err := docker.ContainerStart(ctx, containerID, container.StartOptions{}); err != nil {
		return err
	}

	return nil
}

type LoadResponse struct {
	Stream string `json:"stream"`
	Error  string `json:"error"`
}

// LoadImage 从存档文件加载 image
//
//   - 功能与命令 `docker load -i <archiveFile>` 一样
//
// 参数：
//   - archiveFile: 存档文件
//
// 返回：
//   - docker service 是否返回错误信息
//   - docker service 的返回信息
//   - 错误信息
func LoadImage(archiveFile string) (bool, []string, error) {
	var (
		result  bool     = false
		message []string = make([]string, 0)
	)
	// 打开 tar 存档 文件
	file, err := os.Open(archiveFile)
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

// LoadVolume 从存档文件加载 volume
//
//   - 功能与命令 `docker run --rm -v <newVolumeName>:/volume -v <filePath>:/backup busybox tar xzf /backup/<archiveFile> -C /volume` 一样
//
// 参数：
//   - newVolumeName: 要创建的 volume 名
//   - filePath: 存档文件路径
//   - archiveFile: 存档文件名
//
// 返回：
//   - 错误信息
func LoadVolume(newVolumeName string, filePath string, archiveFile string) error {
	const (
		volumePathInContainer = "/volume" // volume 在容器中的挂载路径
		backupPathInContainer = "/backup" // 备份文件夹在容器中的挂载路径
	)
	backupFileInContainer := color.Sprintf("%s/%s", backupPathInContainer, archiveFile)

	// 创建一个临时容器并挂载 volume
	containerConfig := &container.Config{
		// 基于 busybox 镜像创建容器
		Image: "busybox",
		// 使用 tar 打包 volume 中的文件
		Cmd: []string{"tar", "xzf", backupFileInContainer, "-C", volumePathInContainer},
	}
	hostConfig := &container.HostConfig{
		// 自动删除容器
		AutoRemove: true,
		// 设置挂载点
		Binds: []string{
			color.Sprintf("%s:%s", newVolumeName, volumePathInContainer),
			color.Sprintf("%s:%s", filePath, backupPathInContainer),
		},
	}

	// 创建容器，容器名称留空使其随机生成
	resp, err := docker.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "")
	if err != nil {
		return err
	}
	containerID := resp.ID

	// 启动容器
	if err := docker.ContainerStart(ctx, containerID, container.StartOptions{}); err != nil {
		return err
	}

	return nil
}

// dockerClient 创建 docker 客户端
//
// 返回：
//   - docker 客户端
func dockerClient() *client.Client {
	docker, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		fileName, lineNo := GetCallerInfo()
		color.Printf("%s %s %s\n", DangerText(ErrorInfoFlag), SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
	}
	defer docker.Close()

	return docker
}
