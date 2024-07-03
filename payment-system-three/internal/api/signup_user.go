package api

import (
	"payment-system-three/internal/models"
	"payment-system-three/internal/util"
	"strings"

	"github.com/gin-gonic/gin"
)

// Create User
func (u *HTTPHandler) CreateUser(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBind(&user); err != nil {
		util.Response(c, "invalid request", 400, err.Error(), nil)
		return
	}
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)
	user.DateOfBirth = strings.TrimSpace(user.DateOfBirth)
	user.Phone = strings.TrimSpace(user.Phone)
	user.Address = strings.TrimSpace(user.Address)

	if user.FirstName == "" {
		util.Response(c, "First name must not be empty", 400, nil, nil)
		return
	}
	if user.LastName == "" {
		util.Response(c, "Last name must not be empty", 400, nil, nil)
		return
	}
	if user.Email == "" {
		util.Response(c, "Email must not be empty", 400, nil, nil)
		return
	}
	if user.Password == "" {
		util.Response(c, "Password must not be empty", 400, nil, nil)
		return
	}
	if user.DateOfBirth == "" {
		util.Response(c, "Date of birth must not be empty", 400, nil, nil)
		return
	}
	if user.Phone == "" {
		util.Response(c, "Phone must not be empty", 400, nil, nil)
		return
	}
	if user.Address == "" {
		util.Response(c, "Address must not be empty", 400, nil, nil)
		return
	}

	//Test if email exist
	isEmailExist, _ := u.Repository.FindUserByEmail(user.Email)
	if isEmailExist != nil {
		util.Response(c, "Email already exist", 400, nil, nil)
		return
	}

	// Hash the Password
	hashPassword, err := util.HashPassword(user.Password)
	if err != nil {
		util.Response(c, "Internal server error", 500, err.Error(), nil)
		return
	}
	user.Password = hashPassword

	// Test if user exist
	err = u.Repository.CreateUser(user)
	if err != nil {
		util.Response(c, "User not created", 400, err.Error(), nil)
		return
	}
	util.Response(c, "User created", 200, nil, nil)

	// generate account number
	acctNo, err := util.GenerateAccountNumber()
	if err != nil {
		util.Response(c, "could not generate account number", 500, "Internal server error", nil)
		return
	}

	// Update User struct
	user.AccountNo = acctNo

	// set available balance to zero
	user.AvailableBalance = 0.0

	// func (u *HTTPHandler) TransferFunds(c *gin.Context){
	// declare request body  (need account number and amount)

	// bind JSON data to struct

	// get user from context

	// validate the amount

	// check if account number exist

	// check if amount being transferred is less than the user current balance

	// persist the data into the db
	//}

}
