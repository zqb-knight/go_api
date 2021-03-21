package e

var MsgMap = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	ERROR_DATABASE_INSERT_FAIL:     "数据库插入失败",
	ERROR_DATABASE_QUERY_FAIL:      "数据库查询失败",
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
