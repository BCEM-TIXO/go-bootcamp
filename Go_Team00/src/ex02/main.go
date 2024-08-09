package main

import (
	"client/internal/config"
	"client/internal/database"
	"client/internal/model"
	"client/internal/repository"
)

func main() {
	cfg := config.LoadConfig()

	db := database.NewDB(
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Host,
		cfg.Database.Port,
	)
	repo := repository.NewGormFrequencyRepository(*db)
	repo.Save(
		&model.FrequencyRecord{
			SessionID: "123",
			Frequency: 1000,
			Timestamp: 123,
		},
	)

}
