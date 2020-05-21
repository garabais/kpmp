package kpage

import "math"

type heap struct {
	s        *Solution
	heapSize int
}

func buildHeap(sol *Solution) heap {
	h := heap{s: sol, heapSize: len(sol.Edges)}
	for i := len(sol.Edges) / 2; i >= 0; i-- {
		h.heapify(i)
	}
	return h
}

func (h heap) heapify(i int) {
	l, r := 2*i+1, 2*i+2
	max := i

	if l < h.size() && math.Abs(float64(h.s.vPosition[h.s.Edges[l].Src]-h.s.vPosition[h.s.Edges[l].Dst])) > math.Abs(float64(h.s.vPosition[h.s.Edges[max].Src]-h.s.vPosition[h.s.Edges[max].Dst])) {
		max = l
	}
	if r < h.size() && math.Abs(float64(h.s.vPosition[h.s.Edges[r].Src]-h.s.vPosition[h.s.Edges[r].Dst])) > math.Abs(float64(h.s.vPosition[h.s.Edges[max].Src]-h.s.vPosition[h.s.Edges[max].Dst])) {
		max = r
	}
	//log.Printf("MaxHeapify(%v): l,r=%v,%v; max=%v\t%v\n", i, l, r, max, h.slice)
	if max != i {
		h.s.Edges[i], h.s.Edges[max] = h.s.Edges[max], h.s.Edges[i]
		h.heapify(max)
	}
}

func (h heap) size() int { return h.heapSize }

func heapSort(slice *Solution) {
	h := buildHeap(slice)

	for i := len(h.s.Edges) - 1; i >= 1; i-- {
		h.s.Edges[0], h.s.Edges[i] = h.s.Edges[i], h.s.Edges[0]
		h.heapSize--
		h.heapify(0)
	}

}
