package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ    SignStatus = 1 // QQ
	SignGitee SignStatus = 2 // Gitee
	SignEmail SignStatus = 3 // 邮箱
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	switch s {
	case SignQQ:
		return "QQ"
	case SignGitee:
		return "Gitee"
	case SignEmail:
		return "邮箱"
	default:
		return "其他"
	}
}
