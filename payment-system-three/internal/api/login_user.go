package api

import (
	"net/http"
	"os"
	"payment-system-three/internal/middleware"
	"payment-system-three/internal/models"
	"payment-system-three/internal/util"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// login User
func (u *HTTPHandler) LoginUser(c *gin.Context) {
	var loginRequest *models.LoginRequestUser
	if err := c.ShouldBind(&loginRequest); err != nil {
		util.Response(c, "invalid request", 400, "bad request body", nil)
		return
	}
	if loginRequest.Email == "" || loginRequest.Password == "" {
		util.Response(c, "Please enter your email or password", 400, "bad request body", nil)
		return
	}
	if loginRequest.Password == "" {
		util.Response(c, "Password must not be empty", 400, nil, nil)
		return
	}
	user, err := u.Repository.FindUserByEmail(loginRequest.Email)
	if err != nil {
		util.Response(c, "user does not exist", 404, "user not found", nil)
		return
	}

	// Verify Password
	match := util.CheckPasswordHash(loginRequest.Password, user.Password)
	if !match {
		user.LoginCounter++
		user.UpdatedAt = time.Now()
		err = u.Repository.UpdateUser(user)
		if err != nil {
			util.Response(c, "There is an error occured", 500, err.Error(), nil)
			return
		}
		util.Response(c, "Incorrect Password", 401, nil, nil)
		return
	}
	//Generate token
	accessClaims, refreshClaims := middleware.GenerateClaims(user.Email)

	secret := os.Getenv("JWT_SECRET")

	accessToken, err := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)
	if err != nil {
		util.Response(c, "error generating access token", 500, "error generating access token", nil)
		return
	}
	refreshToken, err := middleware.GenerateToken(jwt.SigningMethodHS256, refreshClaims, &secret)
	if err != nil {
		util.Response(c, "error generating refresh token", 500, "error generating refresh token", nil)
		return
	}
	c.Header("access_token", *accessToken)
	c.Header("refresh_token", *refreshToken)

	util.Response(c, "login successful", http.StatusOK, gin.H{
		"user":          user,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil)
}

// call a protected route
func (u *HTTPHandler) GetUserByEmail(c *gin.Context) {
	_, err := u.GetUserFromContext(c)
	if err != nil {
		util.Response(c, "User not logged in", 500, "user not found", nil)
		return
	}

	email := c.Query("email")

	if email == "" {
		util.Response(c, "email is required", 400, "email is required", nil)
		return
	}

	user, err := u.Repository.FindUserByEmail(email)
	if err != nil {
		util.Response(c, "user not fount", 500, "user not found", nil)
		return
	}

	util.Response(c, "user found", 200, user, nil)
}
