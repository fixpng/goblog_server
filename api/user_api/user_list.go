package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/common"
	"gvb_server/utils/desens"
	"gvb_server/utils/jwts"
)

type UserResponse struct {
	models.UserModel
	RoleID int `json:"role_id"`
}
type UserListRequest struct {
	models.PageInfo
	Role int `json:"role" form:"role"`
}

// UserListView 用户列表
// @Tags 用户管理
// @Summary 用户列表
// @Description 用户列表
// @Param data body UserListRequest    true  "查询参数"
// @Param token header string true "token"
// @Router /api/users [get]
// @Produce json
// @Success 200 {object} res.Response{data=UserResponse}
func (UserApi) UserListView(c *gin.Context) {
	// 如何判断是管理员
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var page UserListRequest
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var users []UserResponse
	list, count, _ := common.ComList(models.UserModel{Role: ctype.Role(page.Role)}, common.Option{
		PageInfo: page.PageInfo,
		Likes:    []string{"nick_name"},
	})

	for _, user := range list {
		//fmt.Println(claims.Role)
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			// 非管理员脱敏
			user.UserName = ""
			user.Tel = desens.DesensitizationTel(user.Tel)
			user.Email = desens.DesensitizationEmail(user.Email)
		}
		users = append(users, UserResponse{
			UserModel: user,
			RoleID:    int(user.Role),
		})
	}

	res.OkWithList(users, count, c)
}
