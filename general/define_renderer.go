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
	CellStyle   = renderer.NewStyle().Align(lipgloss.Center).Padding(0, 1).Bold(false)                        // 单元格样式

	highlightColor    = lipgloss.AdaptiveColor{Light: TabLightColor, Dark: TabDarkColor}
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴") // 不活跃标签的边框
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└") // 活跃标签的边框
	docStyle          = lipgloss.NewStyle().Padding(0, 1, 0, 1)
	inactiveTabStyle  = lipgloss.NewStyle().Border(inactiveTabBorder, true).Padding(0, 2).BorderForeground(highlightColor)
	activeTabStyle    = inactiveTabStyle.Border(activeTabBorder, true).Padding(0, 8)
	windowStyle       = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).UnsetBorderTop().Padding(1, 0).Align(lipgloss.Center).BorderForeground(highlightColor)
)

// tabBorderWithBottom 返回指定样式的边框，用于构建活跃/不活跃标签
//
// 参数：
//   - left: 左边框
//   - middle: 中间框
//   - right: 右边框
//
// 返回：
//   - 指定样式的边框
func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}
