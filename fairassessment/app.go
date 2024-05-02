package main

import (
	"log"
	"github.com/gin-gonic/gin"
	Endless "github.com/fvbock/endless"
	KafkaService "money.com/kafkaservice"
    Migration "money.com/migration"
	wallet "money.com/wallet"
)

func main() {
    
	router := gin.Default();
	apiRoutes := router.Group("/api/v1")
	walletRouter := apiRoutes.Group("/wallet")

	var wallet wallet.WalletService
	walletRouter.POST("/credit-account", wallet.CreditWallet)
	walletRouter.POST("/debit-account", wallet.DebitWallet)
	
	backgroundJobs()
	migrateDatabase :=  func(){
		var migration Migration.Migration = &Migration.MigrationService{}
		migration.MigrateTables();
	}

	migrateDatabase();

	if err := Endless.ListenAndServe("localhost:8091", router); err != nil {
		log.Fatal("failed run app: ", err)
	}

}

func backgroundJobs(){
	var kafkaServiceProducer KafkaService.KafkaProducerService;
	go kafkaServiceProducer.CreateKafKaTopic();

	var kafkaServiceConsumer KafkaService.KafkaConsumerService; 
	go kafkaServiceConsumer.ConsumeKafkaTopics();
}































