package service

import (
	"color/models"
	"color/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
	"strconv"
	"strings"
	"time"
)

func Method1list(c *gin.Context) {
	//var issues = make([]models.IssueIshida, 4)
	Issue := models.Select()
	utils.RespOk(c.Writer, Issue, "返回四条石田测试题")
}
func GetColor(c *gin.Context) {
	Issue := models.SearchColor()
	Issue.Key = rands()
	token := token(c)
	utils.Red.Set(c, token, Issue.Key, 12*time.Hour)
	utils.RespOk(c.Writer, Issue, "返回两张相似色调色块")
}
func Judge_c(c *gin.Context) {
	token := token(c)
	key := c.Query("key")
	cacheKey := strings.Split(utils.Red.Get(c, token).String(), ":")
	if len(cacheKey) != 2 {
		utils.RespFail(c.Writer, "redis没有缓存该键")
		return
	}
	if key != strings.TrimSpace(cacheKey[1]) {
		utils.RespFail(c.Writer, "本题回答错误")
		return
	}
	utils.RespOk(c.Writer, nil, "本题回答正确")
	return
}
func rands() string {
	rand.Seed(uint64(time.Now().UnixNano()))
	return strconv.Itoa(int(rand.Int31n(2) + 1))
}
func token(c *gin.Context) string {
	authHeader, _ := c.Cookie("Authorization")
	if authHeader == "" {
		utils.RespFail(c.Writer, "没有token信息")
	}
	return authHeader
}
