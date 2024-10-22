package tools

import (
	"fmt"

	"github.com/google/uuid"
)

// ToolsRandomGuid 生成随机GUID
func ToolsRandomGuid() (string, error) {
	u, err := uuid.NewRandom() // 使用强随机数生成UUID
	if err != nil {
		return "", err
	}

	// 截取UUID的一部分并格式化为 xxxx-xx-xxxxxxxxxx
	return fmt.Sprintf("%s-%s-%s", u.String()[0:4], u.String()[4:6], u.String()[6:16]), nil
}
