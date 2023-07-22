package repository

import (
	"database/sql"

	"github.com/ivmello/finanzen-api/internal/application/transaction/domain"
)

type repositoryMock struct {
	db *sql.DB
}

func NewTransactionMysqlRepositoryMock(db *sql.DB) TransactionMysqlRepositoryInterface {
	return &repositoryMock{
		db,
	}
}

func (r *repositoryMock) CreateTransaction(transaction domain.TransactionDomain) (domain.TransactionDomain, error) {
	return transaction, nil
}
