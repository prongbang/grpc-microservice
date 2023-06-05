# auth-service

## 1.Create a proto

- [auth.proto](proto/auth.proto)

## 2. Gen go file

```shell
make gen_auth
```

Output

```
auth-service
└── proto
    ├── auth
    │   ├── auth.pb.go
    │   └── auth_grpc.pb.go
    └── auth.proto
```

## 3. Create gRPC server

- [main.go](main.go)

```go
package main

import (
	"context"
	"errors"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	"github.com/prongbang/grpc-microservice/auth-service/proto/auth"
	"google.golang.org/grpc"
)

const port = ":50051"

// Server is used to implement auth.AuthServer
type authServer struct {
	auth.UnimplementedAuthServer
}

func (a *authServer) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	log.Printf("Received: username=%v, password=%v", request.GetUsername(), request.GetPassword())
	
	// TODO implement logic here
	
	return nil, errors.New("401")
}

func NewAuthServer() auth.AuthServer {
	return &authServer{}
}

func main() {
	// Set up a connection tcp server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	auth.RegisterAuthServer(s, NewAuthServer())

	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("Server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
```