package tools

import (
	"encoding/base64"
)

// ToolsBase64Decode 使用base64.StdEncoding解码
func ToolsBase64Decode(s string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}
