/*
File: volume.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-06-30 13:52:44

Description: 子命令 'volume' 的实现
*/

package cli

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/gookit/color"
	"github.com/yhyj/wocker/general"
)

// ListVolumes 输出所有 volume 的信息
func ListVolumes() {
	// 获取 volume 列表
	volumes, err := general.ListVolumes()
	if err != nil {
		fileName, lineNo := general.GetCallerInfo()
		color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
		return
	}

	tableHeader := []string{"Name", "Driver", "Mountpoint"} // 表头
	tableData := [][]string{}                               // 表数据
	rowData := []string{}                                   // 行数据
	for _, volume := range volumes.Volumes {
		// 组装行数据
		rowData = []string{volume.Name, volume.Driver, volume.Mountpoint}
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

// SaveVolumes 将指定 volumes 保存到各自存档文件
//
// 参数：
//   - names: volume name，允许一次保存多个
func SaveVolumes(names []string) {
	if len(names) == 0 {
		color.Printf(general.DangerText(general.SpecifyMessage), "volume", "save")
		return
	}

	// 获取 volume 列表
	volumes, err := general.ListVolumes()
	if err != nil {
		fileName, lineNo := general.GetCallerInfo()
		color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
		return
	}

	// 获取所有 volume 名称
	var volumeNames []string
	for _, volume := range volumes.Volumes {
		volumeNames = append(volumeNames, volume.Name)
	}

	// 获取当前目录
	currentDir, err := os.Getwd()
	if err != nil {
		fileName, lineNo := general.GetCallerInfo()
		color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
		return
	}

	// 参数 names 允许是 volume 的 Name 或 'all'
	if general.SliceContains(names, "all") { // 参数中包含 'all'，将所有 volume 保存到各自存档文件
		for _, volumeName := range volumeNames {
			volumeArchiveFile := color.Sprintf("%s.tar.gz", volumeName)
			if err := general.SaveVolume(volumeName, currentDir, volumeArchiveFile); err != nil {
				fileName, lineNo := general.GetCallerInfo()
				color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
				return
			}
			// 输出信息
			color.Printf("- Save %s -> %s\n", general.FgBlueText(volumeName), general.FgMagentaText(volumeArchiveFile))
		}
	} else { // 参数为 volume 的 Name
		for _, name := range names {
			if !general.SliceContains(volumeNames, name) {
				color.Printf("- Save %s -> %s\n", general.FgBlueText(name), general.DangerText(general.NoSuchVolumeMessage))
				continue
			}
			volumeArchiveFile := color.Sprintf("%s.tar.gz", name)
			if err := general.SaveVolume(name, currentDir, volumeArchiveFile); err != nil {
				fileName, lineNo := general.GetCallerInfo()
				color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
				return
			}
			// 输出信息
			color.Printf("- Save %s -> %s\n", general.FgBlueText(name), general.FgMagentaText(volumeArchiveFile))
		}
	}
}

// LoadVolumes 从存档文件加载 volume
//
// 参数：
//   - files: 存档文件名，允许一次加载多个
func LoadVolumes(files []string) {
	if len(files) == 0 {
		color.Printf(general.DangerText(general.SpecifyMessage), "volume archive file", "load")
		return
	}
}
