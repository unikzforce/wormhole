package main

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"wormhole/cmd/test_agent/generated"
	"wormhole/cmd/test_agent/impl"
)

func main() {
	log.Println("Starting grpc server 1")
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := impl.TestAgentServiceImpl{}

	log.Println("Starting grpc server 2")

	grpcServer := grpc.NewServer()

	log.Println("Starting grpc server 3")

	generated.RegisterTestAgentServiceServer(grpcServer, &s)

	log.Println("Starting grpc server 4")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc server over port 9000: %v", err)
	}
}
