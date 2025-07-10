package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// setupRouter 设置路由
func SetupRouter() *gin.Engine {
	router := gin.Default()
	// TODO: 设置路由
	router.GET("/", func(c *gin.Context) {
		// 返回hello world
		c.String(http.StatusOK, "hello world")
	})
	return router
}
