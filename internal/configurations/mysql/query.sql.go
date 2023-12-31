// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: query.sql

package mysql

import (
	"context"
	"database/sql"
	"time"
)

const createTransaction = `-- name: CreateTransaction :execresult
INSERT INTO transactions (
  id, amount, transaction_type, transaction_date, description, created_at, updated_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?
)
`

type CreateTransactionParams struct {
	ID              string
	Amount          int64
	TransactionType string
	TransactionDate time.Time
	Description     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createTransaction,
		arg.ID,
		arg.Amount,
		arg.TransactionType,
		arg.TransactionDate,
		arg.Description,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
}

const getTransaction = `-- name: GetTransaction :one
SELECT id, amount, transaction_type, transaction_date, description, created_at, updated_at FROM transactions
WHERE id = ? LIMIT 1
`

func (q *Queries) GetTransaction(ctx context.Context, id string) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, getTransaction, id)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.Amount,
		&i.TransactionType,
		&i.TransactionDate,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
