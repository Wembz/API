package controllers

import (
	"net/http"

	"github.com/MongoDB/models"
	"github.com/MongoDB/services"
	"github.com/gin-gonic/gin"
)

// Service section will interact with the database service
type UserController struct {
	UserService services.UserService
}

func New(userService services.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}


//ctx.JSON holds the details from the request the user sends
func (uc *UserController) CreateUser(ctx *gin.Context)  {
	//user variable 
	var user models.User

	//Binding ctx with the user variable we created
	if err := ctx.ShouldBindJSON(&user); err != nil {
	
	//Error Handling	
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//Calling the CreateUser function
	err := uc.UserService.CreateUser(&user)
	//Error Handling
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) GetUser(ctx *gin.Context)  {

	username := ctx.Param("name")
	//Get user
	user, err := uc.UserService.GetUser(&username)
	//Error Handling
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAll(ctx *gin.Context)  {
	users, err := uc.UserService.GetAll()
	//Error Handling
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	//Binding ctx with the user variable we created
	if err := ctx.ShouldBindJSON(&user); err != nil {
	
	//Error Handling	
	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	return
	}
	err := uc.UserService.UpdateUser(&user)
	//Error Handling
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("name")
	// Delete User function
	err := uc.UserService.DeleteUser(&username)
	//Error Handling
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	//Success Message
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

//User Routes
func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/create", uc.CreateUser)
	userroute.GET("/get/:name", uc.GetUser)
	userroute.GET("/getall", uc.GetAll)
	userroute.PATCH("/update", uc.UpdateUser)
	userroute.DELETE("/delete/:name", uc.DeleteUser)

}