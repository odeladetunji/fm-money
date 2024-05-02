package wallet

import (
	"net/http"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	Dto "money.com/dto"
	Entity "money.com/entity"
	Repository "money.com/repository"
	KafKaService "money.com/kafkaservice"
)

type WalletService struct {

}

var walletRepository Repository.WalletRepository = &Repository.WalletRepo{}

func(wall *WalletService) DebitWallet(ctx *gin.Context){
	
	var payload Dto.Payload
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	wallet, err2 := walletRepository.GetWalletByAccountId(payload.AccountId)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
		return 
	}

	if wallet.Id == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "wallet not found"})
		return 
	}

	if wallet.IsLien {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "A transaction is still pending"})
		return 
	}

	if payload.Amount > wallet.Balance {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Amount greater than account balance"})
		return 
	}

	walletBalance  := wallet.Balance - payload.Amount
    wallet.Balance = walletBalance

	currentDate := time.Now()
	var walletTransaction Entity.WalletTransactions
	walletTransaction.Type = "DEBIT"
	walletTransaction.Reference = payload.Reference
	walletTransaction.Amount = payload.Amount
	walletTransaction.CreatedDate = currentDate
	walletTransaction.CreatedBy = "Admin"

    var activity Entity.Activity
	activity.Reference = payload.Reference
	activity.Amount = payload.Amount
	activity.Action = "DEBIT_WALLET"
	activity.CreatedDate = currentDate
	activity.CreatedBy = "Admin"

	var kafkaWalletPayload Dto.KafkaWalletPayload
	kafkaWalletPayload.Wallet = *wallet
	kafkaWalletPayload.WalletTransaction = walletTransaction
	kafkaWalletPayload.Activity = activity

	// Push to Kafka at this point
    var kafKaService KafKaService.KafkaProducerService
	kafKaService.PushToKafkaProducer(&kafkaWalletPayload, "CREDIT_OR_DEBIT")

	ctx.JSON(http.StatusOK, gin.H{"data": "success"})
	
}

func(wall *WalletService) CreditWallet(ctx *gin.Context){
	
	var payload Dto.Payload
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	wallet, err2 := walletRepository.GetWalletByAccountId(payload.AccountId)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
		return 
	}

	if wallet.Id == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "wallet not found"})
		return 
	}

	if wallet.IsLien {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "A transaction is still pending"})
		return 
	}

	walletBalance  := wallet.Balance + payload.Amount
    wallet.Balance = walletBalance

	currentDate := time.Now()
	var walletTransaction Entity.WalletTransactions
	walletTransaction.Type = "CREDIT"
	walletTransaction.Reference = payload.Reference
	walletTransaction.Amount = payload.Amount
	walletTransaction.CreatedDate = currentDate
	walletTransaction.CreatedBy = "Admin"

    var activity Entity.Activity
	activity.Reference = payload.Reference
	activity.Amount = payload.Amount
	activity.Action = "CREDIT_WALLET"
	activity.CreatedDate = currentDate;
	activity.CreatedBy = "Admin"

	var kafkaWalletPayload Dto.KafkaWalletPayload
	kafkaWalletPayload.Wallet = *wallet
	kafkaWalletPayload.WalletTransaction = walletTransaction
	kafkaWalletPayload.Activity = activity

	//PUSH TO KAFKA
	var kafKaService KafKaService.KafkaProducerService
	kafKaService.PushToKafkaProducer(&kafkaWalletPayload, "CREDIT_OR_DEBIT")

	ctx.JSON(http.StatusOK, gin.H{"data": "success"})
	
}

func(wall *WalletService) FetchWallet(ctx *gin.Context){
	
	if len(ctx.Query("account_id")) < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "account_id required"})
		return 
	}

	accountId, err := strconv.Atoi(ctx.Query("account_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	wallet, err2 := walletRepository.GetWalletByAccountId(int32(accountId))
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
		return 
	}

	if wallet.Id == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "wallet not found"})
		return 
	}
	
	ctx.JSON(http.StatusOK, gin.H{"data": "success"})
}





