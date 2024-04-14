package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"wormhole/cmd/test_agent/proto"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Could not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewChatServiceClient(conn)

	message := proto.Message{
		Body: "Hello from the client",
	}

	resp, err := c.SayHello(context.Background(), &message)

	if err != nil {
		log.Fatalf("error while calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", resp.Body)

}
