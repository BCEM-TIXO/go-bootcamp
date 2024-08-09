package service

import (
	"client/internal/analyzer"
	"client/internal/config"
	"client/transmitter"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
)

type frequencyServiceImpl struct {
	config          config.Config
	anomalyDetector analyzer.Analyzer
}

func NewFrequencyService(cfg config.Config, anomalyDetector analyzer.Analyzer) FrequencyService {
	return &frequencyServiceImpl{
		config:          cfg,
		anomalyDetector: anomalyDetector,
	}
}

func (s *frequencyServiceImpl) StartReceivingFrequencies(ctx context.Context) {
	conn, err := grpc.Dial(s.config.GRPCServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := transmitter.NewTransmitterClient(conn)
	stream, err := client.StreamFrequencies(ctx)
	if err != nil {
		log.Fatalf("Failed to call StreamFrequencies: %v", err)
	}

	for {
		frequencyData, err := stream.Recv()
		if err == io.EOF {
			log.Println("Stream has ended")
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive frequency: %v", err)
		}

		sessionID := frequencyData.GetSessionId()
		frequency := frequencyData.GetFrequency()
		timeStamp := frequencyData.GetTimestamp()

		log.Printf("Received frequency: %v for session: %s", frequency, sessionID)

		s.anomalyDetector.ProcessFrequency(sessionID, frequency, timeStamp)
	}
}
