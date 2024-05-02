package dto

import (
	Entity "money.com/entity"
)
type Payload struct {
	AccountId int32 `json:"account_id"`
	Reference string `json:"reference"`
	Amount float64 `json:"amount`
}

type KafkaWalletPayload struct {
	Wallet Entity.Wallet
	WalletTransaction Entity.WalletTransactions
	Activity Entity.Activity
}

