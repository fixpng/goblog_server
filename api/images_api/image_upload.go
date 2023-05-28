package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service"
	"gvb_server/service/image_ser"
	"io/fs"
	"os"
)

// ImageUploadView 上传图片
// @Tags 图片管理
// @Summary 上传图片
// @Description 上传图片
// @Param limit query string false "表示单个参数"
// @Router /api/images [post]
// @Produce json
// @Success 200 {object} res.Response{data=[]image_ser.FileUploadResponse}
func (receiver ImagesApi) ImageUploadView(c *gin.Context) {
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
