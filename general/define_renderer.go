/*
File: define_renderer.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-05-31 11:59:55

Description: 定义渲染器
*/

package general

import (
	"os"

	"github.com/charmbracelet/lipgloss"
)

var (
	renderer    = lipgloss.NewRenderer(os.Stdout)                                                             // 创建一个 lipgloss 渲染器
	HeaderStyle = renderer.NewStyle().Align(lipgloss.Center).Padding(0, 1).Bold(true).Foreground(HeaderColor) // 表头样式
	BorderStyle = renderer.NewStyle().Foreground(BorderColor)                                                 // 边框样式
)
