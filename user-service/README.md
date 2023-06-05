# user-service

## 1.Create a proto

- [user.proto](proto/user.proto)

## 2. Gen go file

```shell
make gen_user
```

Output

```
user-service
└── proto
    ├── user
    │   ├── user.pb.go
    │   └── user_grpc.pb.go
    └── user.proto
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

	"github.com/prongbang/grpc-microservice/user-service/proto/user"
	"google.golang.org/grpc"
)

const port = ":50052"

// Server is used to implement user.UserServer
type userServer struct {
	user.UnimplementedUserServer
}

func (a *userServer) Login(ctx context.Context, request *user.UserRequest) (*user.UserResponse, error) {
	log.Printf("Received: username=%v", request.GetUsername())
	
	// TODO implement logic here
	
	return nil, errors.New("404")
}

func NewUserServer() user.UserServer {
	return &userServer{}
}

func main() {
	// Set up a connection tcp server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
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
```