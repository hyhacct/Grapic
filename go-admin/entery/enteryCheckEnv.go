package entery

import (
	"encoding/json"
	"go-admin/model"
	"go-admin/tools"
	"go-admin/ui"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

var (
	config     = "./default/config/config.json"
	JsonEntery model.StructEntery
)

// 检查环境必备条件
func EnteryCheckEnv() {

	// 配置文件如果不存在则需要用户输入
	if !tools.ToolsIsFileExist(config) {
		getUserInput()
	} else {
		// 否则自动从文件读取配置
		readConfig()
	}
}

// 获取用户输入
func getUserInput() {

	// 第一步，先询问用户是否需要启用SSL
	survey.AskOne(&survey.Select{Message: "是否需要启用SSL?", Options: []string{"true", "false"}, Default: "false"}, &JsonEntery.Https.Enabled)

	// 如果用户启用了SSL，则只需要填写SSL相关的参数即可
	if JsonEntery.Https.Enabled {
		survey.AskOne(&survey.Input{Message: "SSL端口", Default: "443"}, &JsonEntery.Https.Port)
		survey.AskOne(&survey.Input{Message: "SSL证书文件路径"}, &JsonEntery.Https.CertFile)
		survey.AskOne(&survey.Input{Message: "SSL密钥文件路径"}, &JsonEntery.Https.KeyFile)
	} else {
		// 否则只需要提供HTTP端口即可
		survey.AskOne(&survey.Input{Message: "HTTP端口", Default: "9099"}, &JsonEntery.Http.Port)
	}

	// 最后一步，询问用户提供域名或IPV4地址
	survey.AskOne(&survey.Input{Message: "请提供你的域名或IP地址", Default: "localhost"}, &JsonEntery.Address)

	// 序列化配置
	value, err := json.Marshal(JsonEntery)
	if err != nil {
		ui.PrintError("配置序列化失败: " + err.Error())
		os.Exit(1)
	}

	// 将序列化后的数据写出到本地
	err = tools.ToolsWriteFile(config, string(value))
	if err != nil {
		ui.PrintError("配置写出到本地失败: " + err.Error())
		os.Exit(1)
	}
}

// 读取配置文件
func readConfig() {
	value, err := tools.ToolsReadFile(config)
	if err != nil {
		ui.PrintError("配置文件读取失败: " + err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal([]byte(value), &JsonEntery)
	if err != nil {
		ui.PrintError("配置文件反序列化失败: " + err.Error())
		os.Exit(1)
	}
}
