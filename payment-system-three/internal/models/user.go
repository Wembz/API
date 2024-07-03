package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	
	FirstName        string    `json:"first_name" binding:"required"`
	LastName         string    `json:"last_name" binding:"required"`
	Password         string    `json:"password" binding:"required"`
	DateOfBirth      string    `json:"date_of_birth" binding:"required"`
	AccountNo        int       `json: "account_no"`
	AvailableBalance float64   `json: "available_balance"`
	Email            string    `json:"email" binding:"required,email"`
	Phone            string    `json:"phone" binding:"required"`
	Address          string    `json:"address" binding:"required"`
	LoginCounter     int       `json:"login_counter"`
}

type LoginRequestUser struct {
	Email    string `json:"email"`
	Password string `son:"password"`
}

type TransferFunds struct {
	AccountNo int     `json: account_no"`
	Amount    float64 `json: "amount"`
}

type Addfuns struct {
	Amount float64 `json: "amount"`
}
type Transaction struct {
	gorm.Model

	PayerAccount      int       `json: "payer_account"`
	RecipientAccount  int       `json: "recipient_account"`
	TransactionAmount float64   `json: "transaction_account"`
	TransactionDate   time.Time `json: "transaction_date"`
	TransactionType   string    `json: "transaction_type"`
}
