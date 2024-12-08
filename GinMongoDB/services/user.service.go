package services

import "github.com/MongoDB/models"

//Defining API contracts & We'll implement each functions
type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}