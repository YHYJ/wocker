/*
File: define_types.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-07-09 11:09:00

Description: 针对特定数据类型的方法
*/

package general

// Contains 判断字符串切片中是否包含指定字符串
//
// 参数：
//   - slice: 字符串切片
//   - target: 指定字符串
//
// 返回：
//   - 布尔值
func Contains(slice []string, target string) bool {
	for _, str := range slice {
		if str == target {
			return true
		}
	}
	return false
}
