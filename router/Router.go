package router

import (
	"color/service"
	"color/utils"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

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

	//检测---提交后判断
	testGroup.GET("/GetHighest", service.GetHighest)
	testGroup.GET("/RefreshHighest", service.SetHighest)
	testGroup.POST("/JudgeMethod1", service.Judge_m)

	//功能
	r.POST("/attach/upload", service.Upload)
	r.POST("/attach/uploadshi", service.Uploadshi)
	r.GET("/attach/history", service.History)
	return r
}
