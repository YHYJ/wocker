/*
File: define_docker.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-06-30 13:39:57

Description: 创建 docker 客户端
*/

package general

import (
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
