package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yclw/mys_project/apps/api/internal/handler"
)

// setupAuthRoutes 设置认证路由
func setupAuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	{
		// ping
		auth.GET("/ping", handler.Auth.Ping)

		// 用户注册
		auth.POST("/register", handler.Auth.Register)

		// 发送验证码
		auth.POST("/send-code", handler.Auth.SendVerificationCode)
	}
}
