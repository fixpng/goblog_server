package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service/image_ser"
	"gvb_server/utils"
	"gvb_server/utils/jwts"
	"os"
	"path"
	"strings"
	"time"
)

// ImageUploadDataView 上传单个图片，返回图片url
// @Tags 图片管理
// @Summary 上传单个图片，返回图片url
// @Description 上传单个图片，返回图片url
// @Param token header string true "token"
// @Accept multipart/form-data
// @Param limit query string true "文件上传"
// @Router /api/image [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (receiver ImagesApi) ImageUploadDataView(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("参数校验失败", c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	if claims.Role == 3 {
		res.FailWithMessage("游客不可上传图片", c)
		return
	}

	fileName := file.Filename
	basePath := global.Config.Upload.Path

	// 判断上传文件后缀是否在白名单
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	if !utils.InList(suffix, image_ser.WhiteImageList) {
		res.FailWithMessage(fmt.Sprintf("非法文件:%s", suffix), c)
		return
	}

	// 判断大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		res.FailWithMessage(fmt.Sprintf(fmt.Sprintf("图片大小超过设定大小，当前大小为：%.2fMB，设定大小为：%dMB", size, global.Config.Upload.Size), suffix), c)
		return
	}

	//创建文件夹 /uploads/file/user_id
	dirList, err := os.ReadDir(basePath)
	if err != nil {
		res.FailWithMessage("文件目录不存在", c)
		return
	}
	if isInDirEntry(dirList, claims.NickName) {
		//没有创建，创建这个目录
		err := os.Mkdir(path.Join(basePath, claims.NickName), 077)
		if err != nil {
			global.Log.Error(err)
		}
	}
	// 1.如果存在重名，就加随机字符串 时间戳
	// 2.直接+时间戳
	now := time.Now().Format("20060102150405")
	fileName = nameList[0] + "_" + now + "." + suffix
	filePath := path.Join(basePath, claims.NickName, fileName)

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OkWithData("/"+filePath, c)
	return
}

func isInDirEntry(dirList []os.DirEntry, nickName string) bool {
	for _, entry := range dirList {
		if entry.Name() == nickName && entry.IsDir() {
			// 用户创建过
			return true
		}
	}
	return false
}
