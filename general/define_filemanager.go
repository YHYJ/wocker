/*
File: define_filemanager.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-06-30 16:04:26

Description: 文件管理

Notice:
	- 文件参数名为 filePath
	- 文件夹参数名为 folderPath
*/

package general

import (
	"bufio"
	"os"
	"path/filepath"
)

// ReadFile 依次读取文件每行内容
//
// 参数：
//   - filePath: 文件路径
//
// 返回：
//   - 指定行的内容
func ReadFile(filePath string) ([]string, error) {
	// 打开文件
	text, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer text.Close()

	// 创建一个 Scanner 对象
	scanner := bufio.NewScanner(text)

	// 存储读取到的每行内容的切片
	var lines []string

	// 逐行读取文件内容
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// 检查是否出现了读取错误
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// FileExist 判断文件是否存在
//
// 参数：
//   - filePath: 文件路径
//
// 返回：
//   - 文件存在返回 true，否则返回 false
func FileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

// EmptyFile 清空文件内容，文件不存在则创建
//
// 参数：
//   - filePath: 文件路径
//
// 返回：
//   - 错误信息
func EmptyFile(filePath string) error {
	// 打开文件，如果不存在则创建，文件权限为读写
	text, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer text.Close()

	// 清空文件内容
	if err := text.Truncate(0); err != nil {
		return err
	}
	return nil
}

// ListFolderFiles 列出指定路径下的所有文件
//
// 参数：
//   - folderPath: 文件夹路径
//
// 返回：
//   - 文件列表
//   - 错误信息
func ListFolderFiles(folderPath string) ([]string, error) {
	files := []string{}

	// 打开文件夹
	text, err := os.Open(folderPath)
	if err != nil {
		return files, err
	}
	defer text.Close()

	// 读取文件夹中的文件
	fileInfos, err := text.ReadDir(-1)
	if err != nil {
		return files, err
	}

	// 遍历文件夹中的文件
	for _, fileInfo := range fileInfos {
		// 判断是否为文件
		if !fileInfo.IsDir() {
			files = append(files, fileInfo.Name())
		}
	}

	return files, nil
}

// CreateFile 创建文件，包括其父目录
//
// 参数：
//   - file: 文件路径
//
// 返回：
//   - 错误信息
func CreateFile(filePath string) (*os.File, error) {
	if FileExist(filePath) {
		return nil, nil
	}
	// 创建父目录
	folderPath := filepath.Dir(filePath)
	if err := CreateFolder(folderPath); err != nil {
		return nil, err
	}
	// 创建文件
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// CreateFolder 创建文件夹
//
// 参数：
//   - folderPath: 文件夹路径
//
// 返回：
//   - 错误信息
func CreateFolder(folderPath string) error {
	if FileExist(folderPath) {
		return nil
	}
	return os.MkdirAll(folderPath, os.ModePerm)
}

// GoToFolder 进到指定文件夹
//
// 参数：
//   - folderPath: 文件夹路径
//
// 返回：
//   - 错误信息
func GoToFolder(folderPath string) error {
	return os.Chdir(folderPath)
}

// WriteFile 写入内容到文件，文件不存在则创建，不自动换行
//
// 参数：
//   - filePath: 文件路径
//   - content: 内容
//   - mode: 写入模式，追加('a', O_APPEND, 默认)或覆盖('t', O_TRUNC)
//
// 返回：
//   - 错误信息
func WriteFile(filePath, content, mode string) error {
	// 确定写入模式
	writeMode := os.O_WRONLY | os.O_CREATE | os.O_APPEND
	if mode == "t" {
		writeMode = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	}

	// 将内容写入文件
	file, err := os.OpenFile(filePath, writeMode, 0666)
	if err != nil {
		return err
	}
	if _, err = file.WriteString(content); err != nil {
		return err
	}
	return nil
}

// WriteFileWithNewLine 写入内容到文件，文件不存在则创建，自动换行
//
// 参数：
//   - filePath: 文件路径
//   - content: 写入内容
//   - mode: 写入模式，追加('a', O_APPEND, 默认)或覆盖('t', O_TRUNC)
//
// 返回：
//   - 错误信息
func WriteFileWithNewLine(filePath, content, mode string) error {
	// 确定写入模式
	writeMode := os.O_WRONLY | os.O_CREATE | os.O_APPEND
	if mode == "t" {
		writeMode = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	}

	// 将内容写入文件
	file, err := os.OpenFile(filePath, writeMode, 0666)
	if err != nil {
		return err
	}
	if _, err = file.WriteString(content + "\n"); err != nil {
		return err
	}
	return nil
}

// DeleteFile 删除文件，如果目标是文件夹则包括其下所有文件
//
// 参数：
//   - filePath: 文件路径
//
// 返回：
//   - 错误信息
func DeleteFile(filePath string) error {
	if !FileExist(filePath) {
		return nil
	}
	return os.RemoveAll(filePath)
}
