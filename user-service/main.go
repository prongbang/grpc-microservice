package main

import (
	"context"
	"errors"
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	"github.com/prongbang/grpc-microservice/user-service/proto/user"
	"google.golang.org/grpc"
)

const port = ":50052"

// Server is used to implement user.UserServer
type userServer struct {
	user.UnimplementedUserServer
}

func (a *userServer) GetUser(ctx context.Context, request *user.UserRequest) (*user.UserResponse, error) {
	log.Printf("Received: username=%v", request.GetUsername())

	// Mock find user by username
	if request.GetUsername() == "admin" {
		return &user.UserResponse{
			Id:       "1",
			Name:     "Administrator",
			Username: "admin",
			Password: "1234",
		}, nil
	}

	return nil, errors.New("404")
}

func NewUserServer() user.UserServer {
	return &userServer{}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	user.RegisterUserServer(s, NewUserServer())

	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("Server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
