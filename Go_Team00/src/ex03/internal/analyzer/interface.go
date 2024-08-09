package analyzer

type Analyzer interface {
	ProcessFrequency(sessionID string, frequency float64, timeStamp int64)
}
