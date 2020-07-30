package main

import (
	"log"
	"net"

	"github.com/sangianpatrick/grpc-service-demo/service"

	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()

	// initiate service
	service.InitializeUser(server)

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(server.Serve(listener))
}
