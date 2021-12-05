package errmsg

const (
	Success = 200
	Error   = 500

	// jwt类错误
	AuthEmpty    = 1001
	InvalidToken = 1002

	// 用户登录注册类错误
	ErrPassword = 2001
	ErrUsername = 2002

	// 规则相关类错误
	ErrCreateRule      = 3001
	ErrCreateRuleState = 3002
	ErrRuleOffline     = 3003
	ErrOfflineRule     = 3004
	ErrOnlineRule      = 3005
)

var CodeMsg = map[int]string{
	Success:            "成功",
	AuthEmpty:          "请求头是空",
	Error:              "失败",
	InvalidToken:       "token非法",
	ErrPassword:        "登录密码错误",
	ErrUsername:        "用户名不能重复或者为admin",
	ErrCreateRule:      "创建新的规则错误",
	ErrCreateRuleState: "创建新的规则状态错误",
	ErrRuleOffline:     " 这条规则下线",
	ErrOfflineRule:     "规则下线错误",
	ErrOnlineRule:      "规则上线错误",
}
