package kpage

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"

	"github.com/mohae/deepcopy"
)

// Solution represents a solution of a kPage problem
type Solution struct {
	Pages     uint
	Crossings uint
	Vertex    uint
	Edges     []*Edge
	Order     []uint
	vPosition []uint
	adj       []*list.List
}

// BuildAdj creates and fill the adjacency matrix in a pseudo-random way
func (s *Solution) buildAdj() {

	// Change the seed to ensure randomness
	rand.Seed(time.Now().UTC().UnixNano())

	// If the adjacency matrix alredy exits the job is done
	if s.adj != nil {
		return
	}

	// Create the slice if alredy not exist
	s.adj = make([]*list.List, len(s.Order))

	// Create the double linked list of every index of the slice
	for i := 1; i < len(s.adj); i++ {
		s.adj[i] = list.New()
	}

	// Insert a pointer of every edge to the adj matrix
	// Pointer to edge is used to avoid waste of space and unnecesary allocation
	for _, e := range s.Edges {
		l1, l2 := s.adj[e.Src], s.adj[e.Dst]

		// The randomness of rdfs is simulated by select ind a random way to push the new edge to the front or the back of the double linked list
		if rand.Int()%2 == 0 {
			l1.PushFront(e)
		} else {
			l1.PushBack(e)
		}

		if rand.Int()%2 == 0 {
			l2.PushFront(e)
		} else {
			l2.PushBack(e)
		}
	}
}

// OrderVertexes modify the vertexes array to change the order of the vertex without the need of change the edge number reference
func (s *Solution) OrderVertexes(start uint) error {
	// Build the adj matrix if it doesnt exist
	if s.adj == nil {
		s.buildAdj()
	}

	if s.Order[start] == 0 {
		return fmt.Errorf("start value dont have asigned any vertex assigned")
	}

	//Create a stack and push the starting vertex to it
	var st stack
	st.push(s.Order[start])

	// Keep track of the current position to be assign
	i := start

	// Dfs implementation
DFS:
	for curr, ok := st.pop(); ok; curr, ok = st.pop() {
		// If the value is already assign skip it
		if s.vPosition[curr] != 0 {
			continue
		}

		// If the value hasnt been assign assign a value to it
		s.Order[i] = curr
		s.vPosition[curr] = i

		// Add all the vertexes that are connected with the current vertex
		for e := s.adj[curr].Front(); e != nil; e = e.Next() {
			v := e.Value.(*Edge).otherOne(curr)

			// Just add vertexes to the stack if they are alredy not assign
			// This reduce the number of repeting vertexes that enters the stack but doesn't stop the push of vetex that are already on the stack but not assigned.
			if s.vPosition[v] != 0 {
				st.push(v)
			}
		}

		i++

	}

	// Get all conex components
	for i := uint(1); i < uint(len(s.vPosition)); i++ {
		if s.vPosition[i] == 0 {
			st.push(i)
			goto DFS
		}
	}

	return nil
}

// AssignPages iterates over all the edges an assign on page to that edge
func (s *Solution) AssignPages(limit uint) {

	// First sort the edges in decreasing order using heapsort
	heapSort(s)

	var u, v, p, q uint

	s.Crossings = 0

	var bestCross, currCross, page uint

	// Iterate over all the edges and assign it a page
	for i := uint(0); i < uint(len(s.Edges)); i++ {

		// Check which of the 2 vertexes goes first
		if s.vPosition[s.Edges[i].Src] < s.vPosition[s.Edges[i].Dst] {
			u, v = s.vPosition[s.Edges[i].Src], s.vPosition[s.Edges[i].Dst]
		} else {
			v, u = s.vPosition[s.Edges[i].Src], s.vPosition[s.Edges[i].Dst]
		}

		page = 1

		// If the length of the edge is one or n - 1 there's no need to find the best page because it cant generate crossings
		if s.getEdgeLength(i) != 1 || s.getEdgeLength(i) != s.Vertex-1 {

			// Sets all the bits of bestCross to 1 to get the maximun value
			bestCross = ^uint(0)

			// Search in all the pages which generates the minimum number of posible crossings
			for k := uint(1); k <= s.Pages; k++ {
				currCross = uint(0)

				// Loop that counts the posible crossings in the current page
				for j := uint(0); j < i; j++ {
					if k == s.Edges[j].Page {

						// Check which of the 2 vertexes goes first
						if s.vPosition[s.Edges[j].Src] < s.vPosition[s.Edges[j].Dst] {
							p, q = s.vPosition[s.Edges[j].Src], s.vPosition[s.Edges[j].Dst]
						} else {
							q, p = s.vPosition[s.Edges[j].Src], s.vPosition[s.Edges[j].Dst]
						}

						// Check cross condition
						if u < p && p < v && v < q || p < u && u < q && q < v {
							currCross++
						}
					}
				}

				// If this page has a lower number of posibles crossings assign the page to the current one
				if currCross < bestCross {
					bestCross = currCross
					page = k
				}

				// If the best number of crossings is 0 dont check the other pages
				if bestCross == 0 {
					break
				}

			}
		}

		// Once the best page was found assing it to the edge and increase the number of crossings
		s.Edges[i].Page = page
		s.Crossings += bestCross

		// If the limit is already exceeded dont calculate the remainig edges
		if s.Crossings > limit {
			return
		}

	}

}

