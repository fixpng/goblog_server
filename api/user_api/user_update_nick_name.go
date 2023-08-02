package user_api

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/jwts"
	"strings"
)

type UserUpdateNicknameRequest struct {
	NickName string `json:"nick_name" structs:"nick_name"`
	Sign     string `json:"sign" structs:"sign"`
	Link     string `json:"link" structs:"link"`
	Avatar   string `json:"avatar" structs:"avatar"`
}

// UserUpdateNickName 用户修改当前登陆人信息
// @Tags 用户管理
// @Summary 用户修改当前登陆人信息
// @Description 用户修改当前登陆人信息
// @Router /api/user_info [put]
// @Param token header string true "token"
// @Param data body UserUpdateNicknameRequest true "昵称，签名，链接，头像"
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserUpdateNickName(c *gin.Context) {
	var cr UserUpdateNicknameRequest
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var newMaps = map[string]interface{}{}
	maps := structs.Map(cr)
	for key, v := range maps {
		if val, ok := v.(string); ok && strings.TrimSpace(val) != "" {
			newMaps[key] = val
		}
	}

	var userModel models.UserModel
	err = global.DB.Debug().Take(&userModel, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}

	avatar, isOk := newMaps["avatar"]
	s := avatar.(string)
	fmt.Println(len(s), isOk)

	err = global.DB.Model(&userModel).Updates(newMaps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改用户修改当前登陆人信息失败，若怀疑有bug请务必联系管理员", c)
		return
	}
	res.OkWithMessage("修改用户修改当前登陆人信息成功", c)

}
