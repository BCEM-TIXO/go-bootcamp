package main

import "fmt"
import "os"
import "bufio"
import "strconv"
import "sort"
import "math"
import "flag"
import "errors"

type Solution struct {
	numbers []int
	dense   map[int]int

	median float64
	mean   float64
	mode   int
	sd     float64
}

func (v *Solution) getNumbers() error {
	if v.dense == nil {
		v.dense = make(map[int]int)
	}
	var kok error
	kok = nil
	fileScanner := bufio.NewScanner(os.Stdin)

	fileScanner.Split(bufio.ScanLines)
	sum := 0
	for fileScanner.Scan() {
		text_num := fileScanner.Text()
		if len(text_num) == 0 {
			break
		}
		num, kok := strconv.Atoi(text_num)
		if kok == nil {
			v.dense[num]++
			sum += num
			v.numbers = append(v.numbers, num)
		} else {
			return kok
		}
	}
	l := len(v.numbers)
	if l == 0 {
		kok = errors.New("zero len")
		return kok
	}
	sort.Ints(v.numbers)

	max_count := 0
	mode := 0
	for num, count := range v.dense {
		if count > max_count || (count == max_count && num < mode) {
			max_count = count
			mode = num
		}
	}
	var median float64
	if l%2 == 0 {
		median = float64(v.numbers[l/2-1]+v.numbers[l/2]) / 2.0
	} else {
		median = float64(v.numbers[l/2])
	}
	mean := float64(sum) / float64(len(v.numbers))
	sigma_sum := 0.
	for key, val := range v.dense {
		sigma := (float64(key) - mean) * (float64(key) - mean) * float64(val)
		sigma_sum += sigma
	}
	v.mean = mean
	v.median = median
	v.mode = mode
	v.sd = math.Sqrt(sigma_sum / float64(len(v.numbers)))
	return kok
}

func (v Solution) printSol(kok error, median, mean, sd, mode bool) {
	if kok != nil {
		fmt.Printf("Error: %v\n", kok)
		return
	}
	if !median {
		fmt.Printf("Median: %.2g\n", v.median)
	}
	if !mean {
		fmt.Printf("Mean: %.2g\n", v.mean)
	}
	if !sd {
		fmt.Printf("SD: %.2g\n", v.sd)
	}
	if !mode {
		fmt.Printf("Mode: %d\n", v.mode)
	}
}

func main() {

	medianPtr := flag.Bool("median", false, "not print median value")
	meanPtr := flag.Bool("mean", false, "not print mean value")
	sdPtr := flag.Bool("sd", false, "not print sd value")
	modePtr := flag.Bool("mode", false, "not print mode value")
	flag.Parse()

	var tmp Solution
	kok := tmp.getNumbers()
	tmp.printSol(kok, *medianPtr, *meanPtr, *sdPtr, *modePtr)

}
