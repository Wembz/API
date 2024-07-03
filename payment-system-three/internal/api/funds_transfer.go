package api

import (
	"payment-system-three/internal/models"
	"payment-system-three/internal/util"

	"github.com/gin-gonic/gin"
)

// declare request body

//bind JSON data to struct

func (u *HTTPHandler) TransferFunds(c *gin.Context) {
	var funds *models.TransferFunds
	if err := c.ShouldBind(&funds); err != nil {
		util.Response(c, "invalid request", 400, "bad request body", nil)
		return
	}

	//get user from context
	user, err := u.GetUserFromContext(c)
	if err != nil {
		util.Response(c, "User not logged in", 500, "user not found", nil)
		return
	}

	// Validate the amount
	if funds.Amount <= 0 {
		util.Response(c, "Invalid amount", 500, "Bad Request body", nil)
		return
	}

	//check if account exist
	recipient, err := u.Repository.FindUserByAccountNumber(funds.AccountNo)
	if err != nil {
		util.Response(c, "Balance insufficient funds", 400, "Bad Request", nil)
		return
	}

	// persist the date into the db

	err = u.Repository.TransferFunds(user, recipient, funds.Amount)
	if err != nil {
		util.Response(c, "Transfer not possible", 500, "transfer not successful", nil)
		return

	}
	util.Response(c, "Transfer was done successfully", 200, "Transfer successful", nil)

	//View Transaction
	err = u.Repository.TransferFunds(user, recipient, funds.Amount)
	if err != nil {
		util.Response(c, "User unable to view transcation", 500, "Transcation not found", nil)
		return
	}
	util.Response(c, "Transfer was done successfully", 200, "Transfer successful", nil)

}
