package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service"
	"gvb_server/service/image_ser"
	"gvb_server/utils/jwts"
	"io/fs"
	"os"
)

// ImageUploadView 上传图片
// @Tags 图片管理
// @Summary 上传图片
// @Description 上传图片
// @Param token header string true "token"
// @Accept multipart/form-data
// @Param limit query string true "文件上传"
// @Router /api/images [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (receiver ImagesApi) ImageUploadView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	if claims.Role == 3 {
		res.FailWithMessage("游客不可上传图片", c)
		return
	}
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在的文件", c)
		return
	}

	// 判断路径是否存在
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}

	// 不存在就创建
	var resList []image_ser.FileUploadResponse

	for _, file := range fileList {

		ServiceRes := service.ServiceApp.ImageService.ImageUploadService(file)
		if !ServiceRes.IsSuccess {
			resList = append(resList, ServiceRes)
			continue
		}
		// 成功的
		if !global.Config.QiNiu.Enable {
			// 非七牛的本地保存
			err = c.SaveUploadedFile(file, ServiceRes.FileName)
			if err != nil {
				global.Log.Error(err)
				ServiceRes.Msg = err.Error()
				ServiceRes.IsSuccess = false
				resList = append(resList, ServiceRes)
				continue
			}
		}
		resList = append(resList, ServiceRes)
	}
	res.OkWithData(resList, c)

	//fileHeader, err := c.FormFile("image") //单文件
	//if err != nil {
	//	res.FailWithMessage(err.Error(), c)
	//	return
	//}

}
