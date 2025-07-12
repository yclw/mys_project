package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 全局中间件
	router.Use(gin.Logger()) // 日志中间件

	// API路由分组
	apiV1 := router.Group("/api/v1")
	{
		// 用户路由
		setupUserRoutes(apiV1)
	}

	return router
}
