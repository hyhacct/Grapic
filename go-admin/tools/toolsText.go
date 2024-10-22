package tools

import "strings"

// ToolsTextIsNull 判断字符串是否为空, 空字符串返回true, 非空字符串返回false
func ToolsTextIsNull(str string) bool {
	return len(strings.ReplaceAll(str, " ", "")) == 0
}
