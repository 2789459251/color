package models

import (
	"color/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	//validate:"min=6,max=10" `valid:"matches(^1[3-9]{1}\\d{9})"`
	//`valid:"matches(^1[3-9]{1}\\d{9}$)"`
	//`valid:"matches(^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$)"`
	Phone    string `validate:"required"`
	Password string `validate:"required"`
	C_type   int    `json:"C_Type"`
}

func (table *User) TableName() string {
	return "user_basic"
}
func CreateUser(user User) *gorm.DB {
	return utils.DB.Create(&user)
}
func FindUser(phone string) User {
	user := User{}
	utils.DB.Where("phone=?", phone).First(&user)
	return user
}
func EditUserPassword(password, phone string) {
	user := FindUser(phone)
	user.Password, _ = utils.GetPwd(password)
	utils.DB.Updates(&user)
}
