package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"

	pb "main.go/proto"

	"google.golang.org/grpc"
)

// loadDependencies
func loadDependencies() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Smt wrong .env")
	}

	clientHost := os.Getenv("CL_HOST")
	clientPort := os.Getenv("CL_PORT")

	clientAddress := fmt.Sprintf("%s:%s", clientHost, clientPort)

	return clientAddress
}

func main() {
	clientAddress := loadDependencies()

	conn, err := grpc.Dial(clientAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}

	defer conn.Close()

	c := pb.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	userData := User{
		Username: "John",
		Email:    "awesomeJohn@gmail.com",
		Password: "QWERTY",
	}

	res, err := c.RegisterUser(ctx, &pb.User{Username: userData.Username, Email: userData.Email, HashedPassword: userData.Password})

	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}

	log.Printf(`
           Success : %b
           Message : %s
       `, res.Success, res.Message)

}
