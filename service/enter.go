package service

import "gvb_server/service/image_ser"

type ServiceGroup struct {
	ImageService image_ser.ImageService
}

var ServiceApp = new(ServiceGroup)
