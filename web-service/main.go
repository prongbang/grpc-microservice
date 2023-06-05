package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/prongbang/grpc-microservice/web-service/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	authAddress = "localhost:50051"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	// Set up a connection to the server
	authConn, authErr := grpc.Dial(authAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if authErr != nil {
		log.Fatalf("did not connect: %v", authErr)
	}
	defer func(authConn *grpc.ClientConn) { _ = authConn.Close() }(authConn)
	authClient := auth.NewAuthClient(authConn)

	// New creates a new Fiber named instance.
	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	// Routers
	v1 := app.Group("/v1")
	{
		v1.Post("/login", func(c *fiber.Ctx) error {
			data := UserCredentials{}
			if err := c.BodyParser(&data); err != nil {
				return fiber.ErrBadRequest
			}
			log.Println(data)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			resp, err := authClient.Login(ctx, &auth.LoginRequest{Username: data.Username, Password: data.Password})
			if err != nil {
				return c.JSON(fiber.Map{"message": err.Error()})
			}
			return c.JSON(resp)
		})
	}

	log.Fatal(app.Listen(":8000"))
}
