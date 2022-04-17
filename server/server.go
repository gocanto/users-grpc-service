package main

import (
	"context"
	pb "github.com/gocanto/gRPC-Playground/user"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
)

const (
	port = ":50051"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (receiver *UserServer) CreateUser(ctx context.Context, request *pb.NewUserRequest) (*pb.User, error) {
	log.Printf("Received: %v", request.GetName())
	log.Printf("Using context: %v", ctx)

	return &pb.User{
		Id:   rand.Int31n(1000),
		Name: "Gus",
		Age:  37,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &UserServer{})

	log.Printf("Server listing at %v:", listener.Addr())

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
