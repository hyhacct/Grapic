package userSubmit

import (
	"fmt"
	"go-admin/config"
	"go-admin/constant"
	"go-admin/model"
	"go-admin/tools"
	"time"

	"github.com/gin-gonic/gin"
)

func Controllers(g *gin.Context) {
	var (
		err        error
		json       model.StructRequestUserSubmit
		guid       string
		exceedTime int64
	)

	// 绑定json数据
	err = g.ShouldBindJSON(&json)
	if err != nil {
		g.JSON(constant.CODE500, gin.H{"code": constant.CODE500, "msg": err.Error()})
		return
	}

	// 生成一个随机的GUID标识
	guid, err = tools.ToolsRandomGuid()
	if err != nil {
		g.JSON(constant.CODE500, gin.H{"code": constant.CODE500, "msg": fmt.Sprintf("生成Guid失败: %s", err.Error())})
		return
	}

	// 设置超时时长(不为-1时)
	if json.ExceedTime != -1 {
		exceedTime = time.Now().Unix() + int64(json.ExceedTime)
	} else {
		exceedTime = -1
	}

	// 处理数据
	var result = model.ModelData{
		Content:      json.Context,
		Guid:         guid,
		Type:         json.Type,
		Password:     json.Password,
		Once:         json.Once,
		Ips:          json.Ips,
		ExpireTime:   exceedTime,
		UserIp:       g.ClientIP(),
		AllowComment: json.AllowComment,
	}

	_, err = config.Engine.Table(new(model.ModelData)).Insert(&result)
	if err != nil {
		g.JSON(constant.CODE500, gin.H{"code": constant.CODE500, "msg": fmt.Sprintf("增加条目失败: %s", err.Error())})
		return
	}

	g.JSON(constant.CODE200, gin.H{"code": constant.CODE200, "msg": "发送成功", "data": guid})
}
