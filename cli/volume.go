/*
File: volume.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-06-30 13:52:44

Description: 子命令 'volume' 的实现
*/

package cli

import (
	"context"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/docker/docker/api/types/volume"
	"github.com/gookit/color"
	"github.com/yhyj/wocker/general"
)

func ListVolume() {
	docker := general.DockerClient()

	// 获取 volume 列表
	volumes, err := docker.VolumeList(context.Background(), volume.ListOptions{})
	if err != nil {
		fileName, lineNo := general.GetCallerInfo()
		color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
		return
	}

	var (
		name       string
		driver     string
		mountpoint string
	)

	tableHeader := []string{"Name", "Driver", "Mountpoint"} // 表头
	tableData := [][]string{}                               // 表数据
	rowData := []string{}                                   // 行数据
	for _, volume := range volumes.Volumes {
		// 列数据赋值
		name = volume.Name
		driver = volume.Driver
		mountpoint = volume.Mountpoint
		// 组装行数据
		rowData = []string{name, driver, mountpoint}
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

func SaveVolume(name []string) {}

func LoadVolume(file string) {}
