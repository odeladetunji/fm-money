package entity

import (
	"time"
)

type Wallet struct {
	Id int32 `json:"id"`
	AccountId int32 `json:"account_id"`
	Balance float64 `json:"balance"`
	IsLien bool
	CreatedDate time.Time `json:"createdDate"`
	CreatedBy string `json:"createdBy"`
}



