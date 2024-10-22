package router

import (
	"fmt"
	"go-admin/entery"
	"go-admin/ui"

	"github.com/gin-gonic/gin"
)

func InitRouterIndex() {
	// 注册路由
	r := gin.Default()
	r.Use(cors())

	// 更改为发布模式
	gin.SetMode(gin.ReleaseMode)

	// =========================================
	// 静态资源
	r.Static("/r", "./default/dist")

	// 捕获未匹配的路由，返回 index.html 交给 Vue-router 处理
	r.NoRoute(func(c *gin.Context) {
		c.File("./default/dist/index.html")
	})

	// =========================================
	// 注册路由
	main := r.Group("/api/v50")
	{
		InitRouterGroupUser(main) // 注册用户模块路由
	}

	// 如果启用了 SSL
	if entery.JsonEntery.Https.Enabled {
		err := r.RunTLS(fmt.Sprintf(":%s", entery.JsonEntery.Https.Port), entery.JsonEntery.Https.CertFile, entery.JsonEntery.Https.KeyFile)
		if err != nil {
			ui.PrintError(err.Error())
		}
	} else {
		err := r.Run(fmt.Sprintf(":%s", entery.JsonEntery.Http.Port))
		if err != nil {
			ui.PrintError(err.Error())
		}
	}
}

// 解决跨域
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
