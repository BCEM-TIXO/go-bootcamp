package repository

import (
	"client/internal/database"
	"client/internal/model"
)

type gormFrequencyRepository struct {
	db database.Database
}

func NewGormFrequencyRepository(db database.Database) FrequencyRepository {
	return &gormFrequencyRepository{db: db}
}

func (r *gormFrequencyRepository) Save(record *model.FrequencyRecord) error {
	anomaly := model.FrequencyRecord{
		SessionID: record.SessionID,
		Frequency: record.Frequency,
		Timestamp: record.Timestamp,
	}

	r.db.CreateFrequencyRecord(anomaly)

	return nil
}
