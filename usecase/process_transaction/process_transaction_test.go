package process_transaction

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mock_broker "github.com/ricardo-ltm/fullcycle5-microservice/adapter/broker/mock"
	"github.com/ricardo-ltm/fullcycle5-microservice/domain/entity"
	mock_repository "github.com/ricardo-ltm/fullcycle5-microservice/domain/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestProcessTransaction_ExecuteInvalidCreditCardNumber(t *testing.T) {
	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "40000000000000000",
		CreditCardName:            "Fulando da Silva",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    200,
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "invalid credit card number",
	}

	// Mockgen Init
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	// Mockegen Kafak Init
	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), "transaction_result")

	// Usecase Init
	usecase := NewProcessTransaction(repositoryMock, producerMock, "transaction_result")
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransaction_ExecuteRejectedTransaction(t *testing.T) {
	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "5545316191505545",
		CreditCardName:            "Fulando da Silva",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    1200,
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "you don't have limit for this transaction",
	}

	// Mockgen Init
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	// Mockegen Kafak Init
	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), "transaction_result")

	// Usecase Init
	usecase := NewProcessTransaction(repositoryMock, producerMock, "transaction_result")
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransaction_ExecuteApprovedTransaction(t *testing.T) {
	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "5545316191505545",
		CreditCardName:            "Fulando da Silva",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    950,
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}

	// Mockgen Init
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	// Mockegen Kafak Init
	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), "transaction_result")

	// Usecase Init
	usecase := NewProcessTransaction(repositoryMock, producerMock, "transaction_result")
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
