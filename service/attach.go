package service

import (
	"color/dto"
	"color/models"
	"color/utils"
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

// userJson, _ := utils.Red.Get(c, "user:"+token).Result()
// user := dto.UserInfo{}
// json.Unmarshal([]byte(userJson), &user)
func History(c *gin.Context) {
	_, userinfo := User(c)
	utils.RespOk(c.Writer, userinfo.History, "历史记录")
}

// 将存储的数据转成json数组
func UploadFavorite(c *gin.Context) {

	//favorite := dto.Favorite{
	//	Name: c.Request.FormValue("name"), //颜色昵称
	//}
	//c.ShouldBindJSON(favorite.R)
	//c.ShouldBindJSON(favorite.G)
	//c.ShouldBindJSON(favorite.B)
	//c.ShouldBindJSON(favorite.A)

	var favorite dto.Favorite
	if err := c.ShouldBindJSON(&favorite); err != nil {
		utils.RespFail(c.Writer, err.Error())
		return
	}
	if len(favorite.R) != len(favorite.G) || len(favorite.G) != len(favorite.B) || len(favorite.B) != len(favorite.A) {
		utils.RespFail(c.Writer, "数据格式错误")
		return
	}
	_, userinfo := User(c)
	//如果名字存在，重复命名
	if seachFavorite(userinfo.Favorite, favorite.Name) {
		utils.RespFail(c.Writer, "重复命名,请重新请求")
		return
	}
	userinfo.Favorite = append(userinfo.Favorite, favorite)
	dto.RefreshUserInfo(userinfo)
	utils.RespOk(c.Writer, userinfo.Favorite, "ok")
}
func Favorite(c *gin.Context) {
	_, userinfo := User(c)
	utils.RespOk(c.Writer, userinfo.Favorite, "收藏夹")
}
func CancelFavorite(c *gin.Context) {
	name := c.Request.FormValue("name")
	if name == "" {
		utils.RespFail(c.Writer, "颜色名不能为空")
		return
	}
	_, userInfo := User(c)
	favorite := userInfo.Favorite
	favorite, ok := deleteFavorite(favorite, name)
	if !ok {
		utils.RespFail(c.Writer, "删除失败,未查询到该色")
		return
	}
	userInfo.Favorite = favorite
	dto.RefreshUserInfo(userInfo)
	utils.RespOk(c.Writer, favorite, "删除成功")
}
