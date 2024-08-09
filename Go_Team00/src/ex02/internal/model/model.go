package model

import (
	"gorm.io/gorm"
)

type FrequencyRecord struct {
	gorm.Model
	SessionID string  `gorm:"type:varchar(100);not null"`
	Frequency float64 `gorm:"not null"`
	Timestamp int64   `gorm:"index"`
}
