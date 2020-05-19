package kpage

import (
	"fmt"
	"math/rand"
	"time"
)

// Solve recive a list of al the edges, how many vertexes are in the graph and the number of pages
// returns a posible solution to the problem
func Solve(edges []*Edge, maxVertexNo, pages uint) (*Solution, error) {

	for _, e := range edges {
		if e.Src > maxVertexNo || e.Dst > maxVertexNo {
			return nil, fmt.Errorf("Edge source or destination outise of vertex range")
		}

	}

	order := make([]uint, maxVertexNo+1)
	vpos := make([]uint, maxVertexNo+1)

	s := &Solution{
		Pages:     pages,
		Edges:     edges,
		Order:     order,
		vPosition: vpos,
	}

	rand.Seed(time.Now().UTC().UnixNano())

	s.Order[1] = uint(rand.Intn(int(maxVertexNo)) + 1)

	err := s.OrderVertexes(1)
	if err != nil {
		return nil, err
	}

	s.Crossings = s.AssignPages()

	return s, nil
}
