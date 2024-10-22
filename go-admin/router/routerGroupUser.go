package router

import (
	"go-admin/handle/userGetComments"
	"go-admin/handle/userGetContent"
	"go-admin/handle/userIssue"
	"go-admin/handle/userNeedPassword"
	"go-admin/handle/userRead"
	"go-admin/handle/userSubmit"

	"github.com/gin-gonic/gin"
)

// InitRouterGroupUser 初始化用户组路由
func InitRouterGroupUser(group *gin.RouterGroup) {

	chi := group.Group("user")
	{
		chi.POST("submit", userSubmit.MiddlewaresLevel1(), userSubmit.Controllers)                // 用户提交信息
		chi.GET("needPassword", userNeedPassword.Controllers)                                     // 是否需要提供密码
		chi.GET("read", userRead.MiddlewaresLevel1(), userRead.Controllers)                       // 用户读取信息
		chi.GET("get_comments", userGetComments.MiddlewaresLevel1(), userGetComments.Controllers) // 用户读取消息数据
		chi.GET("get_content", userGetComments.MiddlewaresLevel1(), userGetContent.Controllers)   // 用户读取消息数据
		chi.POST("issue", userIssue.MiddlewaresLevel1(), userIssue.Controllers)                   // 游客提交评论
	}
}
