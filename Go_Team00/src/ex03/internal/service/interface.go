package service

import "context"

type FrequencyService interface {
	StartReceivingFrequencies(ctx context.Context)
}
