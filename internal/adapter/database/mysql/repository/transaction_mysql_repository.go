package repository

import (
	"context"
	"database/sql"

	"github.com/ivmello/finanzen-api/internal/application/transaction/domain"
	"github.com/ivmello/finanzen-api/internal/configurations/mysql"
)

type TransactionMysqlRepositoryInterface interface {
	CreateTransaction(transaction domain.TransactionDomain) (domain.TransactionDomain, error)
}

type repository struct {
	db *sql.DB
}

func NewTransactionMysqlRepository(db *sql.DB) TransactionMysqlRepositoryInterface {
	return &repository{
		db,
	}
}

func (r *repository) CreateTransaction(transaction domain.TransactionDomain) (domain.TransactionDomain, error) {
	ctx := context.Background()
	queries := mysql.New(r.db)
	_, err := queries.CreateTransaction(ctx, mysql.CreateTransactionParams{
		ID:              transaction.ID,
		Amount:          transaction.Amount,
		TransactionType: transaction.TransactionType,
		TransactionDate: transaction.TransactionDate,
		Description:     transaction.Description,
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
	})
	if err != nil {
		return domain.TransactionDomain{}, err
	}

	return transaction, nil
}
