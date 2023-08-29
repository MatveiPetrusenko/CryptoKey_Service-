package main

import (
	"context"
	"log"
	"time"

	pb "main.go/proto"

	"google.golang.org/grpc"
)

const (
	ADDRESS = "localhost:50051"
)

type User struct {
	ID       string
	Username string
	Email    string
	Password string
}

func main() {
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}

	defer conn.Close()

	c := pb.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	userData := []User{
		{ID: "1", Username: "John", Email: "awesomeJohn@gmail.com", Password: "QWERTY"},
	}

	for _, todo := range userData {
		res, err := c.RegisterUser(ctx, &pb.User{Id: todo.ID, Username: todo.Username, Email: todo.Email, HashedPassword: todo.Password})

		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}

		log.Printf(`
           Success : %b
           Message : %s
       `, res.Success, res.Message)
	}

}
