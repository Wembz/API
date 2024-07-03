package api

import (
	"os"
	"payment-system-three/internal/middleware"
	"payment-system-three/internal/models"
	"payment-system-three/internal/util"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//Create Admin

func (u *HTTPHandler) CreateAdmin(c *gin.Context) {
	// Var Struct for  Admin
	var admin *models.Admin

	// Error message request for wrong entry
	if err := c.ShouldBind(&admin); err != nil {
		util.Response(c, "invalid request", 400, err.Error(), nil)
		return
	}

	// Admin entry details
	admin.FirstName = strings.TrimSpace(admin.FirstName)
	admin.LastName = strings.TrimSpace(admin.LastName)
	admin.Email = strings.TrimSpace(admin.Email)
	admin.Password = strings.TrimSpace(admin.Password)
	admin.DateOfBirth = strings.TrimSpace(admin.DateOfBirth)
	admin.Phone = strings.TrimSpace(admin.Phone)
	admin.Address = strings.TrimSpace(admin.Address)

	// If admin entry is empty
	if admin.FirstName == "" {
		util.Response(c, "First name must not be empty", 400, nil, nil)
		return
	}
	if admin.LastName == "" {
		util.Response(c, "Last name must not be empty", 400, nil, nil)
		return
	}
	if admin.Email == "" {
		util.Response(c, "Email must not be empty", 400, nil, nil)
		return
	}
	if admin.Password == "" {
		util.Response(c, "Password must not be empty", 400, nil, nil)
		return
	}
	if admin.DateOfBirth == "" {
		util.Response(c, "Date of birth must not be empty", 400, nil, nil)
		return
	}
	if admin.Phone == "" {
		util.Response(c, "Phone must not be empty", 400, nil, nil)
		return
	}
	if admin.Address == "" {
		util.Response(c, "Address must not be empty", 400, nil, nil)
		return
	}

	// If Email Already exist
	isEmailExist, _ := u.Repository.FindAdminByEmail(admin.Email)

	// Error message if Email Already exist
	if isEmailExist != nil {
		util.Response(c, "Email already exist", 400, nil, nil)
		return
	}

	// validate Email
	if !util.ValidatePassword(admin.Password) {
		util.Response(c, "Password acceptence criteria not matched. Password must be At least 6 characters long , Contains at least one uppercase letter, Contains at least one number, Contains at least one special character", 400, nil, nil)
		return
	}

	// Hash the password
	hashedPassword, err := util.HashPassword(admin.Password)
	if err != nil {
		util.Response(c, "Internal server error", 500, err.Error(), nil)
		return
	}
	admin.Password = hashedPassword

	err = u.Repository.CreateAdmin(admin)
	if err != nil {
		util.Response(c, "Admin not created", 400, err.Error(), nil)
		return
	}
	util.Response(c, "Admin created", 200, nil, nil)

}

// Login Admin
func (u *HTTPHandler) LoginAdmin(c *gin.Context) {
	var loginRequest *models.LoginRequestAdmin
	err := c.ShouldBind(&loginRequest)
	if err != nil {
		util.Response(c, "invalid request", 400, err.Error(), nil)
		return
	}

	loginRequest.Email = strings.TrimSpace(loginRequest.Email)
	loginRequest.Password = strings.TrimSpace(loginRequest.Password)

	if loginRequest.Email == "" {
		util.Response(c, "Email must not be empty", 400, nil, nil)
		return
	}
	if loginRequest.Password == "" {
		util.Response(c, "Password must not be empty", 400, nil, nil)
		return
	}

	admin, err := u.Repository.FindAdminByEmail(loginRequest.Email)
	if err != nil {
		util.Response(c, "Email does not exist", 404, err.Error(), nil)
		return
	}

	// Verify the password
	match := util.CheckPasswordHash(loginRequest.Password, admin.Password)
	if !match {
		admin.UpdatedAt = time.Now()
		err = u.Repository.UpdateAdmin(admin)
		if err != nil {
			util.Response(c, "There is an error occured", 500, err.Error(), nil)
			return
		}
		util.Response(c, "Incorrect password", 401, nil, nil)
		return
	}

	accessClaims, refreshClaims := middleware.GenerateClaims(admin.Email)

	secret := os.Getenv("JWT_SECRET")

	accessToken, err := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)
	if err != nil {
		util.Response(c, "Error generating access token", 500, err.Error(), nil)
		return
	}

	refreshToken, err := middleware.GenerateToken(jwt.SigningMethodHS256, refreshClaims, &secret)
	if err != nil {
		util.Response(c, "Error generating refresh token", 500, err.Error(), nil)
		return
	}

	c.Header("access_token", *accessToken)
	c.Header("refresh_token", *refreshToken)

	util.Response(c, "Login successful", 200, gin.H{
		"admin":         admin,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil)
}

// Protected Route
func (u *HTTPHandler) GetAdminByEmail(c *gin.Context) {

	_, err := u.GetAdminFromContext(c)
	if err != nil {
		util.Response(c, "Admin not logged in", 500, err.Error(), nil)
		return
	}

	email := c.Query("email")
	email = strings.TrimSpace(email)
	if email == "" {
		util.Response(c, "Email is required", 400, nil, nil)
		return
	}

	admin, err := u.Repository.FindAdminByEmail(email)
	if err != nil {
		util.Response(c, "Admin not found", 404, err.Error(), nil)
		return
	}

	util.Response(c, "Admin Found", 200, admin, nil)

}



