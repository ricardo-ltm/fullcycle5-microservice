package main

import (
	"database/sql"
	"encoding/json"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ricardo-ltm/fullcycle5-microservice/adapter/broker/kafka"
	"github.com/ricardo-ltm/fullcycle5-microservice/adapter/factory"
	"github.com/ricardo-ltm/fullcycle5-microservice/adapter/presenter/transaction"
	"github.com/ricardo-ltm/fullcycle5-microservice/usecase/process_transaction"
)

func main() {
	db, err := sql.Open("sqlite3", "teste.db")
	if err != nil {
		log.Fatal(err)
	}

	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()

	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}

	kafkaPresenter := transaction.NewTransactionKafkaPresenter()
	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)

	var msgChan = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "goapp",
		"group.id":          "goapp",
	}

	topics := []string{"transactions"}
	consumer := kafka.NewConsumer(configMapConsumer, topics)
	go consumer.Consumer(msgChan)

	usecase := process_transaction.NewProcessTransaction(repository, producer, "transactions_result")

	for msg := range msgChan {
		var input process_transaction.TransactionDtoInput
		json.Unmarshal(msg.Value, &input)
		usecase.Execute(input)
	}
}
