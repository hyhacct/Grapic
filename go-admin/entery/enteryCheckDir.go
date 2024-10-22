package entery

import (
	"go-admin/constant"
	"go-admin/tools"
	"go-admin/ui"
	"os"
)

// 程序启动时必须检查必要目录以及配置是否存在
func EnteryCheckDir() {
	var (
		err error
	)

	var data = []string{
		constant.PATH_DIR_SQL,
		constant.PATH_DIR_LOGS,
		constant.PATH_DIR_FILES,
		constant.PATH_DIR_CONFIGS,
	}

	for _, v := range data {
		if !tools.ToolsCheckDirIsExist(v) {
			err = tools.ToolsCreateDir(v)
			if err != nil {
				ui.PrintError("创建必须目录失败：" + v)
				os.Exit(1)
			}
			ui.PrintSuccess("创建目录成功：" + v)
		} else {
			ui.PrintInfo("目录已存在：" + v)
		}
	}
}
