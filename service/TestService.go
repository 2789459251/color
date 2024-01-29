package service

import (
	"color/dto"
	"color/models"
	"color/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

func Method1list(c *gin.Context) {
	var issues = make([]dto.IssueInfo, 0)
	Issues := models.Select()
	for _, ishida := range Issues {
		issue := dto.IssueInfo{
			Id:  ishida.ID,
			Key: ishida.Key,
		}
		issues = append(issues, issue)
	}
	jsondata, _ := json.Marshal(issues)
	utils.Red.Set(c, token(c), jsondata, 12*time.Hour)
	utils.RespOk(c.Writer, Issues, "返回四条石田测试题")
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

// 将Issue的list存入redis并从redis去出查看答案
func Judge_m(c *gin.Context) {
	options := c.Request.FormValue("options")
	var str string
	Issues, _ := utils.Red.Get(c, token(c)).Result()
	var issuesCache = make([]dto.IssueInfo, 0)
	json.Unmarshal([]byte(Issues), &issuesCache)
	for i, issueCache := range issuesCache {
		if i < len(options) {
			if issueCache.Key == string(options[i]) {
				str += "第" + strconv.Itoa(i+1) + "题回答正确\n"
			} else {
				str += "第" + strconv.Itoa(i+1) + "题回答错误\n"
			}
		} else {
			// 如果 options 长度不足，则假定为错误
			str += "第" + strconv.Itoa(i+1) + "题回答错误\n"
		}
	}
	utils.RespOk(c.Writer, nil, str)
	return
}
