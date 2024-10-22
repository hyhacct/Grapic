package userRead

import (
	"go-admin/config"
	"go-admin/constant"
	"go-admin/model"

	"github.com/gin-gonic/gin"
)

func Controllers(g *gin.Context) {

	var (
		err  error
		json model.ModelData
	)

	guid, _ := g.Get("guid")

	// 从数据库中取出对应数据
	_, err = config.Engine.Table(new(model.ModelData)).Where("`guid`=?", guid).Get(&json)
	if err != nil {
		g.JSON(constant.CODE500, gin.H{"code": constant.CODE500, "msg": err.Error()})
		return
	}

	// 需要预处理一下数据后再返回给客户端

}
