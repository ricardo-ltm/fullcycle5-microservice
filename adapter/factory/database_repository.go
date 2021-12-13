package factory

import (
	"database/sql"

	repo "github.com/ricardo-ltm/fullcycle5-microservice/adapter/repository"
	"github.com/ricardo-ltm/fullcycle5-microservice/domain/repository"
)

type RepositoryDatabaseFactory struct {
	DB *sql.DB
}

func NewRepositoryDatabaseFactory(db *sql.DB) *RepositoryDatabaseFactory {
	return &RepositoryDatabaseFactory{DB: db}
}

func (r RepositoryDatabaseFactory) CreateTransactionRepository() repository.TransactionRepository {
	return repo.NewTransactionRepositoryDb(r.DB)
}
