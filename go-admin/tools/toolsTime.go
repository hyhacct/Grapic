package tools

import (
	"fmt"
	"time"
)

// 时间戳转换到人类可读时间
func ToolsTimeStampToTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

// ToolsTimestampToDescribe 时间戳转换到描述性时间
func ToolsTimestampToDescribe(timestamp int64) string {

	t := time.Unix(timestamp, 0) // 将时间戳转换为time.Time类型
	now := time.Now()
	diff := now.Sub(t)

	switch {
	case diff < time.Minute:
		return "刚刚"
	case diff < time.Hour:
		return fmt.Sprintf("%d分钟前", int(diff.Minutes()))
	case diff < 24*time.Hour:
		return fmt.Sprintf("%d小时前", int(diff.Hours()))
	case diff < 7*24*time.Hour:
		return fmt.Sprintf("%d天前", int(diff.Hours()/24))
	default:
		return "超过一周前"
	}
}
