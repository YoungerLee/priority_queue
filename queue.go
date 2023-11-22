package priority_queue

type LessFunc[T any] func(T, T) bool

type PriorityQueue[T any] struct {
	elems    []T
	lessFunc LessFunc[T]
}

func NewPriorityQueue[T any](elems []T, less LessFunc[T]) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{
		elems:    elems,
		lessFunc: less,
	}
	pq.init()
	return pq
}

func (pq *PriorityQueue[T]) Len() int {
	return len(pq.elems)
}

func (pq *PriorityQueue[T]) less(i, j int) bool {
	return pq.lessFunc(pq.elems[i], pq.elems[j])
}

func (pq *PriorityQueue[T]) swap(i, j int) {
	pq.elems[i], pq.elems[j] = pq.elems[j], pq.elems[i]
}

func (pq *PriorityQueue[T]) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !pq.less(j, i) {
			break
		}
		pq.swap(i, j)
		j = i
	}
}

func (pq *PriorityQueue[T]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && pq.less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !pq.less(j, i) {
			break
		}
		pq.swap(i, j)
		i = j
	}
	return i > i0
}

func (pq *PriorityQueue[T]) init() {
	// heapify
	n := pq.Len()
	if n == 0 {
		return
	}
	for i := n/2 - 1; i >= 0; i-- {
		pq.down(i, n)
	}
}

func (pq *PriorityQueue[T]) Push(v T) {
	pq.elems = append(pq.elems, v)
	pq.up(pq.Len() - 1)
}

func (pq *PriorityQueue[T]) Pop() T {
	var t T
	if len(pq.elems) == 0 {
		return t
	}
	n := pq.Len() - 1
	pq.swap(0, n)
	pq.down(0, n)

	old := pq.elems
	n = len(old)
	elem := old[n-1]
	old[n-1] = t
	pq.elems = old[:n-1]
	return elem
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func (pq *PriorityQueue[T]) Remove(i int) T {
	n := pq.Len() - 1
	if n != i {
		pq.swap(i, n)
		if !pq.down(i, n) {
			pq.up(i)
		}
	}
	return pq.Pop()
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func (pq *PriorityQueue[T]) Fix(i int) {
	if !pq.down(i, pq.Len()) {
		pq.up(i)
	}
}
