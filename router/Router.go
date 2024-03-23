package router

import (
	"color/service"
	"color/utils"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Static("/Asset", "Asset/")
	userGroup := r.Group("/user")
	testGroup := r.Group("/test")
	//用户
	userGroup.POST("/createUser", service.CreateUser)
	userGroup.POST("/loginByPassword", service.LoginByPassword)
	userGroup.POST("/sendCode", service.SendCode)
	userGroup.POST("/loginByCode", service.LoginByCode)
	userGroup.POST("/resetPassword", service.ResetPassword)
	//测试---用户获取
	testGroup.Use(utils.JWTAuth())
	//testGroup.GET("/color", service.GetColor)
	testGroup.GET("/method1", service.Method1)
	testGroup.GET("/methodNextTest", service.GetNextTest)
	testGroup.GET("/methodLastTest", service.GetLastTest)

	//检测---提交后判断
	testGroup.GET("/GetHighest", service.GetHighest)
	testGroup.POST("/RefreshHighest", service.SetHighest)
	testGroup.POST("/JudgeMethod1", service.Judge_m)

	//功能
	attachGroup := r.Group("/attach")
	attachGroup.Use(utils.JWTAuth())

	attachGroup.POST("/upload", service.Upload)
	attachGroup.POST("/uploadshi", service.Uploadshi)
	attachGroup.GET("/history", service.History)
	//收藏夹管理
	attachGroup.POST("/uploadFavorite", service.UploadFavorite)
	attachGroup.GET("/Favorite", service.Favorite)
	attachGroup.DELETE("/cancelFavorite", service.CancelFavorite)

	//文章管理
	//上传文章
	attachGroup.POST("/uploadArticle", service.UploadArticle)
	//获取专题所有文章
	attachGroup.GET("/getArticles", service.GetArticles)
	attachGroup.GET("/getNextArticle", service.GetNextArticle)
	attachGroup.GET("/getLastArticle", service.GetLastArticle)
	return r
}
