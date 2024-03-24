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

	var favorite dto.Favorite
	if err := c.ShouldBindJSON(&favorite); err != nil {
		utils.RespFail(c.Writer, err.Error())
		return
	}

	_, userinfo := User(c)
	//如果名字存在，重复命名
	//if _,ok := seachFavorite(userinfo.Favorite, name);!ok {
	//	utils.RespFail(c.Writer, "您没有名为"+name+"的收藏夹")
	//	return
	//}
	userinfo.Favorite = append(userinfo.Favorite, favorite)
	dto.RefreshUserInfo(userinfo)
	utils.RespOk(c.Writer, userinfo.Favorite, "ok")
}

func CreateFavorite(c *gin.Context) {
	//创建收藏夹
	name := c.Request.FormValue("name")
	var favorite dto.Favorite
	favorite.Name = name
	_, userinfo := User(c)
	//如果名字存在，重复命名
	if _, ok := seachFavorite(userinfo.Favorite, name); ok {
		utils.RespFail(c.Writer, "您已经有名为"+name+"的收藏夹")
		return
	}
	userinfo.Favorite = append(userinfo.Favorite, favorite)
	dto.RefreshUserInfo(userinfo)
	utils.RespOk(c.Writer, userinfo.Favorite, "ok")
}

func ChangeFavorite(c *gin.Context) {
	name := c.Request.FormValue("name")
	newname := c.Request.FormValue("newname")
	if newname == name {
		utils.RespFail(c.Writer, "新名字不能与原来的重复！")
		return
	}
	_, userinfo := User(c)
	//如果名字存在，重复命名
	if _, ok := seachFavorite(userinfo.Favorite, name); !ok {
		utils.RespFail(c.Writer, "您没有名为"+name+"的收藏夹")
		return
	}
	favorites, ok := changFavorite(userinfo.Favorite, name, newname)
	if !ok {
		utils.RespFail(c.Writer, "修改失败")
		return
	}
	userinfo.Favorite = favorites
	dto.RefreshUserInfo(userinfo)
	utils.RespOk(c.Writer, favorites, "修改成功")
}

func AppendColor(c *gin.Context) {
	//添加颜色
	name := c.Request.FormValue("name")
	id := c.Request.FormValue("id")
	_, userinfo := User(c)
	//如果名字存在，重复命名
	if _, ok := seachFavorite(userinfo.Favorite, name); !ok {
		utils.RespFail(c.Writer, "您没有名为"+name+"的收藏夹")
		return
	}
	id_, err := strconv.Atoi(id)
	if err != nil {
		utils.RespFail(c.Writer, "数据转换失败，检查id格式为int")
	}
	color := dto.Color{
		Id: id_,
		R:  c.Request.FormValue("R"),
		G:  c.Request.FormValue("G"),
		B:  c.Request.FormValue("B"),
		A:  c.Request.FormValue("A"),
	}
	favorites, ok := appendColor(userinfo.Favorite, name, color)
	if !ok {
		utils.RespFail(c.Writer, "添加失败")
		return
	}
	userinfo.Favorite = favorites
	dto.RefreshUserInfo(userinfo)
	utils.RespOk(c.Writer, favorites, "颜色添加成功")
	return
}
func DeleteColor(c *gin.Context) {
	//删除颜色
	name := c.Request.FormValue("name")
	id := c.Request.FormValue("id")
	_, userinfo := User(c)
	//如果名字存在，重复命名
	if _, ok := seachFavorite(userinfo.Favorite, name); !ok {
		utils.RespFail(c.Writer, "您没有名为"+name+"的收藏夹")
		return
	}
	id_, err := strconv.Atoi(id)
	if err != nil {
		utils.RespFail(c.Writer, "数据转换失败，检查id格式为int")
	}
	favorites, ok := deleteColor(userinfo.Favorite, name, id_)
	if !ok {
		utils.RespFail(c.Writer, "删除失败，该收藏夹没有id为"+id+"的颜色")
		return
	}
	userinfo.Favorite = favorites
	dto.RefreshUserInfo(userinfo)
	utils.RespOk(c.Writer, favorites, "颜色删除成功")
	return
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
