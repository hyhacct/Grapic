package userNeedPassword

import (
	"go-admin/config"
	"go-admin/constant"
	"go-admin/model"
	"go-admin/tools"

	"github.com/gin-gonic/gin"
)

func Controllers(g *gin.Context) {

	var (
		err  error
		json model.ModelData
	)

	guid := g.Query("guid")

	// 如果guid为空，则返回参数错误
	if tools.ToolsTextIsNull(guid) {
		g.JSON(constant.CODE500, gin.H{"code": constant.CODE500, "msg": "参数错误,guid不能为空"})
		return
	}

	// 从数据库中取出对应数据
	_, err = config.Engine.Table(new(model.ModelData)).Where("`guid`=?", guid).Get(&json)
	if err != nil {
		g.JSON(constant.CODE500, gin.H{"code": constant.CODE500, "msg": err.Error()})
		return
	}

	// 如果数据库内没有这条数据，则报错
	if tools.ToolsTextIsNull(json.Guid) {
		g.JSON(constant.CODE500, gin.H{"code": constant.CODE500, "msg": "数据不存在"})
		return
	}

	// 是否需要密码(true:需要密码，false:不需要密码)
	if tools.ToolsTextIsNull(json.Password) {
		g.JSON(constant.CODE200, gin.H{"code": constant.CODE200, "data": false})
	} else {
		g.JSON(constant.CODE200, gin.H{"code": constant.CODE200, "data": true})
	}
}
