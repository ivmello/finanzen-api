package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	TransactionTypeCredit = "credit"
	TransactionTypeDebit  = "debit"
)

const (
	ErrInvalidAmount          = "invalid amount"
	ErrInvalidTransactionType = "invalid transaction type"
	ErrInvalidDescription     = "invalid description"
	ErrInvalidTransactionDate = "invalid transaction date"
)

type TransactionDomain struct {
	ID              string
	Amount          int64
	TransactionType string
	Description     string
	TransactionDate time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewTransactionDomain(amount int64, transactionType string, transactionDate time.Time, description string) (*TransactionDomain, error) {
	fmt.Println(transactionDate)

	transaction := &TransactionDomain{
		ID:              uuid.New().String(),
		Amount:          amount,
		TransactionType: transactionType,
		Description:     description,
		TransactionDate: transactionDate,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	err := transaction.Validate()
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionDomain) Validate() error {
	if t.Amount <= 0 {
		return errors.New(ErrInvalidAmount)
	}

	if t.TransactionType != TransactionTypeCredit && t.TransactionType != TransactionTypeDebit {
		return errors.New(ErrInvalidTransactionType)
	}

	if len(t.Description) == 0 {
		return errors.New(ErrInvalidDescription)
	}

	if t.TransactionDate.IsZero() {
		return errors.New(ErrInvalidTransactionDate)
	}

	return nil
}

func (t *TransactionDomain) IsDebit() bool {
	return t.TransactionType == TransactionTypeDebit
}

func (t *TransactionDomain) IsCredit() bool {
	return t.TransactionType == TransactionTypeCredit
}
