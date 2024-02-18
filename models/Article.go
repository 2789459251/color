package models

import (
	"color/utils"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Theme    int // 1 2 3 4
	Body     string
	Uploader int
}

func (article *Article) TableName() string {
	return "article_basic"
}
func GenerateArticle(article Article) *gorm.DB {
	return utils.DB.Create(&article)
}
func GetArticleByTheme(theme int) []Article {
	article := []Article{}
	utils.DB.Where("theme = ?", theme).Find(&article)
	return article
}
func FindArticle(theme int, body string) Article {
	article := Article{}
	utils.DB.Where("theme = ?&&body = ?", theme, body).First(&article)
	return article
}
