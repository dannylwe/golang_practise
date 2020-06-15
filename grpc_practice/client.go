package main

import (
	"golang.org/x/net/context"
	"log"
	"time"

	pb "github.com/danny/service/gravatar"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())

	if err != nil {
		log.Fatal("could not connect to grpc server: %v", err)
	}

	defer conn.Close()

	c := pb.NewGravatarServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()

	response, err := c.Generate(ctx, &pb.GravatarRequest{Email: "dhdudy@gmail.com", Size: 10})
	if err != nil {
		log.Fatal("could not run client function 'Generate': %v", err)
	}

	log.Println("response from server:", response)
}