package models

import (
	"color/utils"
	"gorm.io/gorm"
)

type ImageInfo struct {
	gorm.Model
	C_type int    `json:"C_Type"` //颜色基调uint
	Image  string `json:"image"`
	M_type int    `json:"M_Type"` //测试类型
}

func (table *ImageInfo) TableName() string {
	return "ImageInfo_basic"
}
func AddImage(Image ImageInfo) {
	utils.DB.Create(&Image)
}
