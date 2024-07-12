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
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/docker/docker/api/types/image"
	"github.com/gookit/color"
	"github.com/yhyj/wocker/general"
)

const (
	idMinSearchLength = 4  // 用于查找 image ID 的字符串的最小长度
	idMinViewLength   = 12 // 用于显示 image ID 的字符串的最小长度
)

// ListImage 输出所有 image 的信息
func ListImage() {
	docker := general.DockerClient()

	// 获取 image 列表
	images, err := docker.ImageList(context.Background(), image.ListOptions{All: true})
	if err != nil {
		fileName, lineNo := general.GetCallerInfo()
		color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
		return
	}

	var (
		imageRepo    string
		imageTag     string
		imageID      string
		imageCreated string
		imageSize    string
	)

	tableHeader := []string{"Repository", "Tag", "ID", "Created", "Size"} // 表头
	tableData := [][]string{}                                             // 表数据
	rowData := []string{}                                                 // 行数据
	for _, image := range images {
		// 处理原始数据
		imageRepoTag := strings.Split(image.RepoTags[0], ":")
		id := strings.Split(image.ID, ":")
		OriginalSize, sizeUnit := general.Human(float64(image.Size), "B")
		// 列数据赋值
		imageRepo = imageRepoTag[0]
		imageTag = imageRepoTag[1]
		imageID = id[1][:idMinViewLength]
		imageCreated = general.UnixTime2TimeString(image.Created)
		imageSize = color.Sprintf("%6.1f %s", OriginalSize, sizeUnit)
		// 组装行数据
		rowData = []string{imageRepo, imageTag, imageID, imageCreated, imageSize}
		tableData = append(tableData, rowData)
	}

	dataTable := table.New()                                // 创建一个表格
	dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
	dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
	dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
		var style lipgloss.Style

		if row == 0 {
			return general.HeaderStyle // 第一行为表头
		}

		return style
	})

	dataTable.Headers(tableHeader...) // 设置表头
	dataTable.Rows(tableData...)      // 设置单元格

	color.Println(dataTable)
}

// SaveImages 将指定 images 保存到各自 tar 存档文件
//
// 参数：
//   - name: image 的 Repository 或 ID
func SaveImages(names ...string) {
	docker := general.DockerClient()

	// 获取 image 列表
	images, err := docker.ImageList(context.Background(), image.ListOptions{All: true})
	if err != nil {
		fileName, lineNo := general.GetCallerInfo()
		color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
		return
	}

	// 生成 map[Repository]ID
	imagesMap := make(map[string]string)
	for _, image := range images {
		imagesMap[image.RepoTags[0]] = strings.Split(image.ID, ":")[1]
	}

	// 如果 names 没有 Tag，以 'latest' 作为默认值
	for _, name := range names {
		if !strings.Contains(name, ":") {
			names = append(names, name+":latest")
		}
	}

	// 参数 name 允许是 image 的 Repository 或 ID，如果是 Repository，则获取其对应的 ID，如果为 'all'，则将所有 image 保存到各自 tar 存档文件
	if general.SliceContains(names, "all") {
		for imageRepo, imageID := range imagesMap {
			// 将 image 名中的 ':' 替换为 '_'，'/' 替换为 '-'，再与 ID 前 12 位以 '_' 拼接做为存储文件名
			filename := color.Sprintf("%s_%s", strings.Replace(strings.Replace(imageRepo, ":", "_", -1), "/", "-", -1), imageID[:idMinViewLength])

			// 保存 image
			err = general.SaveImage(docker, imageID, filename)
			if err != nil {
				fileName, lineNo := general.GetCallerInfo()
				color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
				return
			}

			// 输出信息
			color.Printf("- Save %s to %s\n", general.FgBlueText(imageRepo), general.FgLightBlueText(filename))
		}
	} else {
		for imageRepo, imageID := range imagesMap {
			if general.SliceContains(names, imageRepo) || general.FaintContains(names, imageID, idMinSearchLength) {
				// 将 image 名中的 ':' 替换为 '_'，'/' 替换为 '-'，再与 ID 前 12 位以 '_' 拼接做为存储文件名
				filename := color.Sprintf("%s_%s", strings.Replace(strings.Replace(imageRepo, ":", "_", -1), "/", "-", -1), imageID[:idMinViewLength])

				// 保存 image
				err = general.SaveImage(docker, imageID, filename)
				if err != nil {
					fileName, lineNo := general.GetCallerInfo()
					color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
					return
				}

				// 输出信息
				color.Printf("- Save %s to %s\n", general.FgBlueText(imageRepo), general.FgLightBlueText(filename))
			}
		}
	}
}

func LoadImage(file string) {}
