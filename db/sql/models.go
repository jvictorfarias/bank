// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"database/sql"
	"time"
)

type Account struct {
	ID        int32     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   string    `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type Entry struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	// Can be negative or positive
	Amount    string       `json:"amount"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type Transfer struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// Must be positive
	Amount    string    `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
