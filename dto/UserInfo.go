package dto

import (
	"color/utils"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type UserInfo struct {
	gorm.Model
	Hightest int
	History  []History  `gorm:"type:json"`
	Favorite []Favorite `gorm:"type:json"`
}

func (user UserInfo) TableName() string {
	return "user_info"
}
func FindUserInfo(id string) UserInfo {
	userInfo := UserInfo{}
	// 查询并获取用户信息
	utils.DB.Table("user_info").Where("id = ?", id).First(&userInfo)

	// 查询并获取 history 字段的值
	var historyJSON string
	utils.DB.Table("user_info").Where("id = ?", id).Select("history").Scan(&historyJSON)
	// 将 historyJSON 解析为 []History 类型
	var history []History
	err := json.Unmarshal([]byte(historyJSON), &history)
	if err != nil {
		// 处理解析错误
		fmt.Println("反序列化错误history", err)
		return userInfo
	}
	// 将解析后的 history 赋值给 userInfo
	userInfo.History = history

	var favoriteJSON string
	utils.DB.Table("user_info").Where("id = ?", id).Select("favorite").Scan(&favoriteJSON)
	// 将 historyJSON 解析为 []History 类型
	var favorite []Favorite
	err = json.Unmarshal([]byte(favoriteJSON), &favorite)
	if err != nil {
		// 处理解析错误
		fmt.Println("反序列化favorite错误", err)
		return userInfo
	}
	// 将解析后的 history 赋值给 userInfo
	userInfo.Favorite = favorite
	return userInfo
}

func RefreshUserInfo(userInfo UserInfo) {
	// 将 History 字段序列化为 JSON 格式
	historyJSON, err := json.Marshal(userInfo.History)
	if err != nil {
		// 处理序列化错误
		fmt.Println("history序列化错误:", err)
		return
	}
	favoriteJSON, err := json.Marshal(userInfo.Favorite)
	if err != nil {
		// 处理序列化错误
		fmt.Println("favorite序列化错误:", err)
		return
	}

	// 更新用户信息
	if userInfo.CreatedAt.IsZero() {
		// 如果 CreatedAt 字段为零值，则将其设置为当前时间
		userInfo.CreatedAt = time.Now()
	}

	// 更新数据库中的用户信息，包括序列化后的 History 字段
	result := utils.DB.Model(&userInfo).Updates(map[string]interface{}{
		"history":  string(historyJSON),
		"favorite": string(favoriteJSON),
	})
	if result.Error != nil {
		// 处理更新错误
		fmt.Println("更新info出错:", result.Error)
		return
	}
}
