package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type History struct {
	Time       time.Time
	Result     string
	ResultInfo ResultInfoSlice `gorm:"type:json"`
}

//func (h *History) MarshalJSON() ([]byte, error) {
//	return json.Marshal(&struct {
//		Time       time.Time    `json:"time"`
//		Result     string       `json:"result"`
//		ResultInfo []ResultInfo `json:"result_info"`
//	}{
//		Time:       h.Time,
//		Result:     h.Result,
//		ResultInfo: h.ResultInfo,
//	})
//}
//
//func (h *History) UnmarshalJSON(data []byte) error {
//	// 在此处解析JSON并赋值给Time、Result和ResultInfo等其他字段
//	type Alias History
//	aux := &struct {
//		Time       time.Time    `json:"time"`
//		Result     string       `json:"result"`
//		ResultInfo []ResultInfo `json:"result_info"`
//	}{}
//	if err := json.Unmarshal(data, &aux); err != nil {
//		return err
//	}
//	h.Time = aux.Time
//	h.Result = aux.Result
//	h.ResultInfo = aux.ResultInfo
//	return nil
//}

// Implement driver.Valuer interface for History type
func (h History) Value() (driver.Value, error) {
	// Marshal History struct to JSON string
	jsonValue, err := json.Marshal(h)
	if err != nil {
		return nil, err
	}
	return string(jsonValue), nil

}

// Implement sql.Scanner interface for History type
func (h *History) Scan(value interface{}) error {
	// Check if the value is nil or not a string
	if value == nil {
		return nil
	}
	strValue, ok := value.(string)
	if !ok {
		return errors.New("输入的并非字符串")
	}

	// Unmarshal JSON string to History struct
	if err := json.Unmarshal([]byte(strValue), h); err != nil {
		return err
	}
	return nil
}
