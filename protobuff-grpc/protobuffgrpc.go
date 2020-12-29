package protobuff_grpc

import (
	"context"
	"errors"
	"log"
	"net"
	"net/mail"

	"github.com/xans-me/protobuff-json-benchmark/protobuff-grpc/proto"
	"google.golang.org/grpc"
)

// Server type
type Server struct{}

// CreateUser handler
func (s *Server) CreateUser(ctx context.Context, user *proto.User) (*proto.Response, error) {
	validationErr := validate(user)
	if validationErr != nil {
		return &proto.Response{
			Code:    500,
			Message: validationErr.Error(),
		}, validationErr
	}

	user.Id = "1000000"
	return &proto.Response{
		Code:    200,
		Message: "OK",
		User:    user,
	}, nil
}

func validate(user *proto.User) error {
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return err
	}

	if len(user.Name) < 4 {
		return errors.New("Name is too short")
	}

	if len(user.Password) < 4 {
		return errors.New("Password is too weak")
	}

	return nil
}

// Start entrypoint
func Start() {
	listener, _ := net.Listen("tcp", ":60000")

	srv := grpc.NewServer()
	proto.RegisterAPIServer(srv, &Server{})
	log.Println(srv.Serve(listener))
}
