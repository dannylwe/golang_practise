package main

import (
	"fmt"
	"net"

	protos "github.com/danny/grpc/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("starting application")
	log := hclog.Default()
	cs := server.NewCurrency(log)
	gs := grpc.NewServer(gs)

	protos.RegisterCurrencyService(gs, cs)

	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Error(err)
	}
	gs.Serve(l)
}