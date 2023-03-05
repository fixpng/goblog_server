package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin       Role = 1 // 管理员
	PermissionUser        Role = 2 // 普通登陆人
	PermissionVisitor     Role = 3 // 游客
	PermissionDisableUser Role = 4 // 被禁用的用户
)

func (role Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(role.String())
}

func (role Role) String() string {
	switch role {
	case PermissionAdmin:
		return "管理员"
	case PermissionUser:
		return "普通登陆人"
	case PermissionVisitor:
		return "游客"
	case PermissionDisableUser:
		return "被禁用的用户"
	default:
		return "其他"
	}
}
