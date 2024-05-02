package entity

import (
	"time"
)

type Activity struct {
	Id int32 `json:"id"`
	Reference string `json:"reference"`
	Action string `json:"action"`
	Amount float64 `json:"amount"`
	CreatedDate time.Time `json:"createdDate"`
	CreatedBy string `json:"createdBy"`
}


