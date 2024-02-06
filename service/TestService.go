package service

import (
	"color/dto"
	"color/models"
	"color/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Method1(c *gin.Context) {
	var issues = make([]dto.IssueInfo, 0)
	Issues := models.Select()
	for _, ishida := range Issues {
		issue := dto.IssueInfo{
			Id:      ishida.ID,
			Key:     ishida.Key,
			ImageID: int(ishida.IssueID),
		}
		issues = append(issues, issue)
	}
	jsondata, _ := json.Marshal(issues)
	token := token(c)
	utils.Red.Set(c, "Result:"+token, jsondata, 12*time.Hour)
	utils.RespOk(c.Writer, Issues, "返回四条石田测试题")
}

// 将Issue的list存入redis并从redis去出查看答案
func Judge_m(c *gin.Context) {
	options := c.Request.FormValue("options")
	Issues, _ := utils.Red.Get(c, "Result:"+token(c)).Result()
	var str string
	var issuesCache = make([]dto.IssueInfo, 0)
	var results = make([]dto.ResultInfo, 0)
	var rets = make([]int, 3)
	json.Unmarshal([]byte(Issues), &issuesCache)
	cnt := 0
	for i, issueCache := range issuesCache {
		if i < len(options) {
			image := dto.SeachImage(issueCache.ImageID)
			if issueCache.Key == string(options[i]) {
				result := dto.ResultInfo{
					Key:   issueCache.Key,
					Mykey: string(options[i]),
					Point: Point(image.C_type),
					Image: image.Image,
					Flag:  true,
				}
				results = append(results, result)
				cnt++
			} else {
				result := dto.ResultInfo{
					Key:   issueCache.Key,
					Mykey: string(options[i]),
					Point: Point(image.C_type),
					Image: image.Image,
					Flag:  false,
				}
				rets[image.C_type-1]++
				results = append(results, result)
			}
		}
	}
	str = ret(rets)
	id, _ := c.Get("userInfoId")
	userJson, _ := utils.Red.Get(c, "user:"+token(c)).Result()
	user := dto.UserInfo{}
	json.Unmarshal([]byte(userJson), &user)
	history := dto.History{
		Time:       time.Now(),
		Result:     "共有4道题，回答正确" + strconv.Itoa(cnt) + "道题;" + str,
		ResultInfo: results,
	}
	if userJson == "" {
		user.ID = int(id.(uint64))
		user.Hightest = 0
	}
	user.History = append(user.History, history)
	toUserJson, _ := json.Marshal(user)
	utils.Red.Set(c, "user:"+token(c), toUserJson, -1)
	utils.RespOk(c.Writer, results, "共有4道题，回答正确"+strconv.Itoa(cnt)+"道题;"+str)
	return
}
func Point(t int) string {
	switch t {
	case 1:
		{
			return "分析：红色色盲"
		}
	case 2:
		{
			return "分析：红色色盲"
		}
	case 3:
		{
			return "分析：蓝紫色盲"
		}
	}
	return ""
}
func ret(t []int) string {
	var str string
	if t[0] != 0 {
		str += "有一定程度红色认知困难 "
	}
	if t[1] != 0 {
		str += "有一定程度绿色认知困难 "
	}
	if t[2] != 0 {
		str += "有一定程度蓝紫色认知困难 "
	}
	return str
}
func GetHighest(c *gin.Context) {
	token := token(c)
	id, _ := c.Get("userInfoId")
	userJson, _ := utils.Red.Get(c, "user:"+token).Result()
	user := dto.UserInfo{}
	json.Unmarshal([]byte(userJson), &user)
	if userJson == "" {
		user.ID = int(id.(uint64))
		user.Hightest = 0
		utils.Red.Set(c, "user:"+token, user, -1)
	}
	utils.RespOk(c.Writer, user.Hightest, "获取最高纪录")
}
func SetHighest(c *gin.Context) {
	score, _ := strconv.Atoi(c.Query("Score"))
	id, _ := c.Get("userInfoId")
	userJson, _ := utils.Red.Get(c, "user:"+token(c)).Result()
	user := dto.UserInfo{}
	json.Unmarshal([]byte(userJson), &user)
	if userJson == "" {
		user.ID = int(id.(uint64))
	}
	if score <= user.Hightest {
		utils.RespFail(c.Writer, "未达到刷新要求")
		return
	}
	user.Hightest = score
	toUserJson, _ := json.Marshal(user)
	utils.Red.Set(c, "user:"+token(c), toUserJson, -1)
	utils.RespOk(c.Writer, nil, "刷新成功")
}
