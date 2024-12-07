package e

var Msg = map[int]string{
	Success:      "操作成功",
	Error:        "操作失败",
	TokenExpired: "token验证失败",
}

func GetMsg(code int) string {
	msg, ok := Msg[code]
	if !ok {
		return Msg[Error]
	}
	return msg
}
