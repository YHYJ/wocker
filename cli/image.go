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
		repository string
		tag        string
		id         string
		created    string
		size       string
	)

	tableHeader := []string{"Repository", "Tag", "ID", "Created", "Size"} // 表头
	tableData := [][]string{}                                             // 表数据
	rowData := []string{}                                                 // 行数据
	for _, image := range images {
		// 处理原始数据
		imageRepoTag := strings.Split(image.RepoTags[0], ":")
		imageID := strings.Split(image.ID, ":")
		OriginalSize, sizeUnit := general.Human(float64(image.Size), "B")
		// 列数据赋值
		repository = imageRepoTag[0]
		tag = imageRepoTag[1]
		id = imageID[1][:12]
		created = general.UnixTime2TimeString(image.Created)
		size = color.Sprintf("%6.1f %s", OriginalSize, sizeUnit)
		// 组装行数据
		rowData = []string{repository, tag, id, created, size}
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

func SaveImage(name []string) {}

func LoadImage(file string) {}
