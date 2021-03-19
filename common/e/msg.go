package e

var MsgMap = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "token已超时",
	ERROR_AUTH_TOKEN:               "token生成失败",
	ERROR_AUTH:                     "token错误",
}

func GetMsg(code int) string {
	msg, ok := MsgMap[code]
	if ok {
		return msg
	}
	return MsgMap[ERROR]
}
