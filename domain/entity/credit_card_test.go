package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("40000000000000000", "Fulano da Silva", 12, 2024, 123)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("5545316191505545", "Fulano da Silva", 12, 2024, 123)
	assert.Nil(t, err)

}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("5545316191505545", "Fulano da Silva", 13, 2024, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5545316191505545", "Fulano da Silva", 0, 2024, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5545316191505545", "Fulano da Silva", 11, 2024, 123)
	assert.Nil(t, err)

}

func TestCreditCardYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1, 0, 0)
	_, err := NewCreditCard("5545316191505545", "Fulano da Silva", 3, lastYear.Year(), 123)
	assert.Equal(t, "invalid expiration year", err.Error())

	_, err = NewCreditCard("5545316191505545", "Fulano da Silva", 3, time.Now().Year(), 123)
	assert.Nil(t, err)
}
