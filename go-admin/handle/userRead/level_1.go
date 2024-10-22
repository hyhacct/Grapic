package userRead

import (
	"go-admin/constant"
	"strings"

	"github.com/gin-gonic/gin"
)

// 中间件预处理
func MiddlewaresLevel1() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 后期有任何要做的预处理都可以放在这里
		guid := c.Query("guid")
		guid = strings.ReplaceAll(guid, " ", "")

		// 如果客户端传入的guid为空，则返回错误信息
		if len(guid) == 0 {
			c.JSON(constant.CODE404, gin.H{"code": constant.CODE404, "msg": "guid为空, 无法匹配到数据"})
			c.Abort()
			return
		}

		c.Set("guid", guid)
		c.Next()
	}
}
