package ph

type Present struct {
	Value int
	Size  int
}

type PresentHeap []Present

func (h PresentHeap) Len() int {
	return len(h)
}

func (h PresentHeap) Less(i, j int) bool {
	if h[i].Value == h[j].Value {
		return h[i].Size < h[j].Size
	}
	return h[i].Value > h[j].Value
}
func (h PresentHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PresentHeap) Push(x interface{}) {
	present := x.(Present)
	*h = append(*h, present)
}

func (h *PresentHeap) Pop() interface{} {
	old := *h
	n := len(old)
	present := old[n-1]
	*h = old[0 : n-1]
	return present
}
