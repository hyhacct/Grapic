package ui

import (
	"fmt"

	"github.com/gookit/color"
)

// PrintSuccess 输出成功色字体
func PrintSuccess(msg string) {
	outPut(msg, color.FgGreen.Render("Success"))
}

// PrintError 输出失败色字体
func PrintError(msg string) {
	outPut(msg, color.FgRed.Render("Error"))
}

// PrintWarning 输出警告色字体
func PrintWarning(msg string) {
	outPut(msg, color.FgYellow.Render("Warning"))
}

// PrintInfo 输出提示色字体
func PrintInfo(msg string) {
	outPut(msg, color.FgCyan.Render("Info"))
}

// 最终输出
func outPut(msg string, header string) {
	normal := color.FgDefault.Render

	color.Println(fmt.Sprintf("[%s] %s", header, normal(msg)))
}
