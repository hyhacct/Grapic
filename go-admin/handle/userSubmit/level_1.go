package userSubmit

import (
	"github.com/gin-gonic/gin"
)

// 中间件预处理
func MiddlewaresLevel1() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 后期有任何要做的预处理都可以放在这里

		c.Next()
	}
}
