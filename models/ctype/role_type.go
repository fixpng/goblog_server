package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin       Role = 1 // 管理员
	PermissionUser        Role = 2 // 用户
	PermissionVisitor     Role = 3 // 游客
	PermissionDisableUser Role = 4 // 黑名单
)

func (role Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(role.String())
}

func (role Role) String() string {
	switch role {
	case PermissionAdmin:
		return "管理员"
	case PermissionUser:
		return "用户"
	case PermissionVisitor:
		return "游客"
	case PermissionDisableUser:
		return "黑名单"
	default:
		return "其他"
	}
}
