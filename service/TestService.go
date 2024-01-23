package service

import (
	"color/models"
	"color/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
	"time"
)

func Method1list(c *gin.Context) {
	//var issues = make([]models.IssueIshida, 4)
	Issue := models.Select()
	utils.RespOk(c.Writer, Issue, "返回四条石田测试题")
}
func GetColor(c *gin.Context) {
	Issue := models.SearchColor()
	Issue.Key = string(rands())
	utils.RespOk(c.Writer, Issue, "返回两张相似色调色块")
}
func rands() int32 {
	rand.Seed(uint64(time.Now().UnixNano()))
	return rand.Int31n(2) + 1
}
