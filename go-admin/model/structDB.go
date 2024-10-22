package model

// 管理员信息表，管理员只有一个
type ModelUser struct {
	Username string // 用户名（默认admin）
	Password string // 密码（默认admin）
	Email    string // 邮箱
	Phone    string // 手机号
	Name     string // 昵称
	Avatar   string // 头像地址
}

// 提交数据保存结构
type ModelData struct {
	CreateTime   int64    `xorm:"created"`  // 条目创建时间
	Delete       int      `xorm:"deleted"`  // 删除标志
	Content      string   `xorm:"not null"` // 内容(经过了base64加密)
	Password     string   // 密码，可以为空(经过了hash加密)
	Guid         string   `xorm:"unique index"` // 唯一标识符，用于查询数据
	Type         string   `xorm:"not null"`     // 类型（text：文本，code：代码，logs：日志）
	Once         bool     // 是否只允许一次性使用（true：访问一次后就卸载，false：无限制访问）
	Ips          []string // 白名单IP列表，使用英文逗号分隔，可以为空
	ExpireTime   int64    // 过期时间，-1表示永不过期
	UserIp       string   // 发布者IP地址
	AllowComment bool     // 是否允许评论（true：允许，false：不允许）
}

// 游客评论数据表
type ModelIssue struct {
	Guid string `xorm:"index"` // 文章唯一标识符
	Ip   string // 评论者IP地址
	Name string // 评论者名称（一般用邮箱就行）
	Time int64  `xorm:"created"` // 评论时间
	Text string // 评论内容
}
