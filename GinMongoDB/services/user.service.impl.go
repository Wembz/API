package services

import (
	"context"
	"errors"

	"github.com/MongoDB/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx 		context.Context
}


func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl {
		usercollection: usercollection,
		ctx: 		ctx,
	}
}

//model.user haas to be similar to its type thats been created in UserService
func (u *UserServiceImpl) CreateUser(user *models.User) error {
//Logic of creating a user: InsertOne will return the context from the result & errors
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	var user *models.User
	//To make a query of a specific name/User
	query := bson.D{bson.E{Key: "user_name", Value: name}}
	//Get the User
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
 return user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	//Slice object for user
	var users []*models.User
	//How to find all the document of user
	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}})
	//Error Handling
	if err != nil {
		return nil, err
	}

	for cursor.Next(u.ctx) {
		 
		var user models.User
	//Decoding a particular user
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err 
		}
	//Adding user results to appended slice	
		users = append(users, &user)
	}
	//Error Handling this method will capture any particular error
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	//Close the cursor
	cursor.Close(u.ctx)

	//if the slice still equal to "0" meaning no documents found
	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}
	return users, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	//filter user by Name
	filter := bson.D{bson.E{Key: "user_name", Value: user.Name}}
	//update each user detail Name, Age & Address
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "user_name", Value: user.Name}, bson.E{Key: "user_age", Value: user.Age}, bson.E{Key: "user_address", Value: user.Address}}}}
	//results  with ctx, filter & update included
	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	// Error Handling
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	//filter user by Name
	filter := bson.D{bson.E{Key: "user_name", Value: name}}
	//results  with ctx, filter & deleted included
	result, _ := u.usercollection.DeleteOne(u.ctx, filter)
		// Error Handling
		if result.DeletedCount != 1 {
			return errors.New("no matched document found for delete")
		}
	return nil
}