package entity

var (
	OK = NewError(0, "ok")

	// server error
	ErrServer    = NewError(100001, "服务异常,请联系管理员")
	ErrParam     = NewError(100002, "参数有误")
	ErrSignParam = NewError(100003, "签名参数有误")
	ErrDebug     = NewError(100004, "Debug错误")
	// service module error
	ErrUserPhone   = NewError(201001, "用户手机号不合法")
	ErrUserCapture = NewError(201002, "用户验证码有误")
)
