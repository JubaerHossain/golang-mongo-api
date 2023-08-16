package gapi

import (
	"github.com/JubaerHossain/golang-mongodb-api/config"
	"github.com/JubaerHossain/golang-mongodb-api/pb"
	"github.com/JubaerHossain/golang-mongodb-api/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	config         config.Config
	userService    services.UserService
	userCollection *mongo.Collection
}

func NewGrpcUserServer(config config.Config, userService services.UserService, userCollection *mongo.Collection) (*UserServer, error) {
	userServer := &UserServer{
		config:         config,
		userService:    userService,
		userCollection: userCollection,
	}

	return userServer, nil
}
