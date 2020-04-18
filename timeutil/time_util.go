package timeutil

import "fmt"

// AddTimeSuffix 增加时间后缀
// 适用于生成sql语句中的时间区间
func AddTimeSuffix(timeStr string) string {
	return fmt.Sprintf("%s 23:59:59", timeStr)
}
