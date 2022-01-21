package error

var (
	// 错误码4xxxx为业务逻辑错误
	BindingError          = Error{40001, "参数绑定失败"}
	EmailError            = Error{40002, "邮箱错误"}
	EmailExistError       = Error{40003, "邮箱已注册"}
	EmailNotFoundError    = Error{40004, "邮箱未注册"}
	PasswordNotMatchError = Error{40005, "账号密码不匹配"}
	UnLoginError          = Error{40006, "未登录"}
	CaptchaError          = Error{40006, "验证码不匹配"}
	// 错误码5xxxx为服务内部错误，msg不需要返回给前端，暂时不写
	DBError    = Error{50000, ""}
	RDError    = Error{50001, ""}
	TokenError = Error{50002, ""}
	AESError   = Error{50003, ""}
	HashError  = Error{50004, ""}
)
