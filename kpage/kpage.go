package kpage

import (
	"fmt"
	"math/rand"
	"time"
)

// Solve recive a list of al the edges, how many vertexes are in the graph and the number of pages
// returns a posible solution to the problem
func Solve(edges []*Edge, maxVertexNo, pages uint) (*Solution, error) {

	// Check that all the edges are valid
	for _, e := range edges {
		if e.Src > maxVertexNo || e.Dst > maxVertexNo || e.Src == 0 || e.Dst == 0 {
			return nil, fmt.Errorf("Edge source or destination outise of vertex range")
		}

	}

	// Creates slices with one more to use the vertexes in base 1
	order := make([]uint, maxVertexNo+1)
	vpos := make([]uint, maxVertexNo+1)

	// Create the solution struct
	s := &Solution{
		Vertex:    maxVertexNo,
		Pages:     pages,
		Edges:     edges,
		Order:     order,
		vPosition: vpos,
	}

	// Change the random seed to ensure randomness
	rand.Seed(time.Now().UTC().UnixNano())

	// assign to the first position in the order a random vertex and order the other ones
	s.Order[1] = uint(rand.Intn(int(s.Vertex)) + 1)

	err := s.OrderVertexes(1)
	if err != nil {
		return nil, err
	}

	// Assign pages to all the edges of the solution
	s.AssignPages()

	return s, nil
}
