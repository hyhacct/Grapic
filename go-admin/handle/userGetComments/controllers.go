package userGetComments

import (
	"fmt"
	"go-admin/config"
	"go-admin/constant"
	"go-admin/entery"
	"go-admin/model"
	"go-admin/tools"

	"github.com/gin-gonic/gin"
)

type modelUrl struct {
	Web string `json:"web"`
	API string `json:"api"`
}

type modelComment struct {
	Name string `json:"name"`
	Time string `json:"time"`
	Text string `json:"text"`
}

type modelMain struct {
	Type         string         `json:"type"`
	ExpireTime   string         `json:"expire_time"`
	IP           string         `json:"ip"`
	URL          modelUrl       `json:"url"`
	AllowComment bool           `json:"allow_comment"`
	Comments     []modelComment `json:"comments"`
}

func Controllers(g *gin.Context) {
	var (
		out_time    string
		comments    []model.ModelIssue
		err         error
		resp_coment []modelComment
		url_web     string
		url_api     string
	)

	json, _ := g.Get("json")
	ip, _ := g.Get("ip")

	// 过期时间
	if json.(model.ModelData).ExpireTime != -1 {
		out_time = tools.ToolsTimeStampToTime(json.(model.ModelData).ExpireTime)
	} else {
		out_time = "永久有效"
	}

	// 查询评论,以创建日期来排序
	err = config.Engine.Table(new(model.ModelIssue)).Where(`guid=?`, json.(model.ModelData).Guid).Desc("time").Find(&comments)
	if err != nil {
		g.JSON(constant.CODE500, gin.H{"code": constant.CODE500, "msg": err.Error()})
		return
	}

	// 迭代处理评论数据
	for _, v := range comments {
		resp_coment = append(resp_coment, modelComment{
			Name: v.Name,
			Time: tools.ToolsTimestampToDescribe(v.Time),
			Text: v.Text,
		})
	}

	// 拼接URL
	if entery.JsonEntery.Https.Enabled {
		if entery.JsonEntery.Https.Port == "443" {
			url_web = fmt.Sprintf("https://%s/view/read/%s", entery.JsonEntery.Address, json.(model.ModelData).Guid)
			url_api = fmt.Sprintf("https://%s/api/v50/user/get_content?guid=%s&password=%s", entery.JsonEntery.Address, json.(model.ModelData).Guid, json.(model.ModelData).Password)
		} else {
			url_web = fmt.Sprintf("https://%s:%s/view/read/%s", entery.JsonEntery.Address, entery.JsonEntery.Https.Port, json.(model.ModelData).Guid)
			url_api = fmt.Sprintf("https://%s:%s/api/v50/user/get_content?guid=%s&password=%s", entery.JsonEntery.Address, entery.JsonEntery.Https.Port, json.(model.ModelData).Guid, json.(model.ModelData).Password)
		}
	} else {
		url_web = fmt.Sprintf("http://%s:%s/view/read/%s", entery.JsonEntery.Address, entery.JsonEntery.Http.Port, json.(model.ModelData).Guid)
		url_api = fmt.Sprintf("http://%s:%s/api/v50/user/get_content?guid=%s&password=%s", entery.JsonEntery.Address, entery.JsonEntery.Http.Port, json.(model.ModelData).Guid, json.(model.ModelData).Password)
	}

	// 拼接组合数据
	data := modelMain{
		Type:         json.(model.ModelData).Type,
		ExpireTime:   out_time,
		IP:           ip.(string),
		AllowComment: json.(model.ModelData).AllowComment,
		URL: modelUrl{
			Web: url_web,
			API: url_api,
		},
		Comments: resp_coment,
	}

	g.JSON(constant.CODE200, gin.H{
		"code": constant.CODE200,
		"data": data,
	})
}
