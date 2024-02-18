package service

import (
	"color/models"
	"color/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

var articleList []models.Article
var i int

func UploadArticle(c *gin.Context) {
	//token := token(c)
	//toke, _ := utils.ParseToken(token)
	//claims := toke.Claims.(*utils.AuthClaims)
	//fmt.Println(claims.UserInfoId)
	if c.Request.FormValue("Theme") == "" {
		utils.RespFail(c.Writer, "主题不能为空！")
		return
	}
	theme, _ := strconv.Atoi(c.Request.FormValue("Theme"))
	if theme < 1 || theme > 4 {
		utils.RespFail(c.Writer, "请输入有效的主题！")
		return
	}
	body := c.Request.FormValue("Body")
	if body == "" {
		utils.RespFail(c.Writer, "文章主体不能为空！")
		return
	}
	//id, err := c.Get("userInfoId")
	id, _ := c.Get("userInfoId")
	//if err {
	//	utils.RespFail(c.Writer, "获取上传用户信息错误")
	//	return
	//}
	//fmt.Println(id)
	article := models.Article{
		Model:    gorm.Model{},
		Theme:    theme,
		Body:     body,
		Uploader: int(id.(uint64)),
	}
	if models.FindArticle(theme, body).Body != "" {
		utils.RespOk(c.Writer, nil, "文章已存在")
		return
	}
	models.GenerateArticle(article)
	utils.RespOk(c.Writer, nil, "上传文章成功")
}
func GetArticles(c *gin.Context) {
	if c.Request.FormValue("Theme") == "" {
		utils.RespFail(c.Writer, "主题不能为空！")
		return
	}
	theme, _ := strconv.Atoi(c.Request.FormValue("Theme"))
	if theme < 1 || theme > 4 {
		utils.RespFail(c.Writer, "请输入有效的主题！")
		return
	}
	articleList = models.GetArticleByTheme(theme)
	utils.RespOk(c.Writer, articleList[0], "成功获取主题文章列表")
}
func GetNextArticle(c *gin.Context) {
	n := len(articleList)
	i++
	i %= n
	utils.RespOk(c.Writer, articleList[i], "获取了一篇文章")
}
func GetLastArticle(c *gin.Context) {
	n := len(articleList)
	i--
	i += n
	i %= n
	utils.RespOk(c.Writer, articleList[i], "获取了一篇文章")
}
