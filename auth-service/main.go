package main

import (
	"context"
	"errors"
	"github.com/prongbang/grpc-microservice/auth-service/proto/user"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	"github.com/prongbang/grpc-microservice/auth-service/proto/auth"
	"google.golang.org/grpc"
)

const port = ":50051"

const (
	userAddress = "localhost:50052"
)

// Server is used to implement auth.AuthServer
type authServer struct {
	auth.UnimplementedAuthServer
	UserClient user.UserClient
}

func (a *authServer) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	log.Printf("Received: username=%v, password=%v", request.GetUsername(), request.GetPassword())

	// Find user from user-service
	resp, err := a.UserClient.GetUser(ctx, &user.UserRequest{Username: request.GetUsername()})
	if err != nil {
		return nil, errors.New("401")
	}

	if request.GetUsername() == resp.Username && request.GetPassword() == resp.Password {
		return &auth.LoginResponse{
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJlbSIsIm5hbWUiOiJkZXYgZGF5IiwiaWF0IjoxNTE2MjM5MDIyfQ.yNC-7RUVZCveMOANZcT7KWMczVkb_T7KnHv3fmMLiCI",
		}, nil
	}

	return nil, errors.New("401")
}

func NewAuthServer(userClient user.UserClient) auth.AuthServer {
	return &authServer{
		UserClient: userClient,
	}
}

func main() {
	// Set up a connection to the user-service
	userConn, userErr := grpc.Dial(userAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if userErr != nil {
		log.Fatalf("Did not connect: %v", userErr)
	}
	defer func(userConn *grpc.ClientConn) { _ = userConn.Close() }(userConn)
	userClient := user.NewUserClient(userConn)

	// Set up a connection tcp server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	auth.RegisterAuthServer(s, NewAuthServer(userClient))

	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("Server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
