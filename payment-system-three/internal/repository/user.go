package repository

import (
	"payment-system-three/internal/models"
	"time"
)

// Find user By Email in Postgres
func (p *Postgres) FindUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	// Error message if Postgres Email database fails
	if err := p.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}


// Create a user in Postgres database

func (p *Postgres) CreateUser(user *models.User) error {

	// Error message if Postgres database fails
	if err := p.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}


// Update a user in Postgres database

func (p *Postgres) UpdateUser(user *models.User) error {

	// Error message if Postgres database fails
	if err := p.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (p *Postgres) TransferFunds(user *models.User, recipient *models.User, amount float64) error {
	tx := p.DB.Begin()

	// deduct the amount from the payer
	user.AvailableBalance -= amount

	// add the amount to the recipient
	recipient.AvailableBalance += amount

	// save the transcation for payer
	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// save the transcation for recipient
	if err := tx.Save(recipient).Error; err != nil {
		tx.Rollback()
		return err
	}

	// save transcation in the table
	transaction := models.Transaction{
		PayerAccount:      user.AccountNo,
		RecipientAccount:  recipient.AccountNo,
		TransactionType:   "debit",
		TransactionAmount: amount,
		TransactionDate:   time.Time{},
	}

	// save the transcation
	if err := tx.Create(transaction).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil

}

// find user by account
func (p *Postgres) FindUserByAccountNumber(account_no int) (*models.User, error) {
	user := &models.User{}

	if err := p.DB.Where("accountNo = ?", account_no).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// View transaction by Account
func (p *Postgres) GetTransactionByAccountNumber(account_no int) (*[]models.Transaction, error) {
	transaction := &[]models.Transaction{}

	if err := p.DB.Where("payer_account = ? OR recipients_account = ? ", account_no, account_no).Find(&transaction).Error; err != nil {
	}
	return transaction, nil
}

func (p Postgres) UpdateFunds(transcation *models.Transaction) error {
	if err := p.DB.Save(transcation).Error; err != nil {
		return err
	}
	return nil
}
