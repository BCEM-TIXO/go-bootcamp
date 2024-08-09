package repository

import "client/internal/model"

type FrequencyRepository interface {
	Save(record *model.FrequencyRecord) error
}
