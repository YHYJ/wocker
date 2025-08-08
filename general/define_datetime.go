/*
File: define_datetime.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-07-09 09:42:49

Description: 处理日期/时间
*/

package general

import "time"

// UnixTime2TimeString Unix 时间戳转换为字符串格式
//
// 参数：
//   - timeStamp: Unix 时间戳
//
// 返回：
//   - 格式化的 Unix 时间戳字符串
func UnixTime2TimeString(unixTime int64) string {
	return time.Unix(unixTime, 0).Format("2006-01-02 15:04:05")
}

// GetCurrentTimestamp 获取当前时间戳
//
// 参数：
//   - format: 时间格式，例如 "20060102150405"（Go 参考时间）
//
// 返回：
//   - 当前时间戳
func GetCurrentTimestamp(format string) string {
	// 获取当前时间
	now := time.Now()

	// 格式化为 YYYYMMDDHHMMSS
	timestamp := now.Format(format)

	return timestamp
}
