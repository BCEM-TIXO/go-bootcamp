package detector

import "math"

// import "fmt"

type Detector struct {
	metrics    metrics
	k          float64
	colibrated bool
}

type metrics struct {
	count      int64
	mean       float64
	dispertion float64
	std        float64
}

func NewDetector(k float64) *Detector {
	return &Detector{k: k,
		colibrated: false}
}

func (m *metrics) update(x float64) {
	m.count++
	differential := (x - m.mean) / float64(m.count)
	newMean := m.mean + differential

	newDispertion := m.dispertion + (x-newMean)*(x-m.mean)

	m.mean = newMean
	m.dispertion = newDispertion
	if m.count == 1 {
		m.std = 0
	} else {
		m.std = math.Sqrt(m.dispertion / float64(m.count-1))
	}

}

func (d *Detector) Update(frequency float64) {
	d.metrics.update(frequency)

	d.colibrated = d.metrics.count >= 75
}

func (d Detector) IsColibrated() bool {
	return d.colibrated
}

func (d Detector) CheckAnomaly(x float64) bool {
	return math.Abs(x-d.metrics.mean) > (d.metrics.std * d.k)
}

func (d *Detector) Analyze(frequency float64) bool {
	if d.IsColibrated() {
		return d.CheckAnomaly(frequency)
	} else {
		d.Update(frequency)
	}
	return false
}
