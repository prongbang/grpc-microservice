# web-service

## 1. Copy a auth.proto file in to project

```
web-service
└── proto
    └── auth.proto
```

## 2. Gen go file

```shell
make gen_auth
```

Output

```
web-service
└── proto
    ├── auth
    │   ├── auth.pb.go
    │   └── auth_grpc.pb.go
    └── auth.proto
```

## 3. Connect to auth-service

```go
package main 

import (
	"context"
	"github.com/prongbang/grpc-microservice/web-service/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	authAddress = "localhost:50051"
)

func main() {
    // Set up a connection to the server
    authConn, authErr := grpc.Dial(authAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if authErr != nil {
        log.Fatalf("did not connect: %v", authErr)
    }
    defer func (authConn *grpc.ClientConn) { _ = authConn.Close() }(authConn)
    authClient := auth.NewAuthClient(authConn)

	// Using function Login in auth-service
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    resp, err := authClient.Login(ctx, &auth.LoginRequest{Username: "admin", Password: "1234"})
    if err != nil {
        // TODO error   
    }
}
```