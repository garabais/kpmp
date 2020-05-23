package kpage

// heap is used to make a heap in the solution edges
type heap struct {
	s        *Solution // Pointer to the solution to acces the *edges slice and the position of the vertexes
	heapSize int       // Current size of the heap
}

func buildHeap(sol *Solution) heap {
	h := heap{s: sol, heapSize: len(sol.Edges)}
	for i := 1 + uint(len(sol.Edges))/2; i > 0; i-- {
		h.heapify(i - 1)
	}
	return h
}

func (h *heap) heapify(i uint) {
	var l, r, max uint

	l, r = 2*i+1, 2*i+2
	max = i

	// This compares the length of the edge using the position in the solution to find th bigger one
	// This is normaly to find the bigger one but in this case we want the edges order from biggest lo smallest
	if l < h.size() && h.s.getEdgeLength(l) < h.s.getEdgeLength(max) {
		max = l
	}
	if r < h.size() && h.s.getEdgeLength(r) < h.s.getEdgeLength(max) {
		max = r
	}

	if max != i {
		h.s.Edges[i], h.s.Edges[max] = h.s.Edges[max], h.s.Edges[i]
		h.heapify(max)
	}
}

func (h *heap) size() uint { return uint(h.heapSize) }

func heapSort(slice *Solution) {
	h := buildHeap(slice)

	for i := len(h.s.Edges) - 1; i >= 1; i-- {
		h.s.Edges[0], h.s.Edges[i] = h.s.Edges[i], h.s.Edges[0]
		h.heapSize--
		h.heapify(0)
	}

}
