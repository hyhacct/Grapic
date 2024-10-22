package config

import (
	"fmt"
	"go-admin/constant"
	"go-admin/model"
	"go-admin/ui"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var Engine *xorm.Engine

// 初始化本地sqlite数据库
func InitSQLite() {
	var (
		err error
	)

	// 初始化数据库
	Engine, err = xorm.NewEngine("sqlite3", fmt.Sprintf("%s/default.db", constant.PATH_DIR_SQL))
	if err != nil {
		ui.PrintError("初始化数据库引擎失败" + err.Error())
		os.Exit(1)
	}

	// 设置名称映射
	Engine.SetMapper(names.GonicMapper{})

	// 同步表结构
	err = Engine.Sync2(new(model.ModelUser), new(model.ModelData), new(model.ModelIssue))
	if err != nil {
		ui.PrintError("同步表结构失败" + err.Error())
		os.Exit(1)
	}

	// done
	ui.PrintInfo("数据库已就绪")
}
