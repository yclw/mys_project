package response

// BusinessCode 业务状态码类型
type BusinessCode int

// 定义通用业务状态码常量
const (
	CodeSuccess BusinessCode = 200
)

// Result 统一API响应结构
type Result struct {
	Code BusinessCode `json:"code"`
	Msg  string       `json:"msg"`
	Data any          `json:"data"`
}

// Success 返回成功响应
func (r *Result) Success(data any) *Result {
	r.Code = CodeSuccess
	r.Msg = "success"
	r.Data = data
	return r
}

// Fail 返回失败响应
func (r *Result) Fail(code BusinessCode, msg string) *Result {
	r.Code = code
	r.Msg = msg
	r.Data = nil
	return r
}

// 便捷构造函数

// NewResult 创建新的Result实例
func NewResult() *Result {
	return &Result{}
}

// OK 快速创建成功响应
func OK(data any) *Result {
	return &Result{
		Code: CodeSuccess,
		Msg:  "success",
		Data: data,
	}
}

// Error 快速创建错误响应
func Error(code BusinessCode, msg string) *Result {
	return &Result{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
