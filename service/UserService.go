package service

import (
	"color/models"
	"color/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"time"
)

func CreateUser(c *gin.Context) {
	//validate := validator.New()
	//Todo : 验证码 找回密码 查看密码
	user := models.User{}
	phone := c.Request.FormValue("phone")
	user2 := models.FindUser(phone)
	if user2.Password != "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1, //失败
			"message": "该号码已被使用",
			"data":    nil,
		})
		return
	}
	user.Phone = phone
	//err := validate.Var(user.Phone, "required,regexp=^1[3-9]{1}\\\\d{9}$")
	if !isMatchPassword(user.Phone) {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1, //失败
			"message": "电话号码无效",
			"data":    nil,
		})
		return
	}
	password := c.Request.FormValue("password")
	repassword := c.Request.FormValue("repassword")
	if !isMatchPassword(password) {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1, //失败
			"message": "密码无效,请输入长度在8-16位的字母数字或特殊字符",
			"data":    nil,
		})
		return
	}
	if password != repassword {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1, //失败
			"message": "密码不一致",
			"data":    nil,
		})
		return
	} else {
		user.Password, _ = utils.GetPwd(password)
		models.CreateUser(user)
		c.JSON(http.StatusOK, gin.H{
			"code":    0, //成功
			"message": "注册成功",
			"data":    user,
		})
		return
	}

}
func LoginByPassword(c *gin.Context) {
	//不要明文存储密码=
	phone := c.Request.FormValue("phone")
	user := models.FindUser(phone)
	fmt.Println(user)
	if user.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1, //失败
			"message": "用户尚未注册",
			"data":    nil,
		})
		return

	}
	password := c.Request.FormValue("password")
	if utils.ComparePwd(user.Password, password) {
		c.JSON(http.StatusOK, gin.H{
			"code":    0, //成功
			"message": "欢迎回来",
			"data":    user,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1, //失败
			"message": "密码错误",
			"data":    nil,
		})
		return
	}
}
func SendCode(c *gin.Context) {
	//post请求->phone
	phone := c.Request.FormValue("phone")
	if !isMatchPhone(phone) {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"massage": "手机号码无效",
			"data":    nil,
		})
	}
	code := utils.GenerateSMSCode()
	fmt.Println("验证码：", code)
	//将验证码存入redis
	utils.Red.Set(c, phone, code, 5*time.Minute)
	c.JSON(http.StatusOK, gin.H{
		"code":    0, //成功
		"message": "验证码已发送，请注意查收",
		"data":    nil,
	})
}
func LoginByCode(c *gin.Context) {
	//post请求->phone
	phone := c.Request.FormValue("phone")
	if !isMatchPhone(phone) {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"massage": "手机号码无效",
			"data":    nil,
		})
	}
	code := c.Request.FormValue("code")
	//查询redis
	cacheCode, _ := utils.Red.Get(c, phone).Result()
	//不一致就不放行
	if code != cacheCode {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "验证码错误",
			"data":    nil,
		})
		return
	}
	//一致就放行->如果用户尚且未注册，直接可以注册并告知默认密码
	user := models.FindUser(phone)
	if user.Password == "" {
		user.Phone = phone
		user.Password = "111111Az*"
		models.CreateUser(user)
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "已自动帮您注册，默认密码为111111Az*",
			"data":    user,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "欢迎回来",
		"data":    user,
	})
	return
}
func ResetPassword(c *gin.Context) {
	phone := c.Request.FormValue("phone")
	password := c.Request.FormValue("password")
	if !isMatchPhone(phone) {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "手机号码无效",
			"data":    nil,
		})
	}
	if !isMatchPassword(password) {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "密码无效",
			"data":    nil,
		})
	}
	models.EditUserPassword(password, phone)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "重置成功",
		"data":    nil,
	})
}
func isMatchPhone(phone string) bool {
	flag, _ := regexp.Match("^1[3-9]{1}\\d{9}", []byte(phone))
	return flag
}
func isMatchPassword(password string) bool {
	flag, _ := regexp.Match("[A-z0-9~@\\*()_]{8,16}", []byte(password))
	return flag
}
