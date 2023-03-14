package image_ser

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/plugins/qiniu"
	"gvb_server/utils"
	"io"
	"mime/multipart"
	"path"
	"strings"
	"time"
)

var (
	// WhiteImageList 图片上传白名单
	WhiteImageList = []string{"jpg", "png", "jpeg", "ico", "tiff", "gif", "svg", "webp"}
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 消息
}

// ImageUploadService 文件上传的方法
func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	fileName := file.Filename
	basePath := global.Config.Upload.Path
	filePath := path.Join(basePath, file.Filename)
	res.FileName = filePath

	// 判断上传文件后缀是否在白名单
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	if !utils.InList(suffix, WhiteImageList) {
		res.Msg = fmt.Sprintf("非法文件:%s", suffix)
		return
	}

	// 判断大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("图片大小超过设定大小，当前大小为：%.2fMB，设定大小为：%dMB", size, global.Config.Upload.Size)
		return
	}

	// 读取文件内容hash
	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err)
	}
	byteData, err := io.ReadAll(fileObj)
	imageHash := utils.MD5(byteData)
	// 去数据库中查这个图片hash是否存在
	var bannerModel models.BannerModel
	err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
	if err == nil {
		// 找到了
		res.FileName = bannerModel.Path
		res.Msg = "图片已存在"
		return
	}

	fileType := ctype.Local
	res.Msg = "图片上传成功"
	res.IsSuccess = true

	// 判断七牛云存储是否启用
	if global.Config.QiNiu.Enable {
		filePath, err = qiniu.UploadImage(byteData, fileName, global.Config.QiNiu.Prefix)
		if err != nil {
			global.Log.Error(err)
			res.Msg = err.Error()
			return
		}
		res.FileName = filePath
		res.IsSuccess = true
		res.Msg = "上传七牛成功"
		fileType = ctype.QiNiu
	} else {

	}

	// 图片入库
	global.DB.Create(&models.BannerModel{
		MODEL: models.MODEL{
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
		},
		Path:      filePath,
		Hash:      imageHash,
		Name:      fileName,
		ImageType: fileType,
	})
	return
}
