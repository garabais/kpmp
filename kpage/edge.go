package kpage

import (
	"fmt"
)

// Edge keeps the necesary information to keep track of and edge and in which page it is
type Edge struct {
	Src  uint
	Dst  uint
	Page uint
}

func (e *Edge) contains(a, b uint) bool {
	return a == e.Src || a == e.Dst || b == e.Src || b == e.Dst
}

func (e *Edge) String() string {
	return fmt.Sprintf("%v %v %v", e.Src, e.Dst, e.Page)
}

func (e *Edge) equals(a, b uint) bool {
	return a == e.Src && b == e.Dst || b == e.Src && a == e.Dst
}

// NewEdge creates an edge and return a pointer to it
func NewEdge(src, dst uint) *Edge {
	return &Edge{Src: src, Dst: dst, Page: 0}
}

func (e *Edge) otherOne(n uint) (uint, error) {
	if e.Src == n {
		return e.Dst, nil
	} else if e.Dst == n {
		return e.Src, nil
	}
	return 0, fmt.Errorf("value not in edge")
}
