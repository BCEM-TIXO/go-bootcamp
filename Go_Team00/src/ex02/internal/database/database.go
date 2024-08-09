package database

import (
	"client/internal/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDB(user, password, dbname, host string, port int) *Database {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&model.FrequencyRecord{})

	return &Database{db: db}
}

func (d *Database) CreateFrequencyRecord(record model.FrequencyRecord) error {
	result := d.db.Create(&record)
	return result.Error
}
