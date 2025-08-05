/*
File: image.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-06-30 13:52:44

Description: 子命令 'image' 的实现
*/

package cli

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/gookit/color"
	"github.com/yhyj/wocker/general"
)

const (
	idMinSearchLength = 4  // 用于查找 image ID 的字符串的最小长度
	idMinViewLength   = 12 // 用于显示 image ID 的字符串的最小长度
)

// ListImages 输出所有 image 的信息
func ListImages() {
	// 获取 image 列表
	images, err := general.ListImages()
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
		originalSize, sizeUnit := general.Human(float64(image.Size), "B")
		// 列数据赋值
		imageRepo = imageRepoTag[0]
		imageTag = imageRepoTag[1]
		imageID = id[1][:idMinViewLength]
		imageCreated = general.UnixTime2TimeString(image.Created)
		imageSize = color.Sprintf("%6.1f %s", originalSize, sizeUnit)
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

// image 信息
type ImageInfo struct {
	Repo string
	Tag  string
	ID   string
}

// image 保存信息
type SaveInfo struct {
	Name string
	File string
}

// SaveImages 将指定 images 保存到各自存档文件
//
// 参数：
//   - names: image 的 Repository 或 ID，允许一次保存多个
func SaveImages(names []string) {
	if len(names) == 0 {
		color.Printf(general.DangerText(general.SpecifyMessage), "image", "save")
		return
	}

	// 获取 image 列表
	images, err := general.ListImages()
	if err != nil {
		fileName, lineNo := general.GetCallerInfo()
		color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
		return
	}

	// 存档文件名及其组成部分
	var (
		imageRepo        string
		imageTag         string
		imageID          string
		imageArchiveFile string
	)

	// 参数 names 允许是 image 的 Repository(:Tag), ID 或 'all'
	if general.SliceContains(names, "all") { // 参数中包含 'all'，将所有 image 保存到各自存档文件
		for _, image := range images {
			imageSplit := strings.Split(image.RepoTags[0], ":")
			imageRepo = imageSplit[0]                 // image Repository
			imageTag = imageSplit[1]                  // image Tag
			imageID = strings.Split(image.ID, ":")[1] // image ID without 'sha256' prefix
			// 将 image Repository 中的 '/' 替换为 '-'，再与 Tag 以及 ID 前 idMinViewLength 位以 '_' 拼接做为存储文件名
			imageArchiveFile = color.Sprintf("%s_%s_%s.dockerimage", strings.Replace(imageRepo, "/", "-", -1), imageTag, imageID[:idMinViewLength])

			// 保存 image
			err = general.SaveImage(imageID, imageArchiveFile)
			if err != nil {
				fileName, lineNo := general.GetCallerInfo()
				color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
				return
			}

			// 输出信息
			color.Printf("%s Save %s -> %s\n", general.PackFlag, general.FgBlueText(imageRepo), general.FgMagentaText(imageArchiveFile))
		}
	} else { // 参数为 image 的 Repository(:Tag) 或 ID
		var saveImages []SaveInfo // 需要保存的 image 信息切片

		for _, name := range names {
			var (
				imageInfo      ImageInfo   // 匹配成功的 image 信息
				matchingImages []ImageInfo // 匹配成功的 image 信息切片

				saveImage SaveInfo // 需要保存的 image 信息
			)

			nameSplit := strings.Split(name, ":")
			if len(nameSplit) == 2 { // name 是 image Repository:Tag，严格匹配 Repository 和 Tag 都符合的 image
				var status bool = false  // 是否匹配成功
				nameRepo := nameSplit[0] // 期望的 image Repository
				nameTag := nameSplit[1]  // 期望的 image Tag
				// 遍历 image 列表，查找与 name 对应的 image
				for _, image := range images {
					imageSplit := func() []string { // image Repository and Tag
						if len(image.RepoTags) == 0 {
							return []string{"", ""}
						}
						return strings.Split(image.RepoTags[0], ":")
					}()
					imageID = strings.Split(image.ID, ":")[1]                  // image ID without 'sha256' prefix
					if imageSplit[0] == nameRepo && imageSplit[1] == nameTag { // 匹配到一致的 Repository 和 Tag
						imageInfo.Repo = imageSplit[0] // image Repository
						imageInfo.Tag = imageSplit[1]  // image Tag
						imageInfo.ID = imageID
						matchingImages = append(matchingImages, imageInfo)
						status = true
					}
				}
				// 没有匹配到一致的 Repository 和 Tag
				if !status {
					color.Printf("%s Save %s -> %s\n", general.PackFlag, general.FgBlueText(name), general.DangerText(general.ReferenceNotExistMessage))
					continue
				}
			} else { // name 是 image ID 或 image Repository （这种情况可能因为 Tag 不同匹配到多个）两种情况
				// 遍历 image 列表，查找与 name 对应的 image
				for _, image := range images {
					imageSplit := func() []string { // image Repository and Tag
						if len(image.RepoTags) == 0 {
							return []string{"", ""}
						}
						return strings.Split(image.RepoTags[0], ":")
					}()
					imageID = strings.Split(image.ID, ":")[1]                      // image ID without 'sha256' prefix
					if imageSplit[0] == name || strings.HasPrefix(imageID, name) { // 匹配到一致的 Repository 或 ID
						imageInfo.Repo = imageSplit[0] // image Repository
						imageInfo.Tag = imageSplit[1]  // image Tag
						imageInfo.ID = imageID
						matchingImages = append(matchingImages, imageInfo)
					}
				}
			}

			if len(matchingImages) == 0 {
				// 没有匹配到 Repository
				color.Printf("%s Save %s -> %s\n", general.PackFlag, general.FgBlueText(name), general.DangerText(general.NoSuchImageMessage))
				continue
			} else {
				for _, image := range matchingImages {
					if image.Repo == "" {
						// 将 ID 前 idMinViewLength 位做为存储文件名
						imageArchiveFile = color.Sprintf("%s.dockerimage", image.ID[:idMinViewLength])
						saveImage.Name = image.ID[:idMinViewLength]
						saveImage.File = imageArchiveFile
					} else {
						// 将 image Repository 中的 '/' 替换为 '-'，再与 Tag 以及 ID 前 idMinViewLength 位以 '_' 拼接做为存储文件名
						imageArchiveFile = color.Sprintf("%s_%s_%s.dockerimage", strings.Replace(image.Repo, "/", "-", -1), image.Tag, image.ID[:idMinViewLength])
						saveImage.Name = color.Sprintf("%s:%s", image.Repo, image.Tag)
						saveImage.File = imageArchiveFile
					}
					saveImages = append(saveImages, saveImage)
				}
			}
		}

		// 保存 image
		for _, image := range saveImages {
			err = general.SaveImage(image.Name, image.File)
			if err != nil {
				fileName, lineNo := general.GetCallerInfo()
				color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
				return
			}
			// 输出信息
			color.Printf("%s Save %s -> %s\n", general.PackFlag, general.FgBlueText(image.Name), general.FgMagentaText(image.File))
		}
	}
}

// LoadImages 从存档文件加载 image
//
// 参数：
//   - files: 存档文件名，允许一次加载多个
func LoadImages(files []string) {
	if len(files) == 0 {
		color.Printf(general.DangerText(general.SpecifyMessage), "image archive file", "load")
		return
	}

	for _, file := range files {
		result, message, err := general.LoadImage(file)
		if err != nil {
			fileName, lineNo := general.GetCallerInfo()
			color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
			return
		}

		if result {
			for _, msg := range message {
				color.Printf("%s Load %s -> %s", general.LoadFlag, general.FgBlueText(file), general.FgMagentaText(strings.Split(msg, ": ")[1]))
			}
		} else {
			for _, msg := range message {
				color.Printf("%s Load %s -> %s\n", general.LoadFlag, general.FgBlueText(file), general.DangerText(msg))
			}
		}
	}
}
