package main

import (
	"context"
	"fmt"
	"log"

	"github.com/MongoDB/controllers"
	"github.com/MongoDB/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	userService    services.UserService
	UserController controllers.UserController
	ctx            context.Context
	usercollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo connection has been established...")

	usercollection = mongoclient.Database("userdb").Collection("users")
	userService = services.NewUserService(usercollection, ctx)
	UserController = controllers.New(userService)
	server = gin.Default()

}
func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	UserController.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":9090"))
}
