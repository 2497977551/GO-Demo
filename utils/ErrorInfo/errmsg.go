package ErrorInfo

const (
	SucCse = 200
	Error  = 500

	// code = 1000...用户模块错误
	ERRUserNameExists   = 1001
	ERRPassWordWrong    = 1002
	ERRUserNoExistent   = 1003
	ERRTokenNoExistent  = 1004
	ERRTokenOverdue     = 1005
	ERRTokenWrong       = 1006
	ERRTokenFormatWrong = 1007
	ERRNoPermission     = 1008
	// code = 2000...文章模块错误
	// code = 3000...分类模块错误
)

var codeMsg = map[int]interface{}{
	SucCse:              "OK",
	Error:               "FALSE",
	ERRUserNameExists:   "用户名已存在",
	ERRPassWordWrong:    "密码错误",
	ERRUserNoExistent:   "用户不存在",
	ERRTokenNoExistent:  "Token不存在",
	ERRTokenOverdue:     "Token已过期",
	ERRTokenWrong:       "Token错误",
	ERRTokenFormatWrong: "Token格式错误",
	ERRNoPermission:     "用户无权限",
}

func GetErrMsg(code int) interface{} {
	return codeMsg[code]
}
