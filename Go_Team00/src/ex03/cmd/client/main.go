package main

import (
	"client/internal/analyzer"
	"client/internal/config"
	"client/internal/database"
	"client/internal/repository"
	"client/internal/service"
	"context"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	log.Printf("Connecting to DB with config: host=%s port=%d user=%s dbname=%s", cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.DBName)
	db := database.NewDB(
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Host,
		cfg.Database.Port,
	)

	repo := repository.NewGormFrequencyRepository(*db)
	anomalyDetector := analyzer.NewAnalyzer(repo, cfg.Anomaly.STDDevMultiplier)

	frequencyService := service.NewFrequencyService(cfg, anomalyDetector)

	ctx := context.Background()
	frequencyService.StartReceivingFrequencies(ctx)
}
