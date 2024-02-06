package service

import (
	"color/dto"
	"color/models"
	"color/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func Upload(c *gin.Context) {
	r := c.Request
	w := c.Writer
	imageInfo, url := upload(r, w, c)
	fmt.Println("url:", url)
	dto.AddImage(*imageInfo)
	utils.RespOk(w, url, "色块上传成功")
}

// todo 1.修改成只上传石田测试题
// todo 2.添加色感测试的生成器
func Uploadshi(c *gin.Context) {
	r := c.Request
	w := c.Writer
	imageInfo, url := upload(r, w, c)
	//models.AddImage(*imageInfo)
	Ishida := models.IssueIshida{}
	Ishida.IssueID = imageInfo.ID
	Ishida.Issue = *imageInfo
	Ishida.OptionA = c.Request.FormValue("optionA")
	Ishida.OptionB = c.Request.FormValue("optionB")
	Ishida.OptionC = c.Request.FormValue("optionC")
	Ishida.OptionD = c.Request.FormValue("optionD")
	Ishida.Key = c.Request.FormValue("key")
	models.AddIshida(Ishida)
	utils.RespOk(w, url, "石田测试题上传成功")
}
func upload(r *http.Request, w http.ResponseWriter, c *gin.Context) (imageInfo *dto.ImageInfo, url string) {
	//获得上传文件
	srcFile, head, err := r.FormFile("file")
	if err != nil {
		utils.RespFail(w, err.Error())
	}
	suffix := ".png"
	srcName := head.Filename
	t := strings.Split(srcName, ".")
	if len(t) > 1 {
		suffix = "." + t[len(t)-1]
	}
	fileName := fmt.Sprintf("%d%04d%s", time.Now().Unix(), rand.Int31(), suffix)
	imageInfo = &dto.ImageInfo{}
	imageInfo.C_type, err = strconv.Atoi(c.Request.FormValue("color"))
	if err != nil {
		utils.RespFail(w, err.Error())
	}
	imageInfo.M_type, err = strconv.Atoi(c.Request.FormValue("method"))
	if err != nil {
		utils.RespFail(w, err.Error())
	}
	if imageInfo.M_type == 1 {
		switch imageInfo.C_type {
		case 1:
			url = "./Asset/Upload/Color/Red/" + fileName
		case 2:
			url = "./Asset/Upload/Color/Orange/" + fileName
		case 3:
			url = "./Asset/Upload/Color/Yellow/" + fileName
		case 4:
			url = "./Asset/Upload/Color/Green/" + fileName
		case 5:
			url = "./Asset/Upload/Color/Blue/" + fileName
		case 6:
			url = "./Asset/Upload/Color/Purple/" + fileName
		}
	} else if imageInfo.M_type == 2 {
		url = "./Asset/Upload/Method1/" + fileName
	} else {
		utils.RespFail(w, "测试方法尚未录入")
	}
	fmt.Println(imageInfo)
	//创建新文件
	newFile, err := os.Create(url)
	if err != nil {
		utils.RespFail(w, err.Error())
	}
	//拷贝文件
	_, err = io.Copy(newFile, srcFile)
	if err != nil {
		utils.RespFail(w, err.Error())
	}
	imageInfo.Image = url
	return imageInfo, url
}
func History(c *gin.Context) {
	token := token(c)
	userJson, _ := utils.Red.Get(c, "user:"+token).Result()
	user := dto.UserInfo{}
	json.Unmarshal([]byte(userJson), &user)
	utils.RespOk(c.Writer, user.History, "历史记录")
}
