package factory

import "github.com/ricardo-ltm/fullcycle5-microservice/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
