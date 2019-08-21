package e

const (
	SUCCESS               = 200
	ERROR                 = 500
	ERROR_INVALID_REQUEST = 100001
	ERROR_INVLIAD_PARA    = 100002
)

var MsgFlags = map[int]string{
	SUCCESS:               "请求完成",
	ERROR:                 "请求失败,服务器内部错误",
	ERROR_INVALID_REQUEST: "无效接口",
	ERROR_INVLIAD_PARA:    "无效的参数",
}

func GetMsg(code int) string {
	if msg, ok := MsgFlags[code]; ok {
		return msg
	}
	return MsgFlags[ERROR]
}
