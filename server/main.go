package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "main.go/proto"
)

const (
	// Port for gRPC server to listen to
	PORT = ":5051"
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

func main() {
	lis, err := net.Listen("tcp", PORT)

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
