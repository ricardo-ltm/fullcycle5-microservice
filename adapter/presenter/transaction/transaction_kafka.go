package transaction

import (
	"encoding/json"

	"github.com/ricardo-ltm/fullcycle5-microservice/usecase/process_transaction"
)

type KafkaPresenter struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func NewTransactionKafkaPresenter() *KafkaPresenter {
	return &KafkaPresenter{}
}

func (kp *KafkaPresenter) Bind(input interface{}) error {
	// o comando colacado entre parenteses é para informat o tipo devido a interface
	kp.ID = input.(process_transaction.TransactionDtoOutput).ID
	kp.Status = input.(process_transaction.TransactionDtoOutput).Status
	kp.ErrorMessage = input.(process_transaction.TransactionDtoOutput).ErrorMessage
	return nil
}

func (kp *KafkaPresenter) Show() ([]byte, error) {
	j, err := json.Marshal(kp)
	if err != nil {
		return nil, err
	}
	return j, nil
}
