package service

import (
	"color/dto"
	"color/models"
	"color/utils"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

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

func rands() string {
	rand.Seed(int64(time.Now().UnixNano()))
	return strconv.Itoa(int(rand.Int31n(2) + 1))
}

func token(c *gin.Context) string {
	authHeader, _ := c.Cookie("Authorization")
	if authHeader == "" {
		utils.RespFail(c.Writer, "没有token信息")
	}
	return authHeader
}

// Todo 刷新token的操作
func User(c *gin.Context) (models.User, dto.UserInfo) {
	id, _ := c.Get("userInfoId")
	user := models.FindUserById(strconv.Itoa(int(id.(uint64))))
	userinfo := dto.FindUserInfo(strconv.Itoa(user.UserInfoId))
	return user, userinfo
}
func deleteFavorite(favorites []dto.Favorite, name string) ([]dto.Favorite, bool) {
	//移位法，切换
	ok := false
	for i, favorite := range favorites {
		if favorite.Name == name {
			favorites = append(favorites[:i], favorites[i+1:]...)
			ok = true
		}
	}
	return favorites, ok
}
func seachFavorite(favorites []dto.Favorite, name string) (dto.Favorite, bool) {
	//移位法，切换
	ok := false
	for _, favorite := range favorites {
		if favorite.Name == name {
			ok = true
			return favorite, ok
		}
	}
	return dto.Favorite{}, ok
}
func changFavorite(favorites []dto.Favorite, name, newname string) ([]dto.Favorite, bool) {
	//移位法，切换
	ok := false
	for i, favorite := range favorites {
		if favorite.Name == name {
			ok = true
			favorites[i].Name = newname
			break
		}
	}
	return favorites, ok
}
func appendColor(favorites []dto.Favorite, name string, color dto.Color) ([]dto.Favorite, bool) {
	//移位法，切换
	ok := false
	for i, favorite := range favorites {
		if favorite.Name == name {
			ok = true
			favorites[i].Color = append(favorite.Color, color)
			ok = true
			break
		}
	}
	return favorites, ok
}
func deleteColor(favorites []dto.Favorite, name string, id int) ([]dto.Favorite, bool) {
	//移位法，切换
	ok := false
	for i, favorite := range favorites {
		if favorite.Name == name {
			favorites[i].Color, ok = deletecolor(favorite.Color, id)
			break
		}
	}
	return favorites, ok
}
func deletecolor(colors []dto.Color, id int) ([]dto.Color, bool) {
	//移位法，切换
	ok := false
	for i, color := range colors {
		if color.Id == id {
			colors = append(colors[:i], colors[i+1:]...)
			ok = true
		}
	}
	return colors, ok
}
