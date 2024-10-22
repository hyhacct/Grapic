package userIssue

import (
	"go-admin/config"
	"go-admin/constant"
	"go-admin/model"
	"go-admin/tools"

	"github.com/gin-gonic/gin"
)

// 中间件预处理
func MiddlewaresLevel1() gin.HandlerFunc {
	return func(c *gin.Context) {

		var (
			jsonRequest   model.StructRequestUserIssue
			jsonModelData model.ModelData
			decodeText    string
			err           error
		)

		// 验证guid和评论等信息
		err = c.ShouldBindJSON(&jsonRequest)
		if err != nil {
			c.JSON(constant.CODE500, gin.H{"code": constant.CODE500, "msg": err.Error()})
			c.Abort()
			return
		}

		if tools.ToolsTextIsNull(jsonRequest.Text) ||
			tools.ToolsTextIsNull(jsonRequest.Email) ||
			tools.ToolsTextIsNull(jsonRequest.GUID) {
			c.JSON(constant.CODE403, gin.H{"code": constant.CODE403, "msg": "参数异常,请检查你的guid,邮箱等信息"})
			c.Abort()
			return
		}

		// 首先尝试一下解密评论内容和邮箱地址，如果无法正常base64解密，则直接返回错误信息(容错性考虑)
		if decodeText, err = tools.ToolsBase64Decode(jsonRequest.Text); err != nil {
			c.JSON(constant.CODE500, gin.H{"code": constant.CODE500, "msg": "评论内容异常,请检查你的评论内容"})
			c.Abort()
			return
		}
		if _, err = tools.ToolsBase64Decode(jsonRequest.Email); err != nil {
			c.JSON(constant.CODE500, gin.H{"code": constant.CODE500, "msg": "邮箱地址异常,请检查你的邮箱地址"})
			c.Abort()
			return
		}

		// 评论内容长度限制
		if len(decodeText) > 300 {
			c.JSON(constant.CODE403, gin.H{"code": constant.CODE403, "msg": "评论内容过长"})
			c.Abort()
			return
		}

		// 需要判断下guid是否存在
		if _, err = config.Engine.Table(new(model.ModelData)).Where("`guid`=?", jsonRequest.GUID).Get(&jsonModelData); err != nil {
			c.JSON(constant.CODE500, gin.H{"code": constant.CODE500, "msg": err.Error()})
			c.Abort()
			return
		}

		if tools.ToolsTextIsNull(jsonModelData.Guid) {
			c.JSON(constant.CODE403, gin.H{"code": constant.CODE403, "msg": "评论失败,文章不存在"})
			c.Abort()
			return
		}

		// 验证成功后放行
		c.Set("jsonRequest", jsonRequest)
		c.Set("jsonModelData", jsonModelData)
		c.Next()
	}
}
