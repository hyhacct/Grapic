package tools

import (
	"io/ioutil"
	"os"
)

// ToolsCheckDirIsExist 目录是否存在，不存在返回false
func ToolsCheckDirIsExist(dir string) bool {
	_, err := os.Stat(dir)
	return err == nil
}

// ToolsCreateDir 创建目录，如果有误，返回错误信息
func ToolsCreateDir(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

// 本地文件是否存在
func ToolsIsFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

// ToolsReadFile 读取文件内容，如果有误，返回错误信息
func ToolsReadFile(filePath string) (string, error) {
	value, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(value), nil
}

// 写出文件内容，如果有误，返回错误信息
func ToolsWriteFile(filePath string, content string) error {
	return ioutil.WriteFile(filePath, []byte(content), 0644)
}
