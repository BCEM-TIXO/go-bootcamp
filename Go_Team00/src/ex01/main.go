package main

import "detector/detector"
import "math/rand"
import "fmt"
import "log"

func main() {
	detector := detector.NewDetector(1)

	values := make([]float64, 100)
	fmt.Printf("%s\n", "Values:")
	for i := range values {
		values[i] = rand.Float64() * 10
		fmt.Printf("%.4f ", values[i])
	}
	fmt.Println()
	for _, v := range values {
		if detector.Analyze(v) {
			log.Printf("Found anomaly: %.4f", v)
		}
	}
}
