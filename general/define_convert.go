/*
File: define_convert.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 12:45:37

Description: 单位和格式数据转换
*/

package general

// Human 存储数据转换为人类可读的格式
//
// 参数：
//   - size: 需要转换的存储数据
//   - initialUnit: 初始单位
//
// 返回：
//   - 转换后的数据
//   - 转换后的单位
func Human(size float64, initialUnit string) (float64, string) {
	allUnits := [...]string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"}

	// 重构单位切片，从 initialUnit 开始
	var units []string
	for i, unit := range allUnits {
		if unit == initialUnit {
			units = allUnits[i:]
			break
		}
	}

	// 数据及转换
	for _, unit := range units {
		if size < 1000 {
			return size, unit
		}
		size /= 1000
	}

	return size, initialUnit
}
