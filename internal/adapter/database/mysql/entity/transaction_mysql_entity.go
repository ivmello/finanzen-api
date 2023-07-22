package model

import "time"

type TransactionMysqlEntity struct {
	ID              string    `json:"id" gorm:"primaryKey"`
	Amount          int64     `json:"amount"`
	TransactionType string    `json:"transaction_type"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
