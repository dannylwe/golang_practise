package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/danny/service/gravatar"
)

const port = ":8080"

type gravatarService struct {}

func(s * gravatarService) Generate(ctx context.Context, in *pb.GravatarRequest) (*pb.GravatarResponse, error) {
	log.Printf("Recieved email %v of size %v from client", in.Email, in.Size)
	return &pb.GravatarResponse{Url: pb.Gravatar(in.Email, uint32(in.Size))}, nil
}

func main() {
	fmt.Println("starting application")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen on port: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterGravatarServiceServer(server, &gravatarService{})

	if err := server.Serve(lis); err != nil {
		log.Fatal("Failed to start server: %v", err)
	}
}