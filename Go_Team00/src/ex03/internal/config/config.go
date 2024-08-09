package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Database          DatabaseConfig
	Anomaly           AnomalyConfig
	GRPCServerAddress string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	MaxConns int
}

type AnomalyConfig struct {
	STDDevMultiplier float64
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Invalid DB_PORT value: %v", err)
	}

	dbMaxConns, err := strconv.Atoi(os.Getenv("DB_MAX_CONNS"))
	if err != nil {
		log.Fatalf("Invalid DB_MAX_CONNS value: %v", err)
	}

	stdDevMultiplier, err := strconv.ParseFloat(os.Getenv("ANOMALY_STDDEV_MULTIPLIER"), 64)
	if err != nil {
		log.Fatalf("Invalid ANOMALY_STDDEV_MULTIPLIER value: %v", err)
	}

	return Config{
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     dbPort,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			MaxConns: dbMaxConns,
		},
		Anomaly: AnomalyConfig{
			STDDevMultiplier: stdDevMultiplier,
		},
		GRPCServerAddress: os.Getenv("GRPC_SERVER_ADDRESS"),
	}
}
