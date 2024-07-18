/*
File: define_message.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-06-30 15:51:14

Description: 定义输出信息及其格式
*/

package general

var (
	ReferenceNotExistMessage = "Reference does not exist"      // 输出文本 - 引用不存在
	NoSuchImageMessage       = "No such image"                 // 输出文本 - 无此镜像
	NoSuchVolumeMessage      = "No such volume"                // 输出文本 - 无此存储卷
	NotVolumeArchiveMessage  = "Not a volume archive file"     // 输出文本 - 不是存储卷存档
	VolumeExistMessage       = "Volume already exists"         // 输出文本 - 存储卷已存在
	SpecifyMessage           = "Please specify the %s to %s\n" // 输出文本 - 请求指示
)
