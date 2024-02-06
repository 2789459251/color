package dto

import "time"

type History struct {
	Time       time.Time
	Result     string
	ResultInfo []ResultInfo
}
