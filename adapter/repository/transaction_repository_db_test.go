package repository

import (
	"os"
	"testing"

	"github.com/ricardo-ltm/fullcycle5-microservice/adapter/repository/fixture"
	"github.com/ricardo-ltm/fullcycle5-microservice/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRepositoryDbInsert(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	repository := NewTransactionRepositoryDb(db)
	err := repository.Insert("1", "1", 12.1, entity.APPROVED, "")
	assert.Nil(t, err)
}
