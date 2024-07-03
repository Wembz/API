package repository

import "payment-system-three/internal/models"

// Find Admin Email in Postgres database
func (p *Postgres) FindAdminByEmail(email string) (*models.Admin, error) {
	admin := &models.Admin{}

	// Error message if Postgres Email database fails
	if err := p.DB.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

// Create Admin in Postgres database
func (p *Postgres) CreateAdmin(admin *models.Admin) error {

	// Error message if Postgres create Admin database fails
	if err := p.DB.Create(admin).Error; err != nil {
		return err
	}

	return nil
}

// Update Admin in Postgres database
func (p *Postgres) UpdateAdmin(admin *models.Admin) error {

	// Error message if Postgres Update Admin database fails
	if err := p.DB.Save(admin).Error; err != nil {
		return err
	}
	return nil
}
