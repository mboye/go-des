package event

// Heap is a heap of events
type Heap []Event

// Len returns the length of the heap
func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return Less(h[i], h[j])
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds an item to the heap
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(Event))
}

// Pop removes an item from the heap
func (h *Heap) Pop() interface{} {
	oldHeap := *h
	heapLength := len(oldHeap)
	head := oldHeap[heapLength-1]
	*h = oldHeap[0 : heapLength-1]
	return head
}
