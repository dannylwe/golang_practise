package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/danny/service/chat"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatal("could not connect to grpc server: %v", err)
	}

	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from the client!!",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatal("Error in calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", response.Body)
}