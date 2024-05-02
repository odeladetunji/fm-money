package kafkaservice

import (
	Kafka "github.com/segmentio/kafka-go"
	"context"
	Dto "money.com/dto"
	JSON "encoding/json"
	"errors"
	"strconv"
    "net"
)

type KafkaProducerService struct {

}

func (kaf *KafkaProducerService) CreateKafkaConnection(kafkaUrl string, topic string) (*Kafka.Writer) {
	return &Kafka.Writer{
		Addr:     Kafka.TCP(kafkaUrl),
		Topic:    topic,
		Balancer: &Kafka.LeastBytes{},
	}
}

func (kaf *KafkaProducerService) KafkaProducer(kafkaWriter *Kafka.Writer, key string, byteArray []byte) {

	message := Kafka.Message{
		Key: []byte(key),
		Value: []byte(byteArray),
	}

	// var req *http.Request;
	errI := kafkaWriter.WriteMessages(context.Background(), message);
	if errI != nil {
		panic(errI.Error());
	}

}

func (kaf *KafkaProducerService) PushToKafkaProducer(kafkaWalletPayload *Dto.KafkaWalletPayload, topicType string) {

	var kafkaUrl string = "164.90.154.224:9092";
	var topic string = "transaction_service_credit_debit";
	var key string;
	var byteArray []byte;
	
	if topicType == "CREDIT_OR_DEBIT" {
		key = "credit_or_debit_wallet_transaction";
		json, err := JSON.Marshal(kafkaWalletPayload);
		if err != nil {
			panic(err.Error());
		}

		byteArray = json;
	}

	//Create Connection
	var kafkaProducerService KafkaProducerService;
	kafkaWriter := kafkaProducerService.CreateKafkaConnection(kafkaUrl, topic);
	defer kafkaWriter.Close();

	//Push to Producer
	kafkaProducerService.KafkaProducer(kafkaWriter, key, byteArray);

}


func (kaf *KafkaProducerService) checkIfTopicExists(topic string) (bool, error){

	conn, err := Kafka.Dial("tcp", "164.90.154.224:9092");
	if err != nil {
		return false, errors.New(err.Error());
	}
	defer conn.Close()

	partitions, errR := conn.ReadPartitions()
	if errR != nil {
		return false, errors.New(errR.Error());
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}

	var topicIsPresent bool = false;
	for k := range m {
		if k == topic {
			topicIsPresent = true;
		}
	}

	return topicIsPresent, nil;

}

func (kaf *KafkaProducerService) CreateKafKaTopic(){
	var kafkaProducerService KafkaProducerService;

	topicCreation := func(topic string){
		topicIsPresent, errTop := kafkaProducerService.checkIfTopicExists(topic);
		if errTop != nil {
			panic(errTop.Error())
		}

		if errTop == nil && !topicIsPresent {
			conn, err := Kafka.Dial("tcp", "164.90.154.224:9092")
			if err != nil {
				panic(err.Error())
			}
			defer conn.Close();
		
			controller, err := conn.Controller()
			if err != nil {
				panic(err.Error())
			}
			var controllerConn *Kafka.Conn
			controllerConn, err = Kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
			if err != nil {
				panic(err.Error())
			}
	
			defer controllerConn.Close();
		
			topicConfigs := []Kafka.TopicConfig{
				{
					Topic:             topic,
					NumPartitions:     1,
					ReplicationFactor: 1,
				},
			}
		
			err = controllerConn.CreateTopics(topicConfigs...);
			if err != nil {
				panic(err.Error())
			}
		}
	}

	var topicList []string = []string{"transaction_service_credit_debit"};
    for i := 0; i < len(topicList); i++ {
		topicCreation(topicList[i]);
	}
}
