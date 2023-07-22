package dto

import "time"

type CreateTransactionInputDto struct {
	Amount          int64  `json:"amount" validate:"required,min=1"`
	TransactionType string `json:"transaction_type" validate:"required"`
	TransactionDate string `json:"transaction_date" validate:"required"`
	Description     string `json:"description" validate:"required"`
}

type CreateTransactionOutputDto struct {
	ID              string    `json:"id"`
	Amount          int64     `json:"amount"`
	TransactionType string    `json:"transaction_type"`
	TransactionDate time.Time `json:"transaction_date"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
