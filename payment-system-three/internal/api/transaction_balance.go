package api

import (
	"payment-system-three/internal/util"

	"github.com/gin-gonic/gin"
)

func (u *HTTPHandler) ViewUserTransactionHistory(c *gin.Context) {

	user, err := u.GetUserFromContext(c)
	if err != nil {
		util.Response(c, "User not logged in", 500, "user not found", nil)
		return
	}

	transaction, err := u.Repository.GetTransactionByAccountNumber(user.AccountNo)
	if err != nil {
		util.Response(c, "Transaction not found", 400, "Bad request body", nil)
		return
	}

	util.Response(c, "Transaction found", 200, "Successful", nil)

	c.IndentedJSON(200, gin.H{"Transaction History": transaction})
}
