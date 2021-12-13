package kafka

import (
	"testing"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ricardo-ltm/fullcycle5-microservice/adapter/presenter/transaction"
	"github.com/ricardo-ltm/fullcycle5-microservice/domain/entity"
	"github.com/ricardo-ltm/fullcycle5-microservice/usecase/process_transaction"
	"github.com/stretchr/testify/assert"
)

func TestProducerPublish(t *testing.T) {

	expectedOutput := process_transaction.TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "you don't have limit for this transaction",
	}

	// configuração para conexão com o kafka
	configMap := ckafka.ConfigMap{
		// este comando faz o kafka entender que será usado como teste
		"test.mock.num.brokers": 3,
	}

	producer := NewKafkaProducer(&configMap, transaction.NewTransactionKafkaPresenter())
	err := producer.Publish(expectedOutput, []byte("1"), "test")
	assert.Nil(t, err)
}
