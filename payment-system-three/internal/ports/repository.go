package ports

import "payment-system-three/internal/models"

type Repository interface {
	FindUserByEmail(email string) (*models.User, error)
	FindAdminByEmail(email string) (*models.Admin, error)
	TokenInBlacklist(token *string) bool
	CreateUser(user *models.User) error
	CreateAdmin(admin *models.Admin) error
	UpdateUser(user *models.User) error
	UpdateAdmin(user *models.Admin) error
	FindUserByAccountNumber(account_no int) (*models.User, error)
	TransferFunds(user *models.User, recipient *models.User, amount float64) error
	//ViewUserBalance(user *models.User)error
	//ViewUserTranscationHistory(transaction *models.Transaction)error
	//AdminViewAllTransactionsHistory(transaction *models.Transaction)error
	//AdminViewUserBalance(user *models.User)error
	GetTransactionByAccountNumber(account_no int) (*[]models.Transaction, error)
}
