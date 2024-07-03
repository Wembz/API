package repository

import "payment-system-three/internal/models"

// TokenInBlacklist checks if token is already in the blacklist collection
func (p *Postgres) TokenInBlacklist(token *string) bool {
	tok := &models.Blacklist{}
	if err := p.DB.Where("token = ?", token).First(&tok).Error; err != nil {
		return false
	}
	return true
}
