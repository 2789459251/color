package models

import (
	"color/utils"
	"fmt"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
	"time"
)

type IssueIshida struct {
	gorm.Model
	IssueID uint
	Issue   ImageInfo `gorm:"foreignKey:IssueID"`
	OptionA string
	OptionB string
	OptionC string
	OptionD string
	Key     string
}

func (table *IssueIshida) TableName() string {
	return "IssueIshida_basic"
}
func AddIshida(Ishida IssueIshida) {
	utils.DB.Create(&Ishida)
}
func Select() []IssueIshida {
	issueIshida := []IssueIshida{}
	rand.Seed(uint64(time.Now().UnixNano()))
	utils.DB.Order("RAND()").Limit(4).Find(&issueIshida)
	for i := range issueIshida {
		utils.DB.Where("ID = ?", issueIshida[i].IssueID).First(&issueIshida[i].Issue)
		fmt.Println("图片信息：", issueIshida[i].Issue)
	}
	return issueIshida
}
