package model

type StructEntery struct {
	Address string `json:"address"` // 服务地址(没有域名就填写IPV4地址)

	Http struct {
		Port string `json:"port"` // 服务端口
	} `json:"http"` // HTTP配置

	Https struct {
		Enabled  bool   `json:"enabled"`   // 是否启用SSL
		Port     string `json:"port"`      // SSL服务端口
		CertFile string `json:"cert_file"` // SSL证书路径
		KeyFile  string `json:"key_file"`  // SSL密钥路径
	} `json:"https"` // HTTPS配置
}
