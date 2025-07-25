package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yclw/mys_project/apps/api/internal/client"
	"github.com/yclw/mys_project/apps/api/internal/errs"
	"github.com/yclw/mys_project/pkg/common/response"
	v1 "github.com/yclw/mys_project/pkg/protobuf/gen/user/v1"
)

var User = &HandlerUser{}

type HandlerUser struct {
}

func (h *HandlerUser) Ping(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// 接收参数

	// 校验参数

	// 获取用户服务客户端

	// 调用grpc服务
	userClient := client.User()
	if userClient == nil {
		c.JSON(http.StatusOK, errs.ServiceUnavailable.ToResult())
		return
	}
	resp, err := userClient.Ping(ctx, &v1.PingRequest{})
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, response.OK(resp))
}

func (h *HandlerUser) Register(c *gin.Context) {
	_, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// 接收参数

	// 验证参数

	// 获取用户服务客户端

	// 调用grpc服务

	// 返回成功响应
}
func (h *HandlerUser) SendVerificationCode(c *gin.Context) {
	_, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// 接收参数

	// 验证参数

	// 获取用户服务客户端

	// 调用grpc服务

	// 返回成功响应
}
