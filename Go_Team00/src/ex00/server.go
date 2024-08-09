package main

import (
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"server/transmitter"
	"time"
)

type server struct {
	transmitter.UnimplementedTransmitterServer
}

func (*server) mustEmbedUnimplementedTransmitterServer() {}

func genRandInRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func (s *server) StreamFrequencies(stream transmitter.Transmitter_StreamFrequenciesServer) error {
	maxMean := 10.
	minMean := -10.
	maxSTD := 1.5
	minSTD := 0.3
	mean := genRandInRange(minMean, maxMean)
	stdDev := genRandInRange(minSTD, maxSTD)
	sessionID := uuid.New().String()
	for {

		frequency := rand.NormFloat64()*stdDev + mean
		timestamp := time.Now().Unix()

		freq := &transmitter.Frequency{
			SessionId: sessionID,
			Frequency: frequency,
			Timestamp: timestamp,
		}

		if err := stream.Send(freq); err != nil {
			log.Printf("Error sending frequency: %v", err)
			return err
		}

		time.Sleep(time.Second)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	transmitter.RegisterTransmitterServer(s, &server{})

	log.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
