package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type ResultInfo struct {
	Key   string
	Mykey string
	Point string
	Image string
	Flag  bool
}
type ResultInfoSlice []ResultInfo

// Value 将 ResultInfoSlice 转换为 JSON 格式的字符串
func (ri ResultInfoSlice) Value() (driver.Value, error) {
	// 将 ResultInfo 结构体切片转换为 JSON 格式的字符串
	value, err := json.Marshal(ri)
	if err != nil {
		return nil, err
	}
	return string(value), nil
}

// Scan 将数据库中的值解析为 ResultInfoSlice 结构体切片
func (ri *ResultInfoSlice) Scan(value interface{}) error {
	// 将数据库中的值解析为字符串
	stringValue, ok := value.(string)
	if !ok {
		return errors.New("不是 ResultInfo 切片类型")
	}

	// 将 JSON 格式的字符串解析为 ResultInfo 结构体切片
	var resultInfoSlice []ResultInfo
	if err := json.Unmarshal([]byte(stringValue), &resultInfoSlice); err != nil {
		return err
	}

	*ri = resultInfoSlice

	return nil
}
