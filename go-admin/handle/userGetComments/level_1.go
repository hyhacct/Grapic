package userGetComments

import (
	"go-admin/config"
	"go-admin/constant"
	"go-admin/model"
	"go-admin/tools"
	"time"

	"github.com/gin-gonic/gin"
)

// 中间件预处理
func MiddlewaresLevel1() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			ip       string
			guid     string
			password string
			json     model.ModelData
			err      error
		)

		guid = c.Query("guid")
		password = c.Query("password")
		ip = c.ClientIP()

		// 检查参数是否为空
		if tools.ToolsTextIsNull(guid) {
			c.JSON(constant.CODE403, gin.H{"code": constant.CODE403, "msg": "GUID不能为空"})
			c.Abort()
			return
		}

		// 检查guid是否有效
		_, err = config.Engine.Table(new(model.ModelData)).Where("`guid`=?", guid).Get(&json)
		if err != nil {
			c.JSON(constant.CODE500, gin.H{"code": constant.CODE500, "msg": err.Error()})
			c.Abort()
			return
		}

		// 检查GUID是否存在
		if tools.ToolsTextIsNull(json.Guid) {
			c.JSON(constant.CODE403, gin.H{"code": constant.CODE403, "msg": "当前GUID不存在"})
			c.Abort()
			return
		}

		// 检查密码是否正确
		if !tools.ToolsTextIsNull(json.Password) {
			if password != json.Password {
				c.JSON(constant.CODE403, gin.H{"code": constant.CODE403, "msg": "访问密码错误"})
				c.Abort()
				return
			}
		}

		// 如果白名单列表成员大于0表示启用了白名单
		if len(json.Ips) > 0 {

			allow := false

			for _, v := range json.Ips {
				if ip == v {
					allow = true
					break
				}
			}

			if !allow {
				c.JSON(constant.CODE403, gin.H{"code": constant.CODE403, "msg": "抱歉,您没有访问权限"})
				c.Abort()
				return
			}
		}

		// 如果到期时间不等于-1表示此篇文章有过期时间限制
		if json.ExpireTime != -1 {
			if time.Now().Unix() > json.ExpireTime {
				c.JSON(constant.CODE403, gin.H{"code": constant.CODE403, "msg": "此篇文章已过期"})
				c.Abort()
				return
			}
		}

		c.Set("json", json)
		c.Set("ip", ip)
		c.Set("guid", guid)
		c.Set("password", password)
		c.Next()
	}
}
