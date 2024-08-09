package analyzer

import (
	"client/internal/analyzer/detector"
	"client/internal/model"
	"client/internal/repository"
	"fmt"
)

type analyzer struct {
	repo     repository.FrequencyRepository
	detector *detector.Detector
}

func NewAnalyzer(repo repository.FrequencyRepository, k float64) Analyzer {
	return &analyzer{
		repo:     repo,
		detector: detector.NewDetector(k),
	}
}

func (a *analyzer) ProcessFrequency(sessionID string, frequency float64, timeStamp int64) {
	res := a.detector.Analyze(frequency)
	if res {
		record := model.FrequencyRecord{
			SessionID: sessionID,
			Frequency: frequency,
			Timestamp: timeStamp,
		}

		if err := a.repo.Save(&record); err != nil {
			fmt.Errorf("%w", err)
		}
	}
}

// func (a *analyzer) isAnomaly(frequency float64) bool {

// 	return a.detector.
// }
