package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model

	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Password    string `json:"password" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Phone       string `json:"phone" binding:"required"`
	Address     string `json:"address" binding:"required"`
}

type LoginRequestAdmin struct {
	Email    string `json:"email"`
	Password string `son:"password"`
}

type TransferMoney struct {
	Amount    float64 `json: "amount"`
	AccountNo int     `json: account_no"`
}
