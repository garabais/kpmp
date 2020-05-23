package kpage

import (
	"fmt"
)

// Edge keeps the necesary information to keep track of and edge and in which page it is
type Edge struct {
	Src  uint // Source vertex of the edge
	Dst  uint // Destination vertex of the edge
	Page uint // Page where the edge is assign
}

// String implements stringer interface to print and edge using fmt
func (e *Edge) String() string {
	return fmt.Sprintf("%v %v %v", e.Src, e.Dst, e.Page)
}

// NewEdge creates an edge and return a pointer to it
func NewEdge(src, dst uint) *Edge {
	return &Edge{Src: src, Dst: dst, Page: 0}
}

// otherOne recive one of the a number of vertex and if the given vertex is part of the edge returns the other vertex
func (e *Edge) otherOne(n uint) uint {
	if e.Src == n {
		return e.Dst
	}
	return e.Src
}
