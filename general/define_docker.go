/*
File: define_docker.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-06-30 13:39:57

Description: 创建 docker 客户端
*/

package general

import "github.com/docker/docker/client"

// DockerClient 创建 docker 客户端
//
// 返回：
//   - docker 客户端
//   - 错误信息
func DockerClient() (*client.Client, error) {
	docker, err := client.NewClientWithOpts(client.FromEnv)
	defer docker.Close()

	return docker, err
}
