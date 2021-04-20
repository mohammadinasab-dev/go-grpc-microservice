package grpcserver

import (
	"context"

	"github.com/mohammadinasab-dev/grpc/cmd"
	"github.com/mohammadinasab-dev/grpc/configuration"
	"github.com/mohammadinasab-dev/grpc/data"
)

type Grpcserver struct {
	dbHandler *data.SQLHandler
	cmd.UnimplementedUserServiceServer
}

func NewGrpcServer(config configuration.Config) (*Grpcserver, error) { //"user:1234@/people"
	db, err := data.CreateDBConnection(config)
	if err != nil {
		return nil, err
	}
	return &Grpcserver{
		dbHandler: db,
	}, err
}

func (server *Grpcserver) GetUser(ctx context.Context, r *cmd.Request) (*cmd.User, error) {
	user, err := server.dbHandler.DBGetUserByID(r.GetName())
	if err != nil {
		return nil, err
	}
	return convertToGrpcUser(user), nil
}

func (server *Grpcserver) GetUsers(r *cmd.Request, stream cmd.UserService_GetUsersServer) error {
	users, err := server.dbHandler.DBGetUsers()
	if err != nil {
		return err
	}
	for _, user := range users {
		grpcUser := convertToGrpcUser(user)
		err := stream.Send(grpcUser)
		if err != nil {
			return err
		}
	}
	return nil
}

func convertToGrpcUser(user data.User) *cmd.User {
	return &cmd.User{
		UserID:   int32(user.UserID),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

// func (server *Grpcserver) mustEmbedUnimplementedUserServiceServer() {
// 	log.Println("mustEmbedUnimplementedUserServiceServer method implemented")
// }
