package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"payment-system-three/internal/api"
	"payment-system-three/internal/middleware"
	"payment-system-three/internal/ports"

	"time"
)

// SetupRouter is where router endpoints are called
func SetupRouter(handler *api.HTTPHandler, repository ports.Repository) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//POSTMAN DATA
	r := router.Group("/")
	{
		r.GET("/", handler.Readiness)
		r.POST("/create", handler.CreateUser)
		r.POST("/login", handler.LoginUser)
		r.POST("/admin/create", handler.CreateAdmin)	
		r.POST("/admin/login", handler.LoginAdmin)
	}

	//user := r.Group("/user")
	//{
	// user.POST("/create", handler.CreateUser)
	//	user.POST("/login", handler.LoginUser)
	// }

	// AuthorizeAdmin authorizes all the authorized users haldlers
	authorizeAdmin := r.Group("/admin")
	{
		//authorizeAdmin.POST("/create", handler.CreateAdmin)
	//	authorizeAdmin.POST("/login", handler.LoginAdmin)
	}
	authorizeAdmin.Use(middleware.AuthorizeAdmin(repository.FindUserByEmail, repository.TokenInBlacklist))
	{
		authorizeAdmin.GET("/user", handler.GetUserByEmail)
	}

	return router
}
