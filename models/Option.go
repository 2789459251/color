package models

import "gorm.io/gorm"

type Option struct {
	gorm.Model
	C_type int    `json:"C_Type"` //颜色基调uint
	Image  string `json:"image"`
	M_type int    `json:"M_Type"` //测试类型
}

func (table *Option) TableName() string {
	return "option_basic"
}
