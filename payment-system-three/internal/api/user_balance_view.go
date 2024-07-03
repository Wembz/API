package api

import (
	"payment-system-three/internal/util"

	"github.com/gin-gonic/gin"
)

// View Balance
func (u *HTTPHandler) ViewUserBalance(c *gin.Context) {

	// Get user from context
	user, err := u.GetUserFromContext(c)
	if err != nil {
		util.Response(c, "User not logged in", 500, "user not found", nil)
		return
	}

	// checking balance
	util.Response(c, "Balance retrieved successfully", 200, "sucess", nil)
	c.IndentedJSON(200, gin.H{"balance": user.AvailableBalance})

}
