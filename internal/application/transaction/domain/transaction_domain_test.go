package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTransactionDomain(t *testing.T) {
	t.Run("should create a new transaction", func(t *testing.T) {
		transaction, err := NewTransactionDomain(1000, TransactionTypeCredit, time.Now(), "any description")

		assert.Nil(t, err)
		assert.Equal(t, int64(1000), transaction.Amount)
		assert.Equal(t, "credit", transaction.TransactionType)
	})

	t.Run("should return error ErrInvalidAmount", func(t *testing.T) {
		transaction, err := NewTransactionDomain(-1, TransactionTypeCredit, time.Now(), "any description")

		assert.Nil(t, transaction)
		assert.Equal(t, ErrInvalidAmount, err.Error())
	})

	t.Run("should return error ErrInvalidTransactionType", func(t *testing.T) {
		transaction, err := NewTransactionDomain(100, "another-type", time.Now(), "any description")

		assert.Nil(t, transaction)
		assert.Equal(t, ErrInvalidTransactionType, err.Error())
	})

	t.Run("should return error ErrInvalidDescription", func(t *testing.T) {
		transaction, err := NewTransactionDomain(100, TransactionTypeCredit, time.Now(), "")
		assert.Nil(t, transaction)
		assert.Equal(t, ErrInvalidDescription, err.Error())
	})

	t.Run("should return error ErrInvalidTransactionDate", func(t *testing.T) {
		var transactionDate time.Time
		transaction, err := NewTransactionDomain(100, TransactionTypeCredit, transactionDate, "description")
		assert.Nil(t, transaction)
		assert.Equal(t, ErrInvalidTransactionDate, err.Error())
	})

	t.Run("should return debit transaction", func(t *testing.T) {
		transaction, err := NewTransactionDomain(100, TransactionTypeDebit, time.Now(), "description")

		assert.Nil(t, err)
		assert.True(t, transaction.IsDebit())
	})

	t.Run("should return credit transaction", func(t *testing.T) {
		transaction, err := NewTransactionDomain(100, TransactionTypeCredit, time.Now(), "description")

		assert.Nil(t, err)
		assert.True(t, transaction.IsCredit())
	})
}
