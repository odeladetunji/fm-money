package kafkaservice 

import (
	Kafka "github.com/segmentio/kafka-go"
	// JSON "encoding/json"
	"time"
	"fmt"
	"context"
	Entity "money.com/entity"
	Repository "money.com/repository"
	Dto "money.com/dto"
	JSON "encoding/json"
)

type KafkaConsumerService struct {

}

var kafkaService KafkaConsumerService;

func (kaf *KafkaConsumerService) ConnectToKafka(kafkaURL string, topic string) *Kafka.Reader {
	fmt.Println("Kafka B")
	brokers := []string{kafkaURL};
	return Kafka.NewReader(Kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  "wallet-group",
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		CommitInterval: time.Second,
	});
}

func (kaf *KafkaConsumerService) ConsumeKafkaTopics(){

	var kafkaURL string = "164.90.154.224:9092";
	var topic string = "transaction_service_credit_debit";
	var kafkaReader *Kafka.Reader = kafkaService.ConnectToKafka(kafkaURL, topic);

	for {

        fmt.Println("Start  .....");
		ctx := context.Background();
		message, err := kafkaReader.FetchMessage(ctx);
		if err != nil {
			fmt.Println(err.Error());
		}

		if message.Value != nil {
			fmt.Println("Code block entered");
			if string(message.Key) == "credit_or_debit_wallet_transaction" {
				errk := kafkaService.PersistWalletTransactionToDB(message.Value);
				if errk != nil {
					break
				}
			}
		}

		if err := kafkaReader.CommitMessages(ctx, message); err != nil {
			fmt.Println("failed to commit messages:", err.Error())
		}
	}
}

func (kaf *KafkaConsumerService) PersistWalletTransactionToDB(binaryValue []byte) error {

	var kafkaWalletPayload Dto.KafkaWalletPayload
	err := JSON.Unmarshal(binaryValue, &kafkaWalletPayload)
	if err != nil {
		return err
	}

	var wallet Entity.Wallet
	wallet = kafkaWalletPayload.Wallet
	wallet.IsLien = false

	var activity Entity.Activity
	activity = kafkaWalletPayload.Activity

	var walletTrans Entity.WalletTransactions
	walletTrans = kafkaWalletPayload.WalletTransaction

	var walletRepository Repository.WalletRepository = &Repository.WalletRepo{}
	err2 := walletRepository.CreditOrDebitWallet(context.TODO(), wallet, walletTrans, activity)
	if err2 != nil {
		return err
	}

	return nil
}



