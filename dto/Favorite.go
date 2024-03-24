package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Favorite struct {
	Name  string
	Color []Color `json:"color"`
}
type Color struct {
	Id int
	R  string
	G  string
	B  string
	A  string
}

func (f Favorite) Value() (driver.Value, error) {
	// 将 Favorite 结构体转换为 JSON 格式的字符串
	value, err := json.Marshal(f)
	if err != nil {
		return nil, err
	}
	return string(value), nil
}

// Scan 方法将数据库中的值转换为 Favorite 结构体
func (f *Favorite) Scan(value interface{}) error {
	// 将数据库中的值解析为字符串
	stringValue, ok := value.(string)
	if !ok {
		return errors.New("不是Favorite类型")
	}

	// 将 JSON 格式的字符串解析为 Favorite 结构体
	if err := json.Unmarshal([]byte(stringValue), f); err != nil {
		return err
	}

	return nil
}
