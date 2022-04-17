package main

import (
	pb "github.com/gocanto/gRPC-Playground/user"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Failed to close connection: %v", err)
		}
	}(conn)

	client := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	var users = make(map[string]int32)
	users["gus"] = 37
	users["liane"] = 34

	for name, age := range users {
		item, err := client.CreateUser(ctx, &pb.NewUserRequest{
			Name: name,
			Age:  age,
		})

		if err != nil {
			log.Fatalf("Failed to create user: %v", err)
		}

		log.Printf(`
			User Details: 
				Name: %s
				Age: %d
				Id: %d`, item.GetName(), item.GetAge(), item.GetId(),
		)
	}
}
