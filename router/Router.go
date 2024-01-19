package router

import (
	"color/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	//用户
	r.POST("/user/CreateUser", service.CreateUser)
	r.POST("/user/loginByPassword", service.LoginByPassword)
	r.POST("/user/sendCode", service.SendCode)
	r.POST("/user/loginByCode", service.LoginByCode)
	r.POST("/user/resetPassword", service.ResetPassword)
	//测试
	r.GET("/test/getColorlist", service.GetColorlist)
	r.GET("/test/Method1List", service.Method1list)
	return r
}
