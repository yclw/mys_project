package response

import "fmt"

// ErrorCode 错误码类型，与BusinessCode保持一致
type ErrorCode = BusinessCode

// BError 业务错误结构
type BError struct {
	Code ErrorCode `json:"code"`
	Msg  string    `json:"msg"`
}

func (e *BError) Error() string {
	return fmt.Sprintf("code:%v,msg:%s", e.Code, e.Msg)
}

// NewError 创建新的业务错误
func NewError(code ErrorCode, msg string) *BError {
	return &BError{
		Code: code,
		Msg:  msg,
	}
}

// ToResult 将错误转换为API响应
func (e *BError) ToResult() *Result {
	return &Result{
		Code: e.Code,
		Msg:  e.Msg,
		Data: nil,
	}
}
