/*
File: define_types.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-07-09 11:09:00

Description: 针对特定数据类型的方法
*/

package general

import "slices"

// SliceContains 判断字符串切片中是否包含指定字符串
//
// 参数：
//   - slice: 字符串切片
//   - target: 指定字符串
//
// 返回：
//   - 布尔值
func SliceContains(slice []string, target string) bool {
	if slices.Contains(slice, target) {
		return true
	}
	return false
}

// FaintContains 判断字符串切片中是否包含指定字符串
//
//   - 该方法为模糊匹配，即指定字符串前 length 位与切片中的字符串前 length 位匹配成功即可
//
// 参数：
//   - slice: 字符串切片
//   - target: 指定字符串
//   - length: 指定字符串前 length 位作为判断依据
//
// 返回：
//   - 布尔值
func FaintContains(slice []string, target string, length int) bool {
	if len(target) < length {
		return false
	}

	for _, str := range slice {
		strLength := len(str)
		if strLength >= length && strLength <= len(target) && str[:length] == target[:length] {
			return true
		}
	}
	return false
}
