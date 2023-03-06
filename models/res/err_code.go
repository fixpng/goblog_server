package res

type ErrorCode int

const (
	SettingsError ErrorCode = 1001 //系统错误
	ArgumentError ErrorCode = 1002 //参数错误
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "系统错误",
		ArgumentError: "参数错误",
	}
)
