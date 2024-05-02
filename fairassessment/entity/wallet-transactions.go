package entity

import (
	"time"
)

type WalletTransactions struct {
	Id int32 `json:"id"`
	Reference string `json:"reference"`
	AccountId int32 `json:"account_id"`
	Type string `json:"type"`
	Amount float64 `json:"amount"`
	CreatedDate time.Time `json:"createdDate"`
	CreatedBy string `json:"createdBy"`
}
