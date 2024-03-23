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

var ishidas = make([]models.IssueIshida, 0)
var temp int

func Method1(c *gin.Context) {
	var issues = make([]dto.IssueInfo, 0)
	ishidas = models.Select()
	for _, ishida := range ishidas {
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
	utils.RespOk(c.Writer, ishidas[0], "已获取8道石田测试题")
}
func GetNextTest(c *gin.Context) {
	n := len(ishidas)
	temp++
	if temp >= n {
		utils.RespFail(c.Writer, "已到达最后一题")
		temp = n - 1
		return
	}
	i := temp + 1
	utils.RespOk(c.Writer, ishidas[temp], "获得第"+strconv.Itoa(i)+"题")
}
func GetLastTest(c *gin.Context) {
	temp--
	if temp < 0 {
		utils.RespFail(c.Writer, "已到达第一题")
		temp = 0
		return
	}
	i := temp + 1
	utils.RespOk(c.Writer, ishidas[temp], "获得第"+strconv.Itoa(i)+"题")
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
	//userJson, _ := utils.Red.Get(c, "user:"+token(c)).Result()
	//userInfo := dto.UserInfo{}
	//json.Unmarshal([]byte(userJson), &userInfo)
	//这里有数字
	user := models.FindUserById(strconv.Itoa(int(id.(uint64))))
	userInfo := dto.FindUserInfo(strconv.Itoa(user.UserInfoId)) //id不对把
	history := dto.History{
		Time:       time.Now(),
		Result:     "共有8道题，回答正确" + strconv.Itoa(cnt) + "道题;" + str,
		ResultInfo: results,
	}
	if userInfo.History == nil {
		userInfo.ID = uint(id.(uint64))
		userInfo.Hightest = 0
		userInfo.History = make([]dto.History, 0)
	}
	userInfo.History = append(userInfo.History, history)

	dto.RefreshUserInfo(userInfo)
	utils.RespOk(c.Writer, results, "共有8道题，回答正确"+strconv.Itoa(cnt)+"道题;"+str)
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
			return "分析：绿色色盲"
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
	id, _ := c.Get("userInfoId")
	userInfo := dto.FindUserInfo(strconv.Itoa(int(id.(uint64))))
	//if userInfo.History == nil {
	//	userInfo.ID = uint(id.(uint64))
	//	userInfo.Hightest = 0
	//	dto.RefreshUserInfo(userInfo)
	//}
	utils.RespOk(c.Writer, userInfo.Hightest, "获取最高纪录")
}
func SetHighest(c *gin.Context) {
	score, _ := strconv.Atoi(c.Request.FormValue("Score"))
	id, _ := c.Get("userInfoId")
	userInfo := dto.FindUserInfo(strconv.Itoa(int(id.(uint64))))
	userInfo.ID = uint(id.(uint64))
	if score <= userInfo.Hightest || score <= 0 {
		utils.RespFail(c.Writer, "未达到刷新要求")
		return
	}
	userInfo.Hightest = score
	dto.RefreshUserInfo(userInfo)
	utils.RespOk(c.Writer, nil, "刷新成功")
}
