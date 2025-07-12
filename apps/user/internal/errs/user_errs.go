package errs

import "github.com/yclw/mys_project/pkg/common/response"

// 用户服务错误码定义 (10102xxx)
var (
	// 验证码相关错误
	CaptchaNotExist = response.NewError(10102002, "验证码不存在或者已过期")
	CaptchaError    = response.NewError(10102003, "验证码错误")
)
