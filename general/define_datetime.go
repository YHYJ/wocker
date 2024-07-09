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

// ParseDateTime 按照指定格式将字符串解析为日期/时间
//
// 参数：
//   - format: 解析时间日期的格式
//   - datetimeStr: 待解析的日期时间字符串
//
// 返回：
//   - time.Parse 解析结果
//   - 错误信息
func ParseDateTime(format, datetimeStr string) (time.Time, error) {
	// 解析时间字符串
	parsedTime, err := time.Parse(format, datetimeStr)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}
