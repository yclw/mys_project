package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yclw/mys_project/apps/api/internal/handler"
)

// setupUserRoutes 设置用户路由
func setupUserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/user")
	{
		// ping
		user.GET("/ping", handler.User.Ping)

		// 用户注册
		user.POST("/register", handler.User.Register)

		// 发送验证码
		user.POST("/send-code", handler.User.SendVerificationCode)
	}
}
