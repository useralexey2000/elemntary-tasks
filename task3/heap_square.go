package main

type TrSquare struct {
	Name   string
	Square float64
}

type TrSquaresHeap []*TrSquare

func (h TrSquaresHeap) Len() int {
	return len(h)
}

func (h TrSquaresHeap) Less(i, j int) bool {
	return h[i].Square > h[j].Square
}

func (h TrSquaresHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *TrSquaresHeap) Push(x interface{}) {
	*h = append(*h, x.(*TrSquare))
}

func (h *TrSquaresHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
