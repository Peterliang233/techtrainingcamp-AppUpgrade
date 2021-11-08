package errmsg

const (
	Success = 200
	Error   = 500

	// jwt类错误
	AuthEmpty    = 1001
	InvalidToken = 1002

	// 登录类错误
	ErrPassword = 2001
)

var CodeMsg = map[int]string{
	Success:      "成功",
	AuthEmpty:    "请求头是空",
	Error:        "失败",
	InvalidToken: "token非法",
	ErrPassword:  "登录密码错误",
}
