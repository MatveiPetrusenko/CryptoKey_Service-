package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	pb "main.go/proto"
)

type passwordStorageSystem struct {
	pb.UnimplementedAuthServiceServer
}

func (s *passwordStorageSystem) RegisterUser(ctx context.Context, in *pb.User) (*pb.Result, error) {
	log.Printf("Received: %s %s %s", in.GetUsername(), in.GetEmail(), in.GetHashedPassword())

	r := &pb.Result{
		Success: true,
		Message: "User " + in.GetUsername() + " registered",
	}

	return r, nil
}

// loadDependencies ...
func loadDependencies() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Smt wrong .env")
	}

	serverHost := os.Getenv("SE_HOST")
	serverPort := os.Getenv("SE_PORT")

	serverAddress := fmt.Sprintf("%s:%s", serverHost, serverPort)

	return serverAddress
}

func main() {
	serverAddress := loadDependencies()

	lis, err := net.Listen("tcp", serverAddress)

	if err != nil {
		log.Fatalf("failed connection: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterAuthServiceServer(s, &passwordStorageSystem{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
