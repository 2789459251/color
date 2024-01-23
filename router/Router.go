package router

import (
	"color/service"
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
	//测试
	testGroup.GET("/getColorlist", service.GetColor)
	testGroup.GET("/method1List", service.Method1list)
	//功能
	r.POST("/attach/upload", service.Upload)
	r.POST("/attach/uploadshi", service.Uploadshi)
	return r
}
