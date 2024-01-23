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
	if !isMatchPhone(user.Phone) {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1, //失败
			"message": "电话号码无效",
			"data":    nil,
		})
		return
	}
	password := c.Request.FormValue("password")
	repassword := c.Request.FormValue("repassword")
	if !isStrongPassword(password) {
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
		if !signed(user, c) {
			return
		}
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
		"message": "验证码已发送，请注意查收,code:" + string(code),
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
		user.Password, _ = utils.GetPwd("111111Az*")
		models.CreateUser(user)
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "已自动帮您注册，默认密码为111111Az*",
			"data":    user,
		})
		return
	}
	if !signed(user, c) {
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
	user := models.FindUser(phone)
	if !isMatchPhone(phone) {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "手机号码无效",
			"data":    nil,
		})
		return
	}
	if user.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户尚未注册",
			"data":    nil,
		})
		return
	}
	if !isStrongPassword(password) {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "密码无效",
			"data":    nil,
		})
		return
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
	if len(phone) != 11 {
		flag = false
	}
	return flag
}

func isStrongPassword(password string) bool {
	// 密码长度在8到16之间
	if len(password) < 8 || len(password) > 16 {
		return false
	}

	hasUpperCase := false
	hasLowerCase := false
	hasDigit := false
	hasSpecialChar := false

	for _, char := range password {
		ascii := int(char)

		// 检查大写字母
		if ascii >= 65 && ascii <= 90 {
			hasUpperCase = true
		}

		// 检查小写字母
		if ascii >= 97 && ascii <= 122 {
			hasLowerCase = true
		}

		// 检查数字
		if ascii >= 48 && ascii <= 57 {
			hasDigit = true
		}

		// 检查特殊字符
		if (ascii >= 33 && ascii <= 47) || (ascii >= 58 && ascii <= 64) || (ascii >= 91 && ascii <= 96) || (ascii >= 123 && ascii <= 126) {
			hasSpecialChar = true
		}
	}

	// 检查是否满足所有条件
	return hasUpperCase && hasLowerCase && hasDigit && hasSpecialChar
}
func signed(user models.User, c *gin.Context) bool {
	// 查询数据库，通过用户密码拿到 userId
	userId := user.ID
	// token 过期时间 12 h，Time 类型
	var expiredTime = time.Now().Add(12 * time.Hour)

	// 生成 token string
	tokenStr, tokenErr := utils.GenerateToken(uint64(userId), expiredTime)
	if tokenErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "未能生成令牌token",
			"data":    nil,
		})
		return false
	}
	// 设置响应头信息的 token
	c.SetCookie("Authorization", tokenStr, 60, "/", "127.0.0.1", false, true)
	return true
}
