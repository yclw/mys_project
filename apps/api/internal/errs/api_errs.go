package errs

import "github.com/yclw/mys_project/pkg/common/response"

// API服务错误码定义 (10101xxx)
var (
	// 请求相关错误
	RequestInvalid = response.NewError(10101001, "请求参数错误")

	// 服务调用相关错误
	ServiceUnavailable = response.NewError(10101101, "服务不可用")
)
