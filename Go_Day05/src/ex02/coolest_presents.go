package ph

import (
	"container/heap"
	"errors"
)

func getNCoolestPresents(presents []Present, n int) ([]Present, error) {
	if n < 0 || n > len(presents) {
		return nil, errors.New("invalid n")
	}
	ph := make(PresentHeap, len(presents))
	for i, v := range presents {
		ph[i].Value = v.Value
		ph[i].Size = v.Size
	}

	heap.Init(&ph)
	coolestPresents := make([]Present, n)
	for i := 0; i < n; i++ {
		present := heap.Pop(&ph).(Present)
		coolestPresents[i] = present
	}

	return coolestPresents, nil
}