// Copy return a complete copy of the solution if posible
func (s *Solution) Copy() (*Solution, error) {
	var sCopy *Solution
	var err error

	temp := deepcopy.Copy(s)

	sCopy, ok := temp.(*Solution)

	if !ok {
		err = fmt.Errorf("unable to copy solution")
	}

	// deepcopy only copy exported values, wee neew to manually copy some unexported values
	sCopy.vPosition = make([]uint, len(s.vPosition))
	copy(sCopy.vPosition, s.vPosition)

	return sCopy, err
}

// Swap 2 values based on s.Order value
func (s *Solution) Swap(a, b uint) {
	s.vPosition[s.Order[a]], s.vPosition[s.Order[b]], s.Order[a], s.Order[b] = s.vPosition[s.Order[b]], s.vPosition[s.Order[a]], s.Order[b], s.Order[a]
}

// ResetFrom erase all the vertex order starting from index and do rdfs to reordenate
func (s *Solution) ResetFrom(index uint) error {
	// Set all the vertex next to the start index as unassign
	for i := index; i < s.Vertex+1; i++ {
		s.vPosition[s.Order[i]] = 0
	}

	err := s.OrderVertexes(index)
	return err
}

// getEdgeLength gives the length of an edge using the position of it vetexes in the current solution
// getEdgeLength dont chech if the index of the edge exists because its an unexported function and its always called from other functions inside the package that already check it
func (s *Solution) getEdgeLength(index uint) uint {
	// Check which of the values has a higher position to avoid an overflow in the substraction
	// The values of s.vPosition are unsigned integers so an overflow will cause the number become close to 2^64 - 1 getting a wrong length value
	if s.vPosition[s.Edges[index].Src] < s.vPosition[s.Edges[index].Dst] {
		return s.vPosition[s.Edges[index].Dst] - s.vPosition[s.Edges[index].Src]
	}
	return s.vPosition[s.Edges[index].Src] - s.vPosition[s.Edges[index].Dst]
}

// CalculateCrossings erase the previus Crossings value and recalculate it from zero
func (s *Solution) CalculateCrossings() uint {
	cross := uint(0)

	var u, v, p, q uint

	for i := 0; i < len(s.Edges); i++ {

		if s.vPosition[s.Edges[i].Src] < s.vPosition[s.Edges[i].Dst] {
			u, v = s.vPosition[s.Edges[i].Src], s.vPosition[s.Edges[i].Dst]
		} else {
			v, u = s.vPosition[s.Edges[i].Src], s.vPosition[s.Edges[i].Dst]
		}

		for j := i + 1; j < len(s.Edges); j++ {
			if s.Edges[i].Page == s.Edges[j].Page {

				if s.vPosition[s.Edges[j].Src] < s.vPosition[s.Edges[j].Dst] {
					p, q = s.vPosition[s.Edges[j].Src], s.vPosition[s.Edges[j].Dst]
				} else {
					q, p = s.vPosition[s.Edges[j].Src], s.vPosition[s.Edges[j].Dst]
				}

				if u < p && p < v && v < q || p < u && u < q && q < v {
					cross++
				}
			}

		}
	}

	return cross
}
