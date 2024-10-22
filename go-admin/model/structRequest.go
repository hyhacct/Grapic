package model

// StructRequestUserSubmit 客户端请求创建数据的结构体
type StructRequestUserSubmit struct {
	Type         string   `json:"type"`
	Once         bool     `json:"once"`
	ExceedTime   int64    `json:"exceed_time"`
	Password     string   `json:"password"`
	Ips          []string `json:"ips"`
	Context      string   `json:"context"`
	AllowComment bool     `json:"allow_comment"`
}

// StructRequestUserIssue 客户端请求创建评论的结构体
type StructRequestUserIssue struct {
	Email string `json:"email"` // 邮箱
	GUID  string `json:"guid"`  // 唯一标识符
	Text  string `json:"text"`  // 评论内容
}
