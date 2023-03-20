package desens

import "strings"

// DesensitizationEmail 脱敏邮箱
func DesensitizationEmail(email string) string {
	// Congee1997@outlook.com	C****@outlook.com
	// 897491068@qq.com			8****@qq.com
	eList := strings.Split(email, "@")
	if len(eList) != 2 {
		return ""
	}
	return eList[0][:1] + "****@" + eList[1]
}

// DesensitizationTel 脱敏手机号
func DesensitizationTel(tel string) string {
	// 18666666371 87666657
	// 18******371 87****57
	if len(tel) == 11 {
		return tel[:3] + "*****" + tel[8:]
	} else if len(tel) == 8 {
		return tel[:2] + "****" + tel[6:]
	}
	return ""
}
