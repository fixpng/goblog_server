package user_ser

import (
	"gvb_server/service/redis_ser"
	"gvb_server/utils/jwts"
	"time"
)

type UserService struct {
}

func (UserService) Logout(claims *jwts.CustomClaims, token string) error {
	// 需要计算距离现在的过期时间
	exp := claims.ExpiresAt // 过期时间
	now := time.Now()
	diff := exp.Time.Sub(now)
	return redis_ser.Logout(token, diff)
}
