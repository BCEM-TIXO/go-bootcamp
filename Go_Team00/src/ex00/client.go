package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"server/transmitter"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := transmitter.NewTransmitterClient(conn)

	stream, err := client.StreamFrequencies(context.Background())
	if err != nil {
		log.Fatalf("Failed to call StreamFrequencies: %v", err)
	}

	for {
		frequency, err := stream.Recv()
		if err != nil {
			log.Fatalf("Failed to receive frequency: %v", err)
		}
		log.Printf("Received frequency: %v", frequency)
	}
}
