package userIssue

import (
	"go-admin/config"
	"go-admin/constant"
	"go-admin/model"

	"github.com/gin-gonic/gin"
)

func Controllers(g *gin.Context) {

	// c.Set("jsonRequest", jsonRequest)
	// c.Set("jsonModelData", jsonModelData)

	jsonRequest, _ := g.Get("jsonRequest")
	jsonModelData, _ := g.Get("jsonModelData")

	var (
		err error
	)

	var data = model.ModelIssue{
		Guid: jsonModelData.(model.ModelData).Guid,
		Ip:   g.ClientIP(),
		Name: jsonRequest.(model.StructRequestUserIssue).Email,
		Text: jsonRequest.(model.StructRequestUserIssue).Text,
	}

	_, err = config.Engine.Insert(&data)
	if err != nil {
		g.JSON(constant.CODE500, gin.H{"code": constant.CODE500, "msg": err.Error()})
		return
	}

	g.JSON(constant.CODE200, gin.H{"code": constant.CODE200, "msg": "评论成功"})
}
