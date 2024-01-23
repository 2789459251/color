package models

import (
	"color/utils"
	"gorm.io/gorm"
)

type IssueColor struct {
	gorm.Model
	Color  string
	ImageA string
	ImageB string
	Key    string
}

func (table *IssueColor) TableName() string {
	return "IssueColor_basic"
}
func AddColor(color IssueColor) {
	utils.DB.Create(&color)
}
func SearchColor() IssueColor {
	issueColor := IssueColor{}
	utils.DB.Order("RAND()").Limit(1).First(&issueColor)
	return issueColor
}
