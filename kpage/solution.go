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

func (s *Solution) buildAdj() {

	rand.Seed(time.Now().UTC().UnixNano())

	if s.adj != nil {
		return
	}
	s.adj = make([]*list.List, len(s.Order))
	for i := 1; i < len(s.adj); i++ {
		s.adj[i] = list.New()
	}

	for _, e := range s.Edges {
		l1, l2 := s.adj[e.Src], s.adj[e.Dst]

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
	if s.adj == nil {
		s.buildAdj()
	}

	if s.Order[start] == 0 {
		return fmt.Errorf("start value dont have asigned any vertex assigned")
	}

	var st stack

	st.push(s.Order[start])

	i := start

DSL:
	for curr, ok := st.pop(); ok; curr, ok = st.pop() {
		if s.vPosition[curr] != 0 {
			continue
		}

		s.Order[i] = curr
		s.vPosition[curr] = i

		for e := s.adj[curr].Front(); e != nil; e = e.Next() {
			v, err := e.Value.(*Edge).otherOne(curr)
			if err != nil {
				return err
			}
			st.push(v)
		}

		i++

	}

	// get all conex components
	for i := uint(1); i < uint(len(s.vPosition)); i++ {
		if s.vPosition[i] == 0 {
			st.push(i)
			goto DSL
		}
	}

	return nil
}

// AssignPages iterates over all the edges an assign on page to that edge
func (s *Solution) AssignPages() {

	// sort.Slice(s.Edges, func(i, j int) bool {
	// 	// This function is called by the function as Less, but we are insted using as More so we use ">" insted of "<"
	// 	return math.Abs(float64(s.vPosition[s.Edges[i].Src]-s.vPosition[s.Edges[i].Dst])) > math.Abs(float64(s.vPosition[s.Edges[j].Src]-s.vPosition[s.Edges[j].Dst]))
	// })

	heapSort(s)

	var u, v, p, q uint

	s.Crossings = 0

	var bestCross, currCross, page uint

	for i := 0; i < len(s.Edges); i++ {

		if s.vPosition[s.Edges[i].Src] < s.vPosition[s.Edges[i].Dst] {
			u, v = s.vPosition[s.Edges[i].Src], s.vPosition[s.Edges[i].Dst]
		} else {
			v, u = s.vPosition[s.Edges[i].Src], s.vPosition[s.Edges[i].Dst]
		}

		bestCross = ^uint(0) // Sets all the bits of bestCross to 1 to get the maximun value
		page = 0
		for k := uint(1); k <= s.Pages; k++ {
			currCross = uint(0)

			for j := 0; j < i; j++ {
				if k == s.Edges[j].Page {

					if s.vPosition[s.Edges[j].Src] < s.vPosition[s.Edges[j].Dst] {
						p, q = s.vPosition[s.Edges[j].Src], s.vPosition[s.Edges[j].Dst]
					} else {
						q, p = s.vPosition[s.Edges[j].Src], s.vPosition[s.Edges[j].Dst]
					}

					if u < p && p < v && v < q || p < u && u < q && q < v {
						currCross++
					}
				}
			}

			if currCross < bestCross {
				bestCross = currCross
				page = k
			}

		}
		s.Edges[i].Page = page
		s.Crossings += bestCross

	}

	// s.CalculateCrossings()
}

// CalculateCrossings erase the previus Crossings value and recalculate it from zero
func (s *Solution) CalculateCrossings() {
	s.Crossings = 0

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

				if u < p && p < v && v < q {
					s.Crossings++
				}
			}

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
	for i := index; i < s.Vertex+1; i++ {
		s.vPosition[s.Order[i]] = 0
	}

	err := s.OrderVertexes(index)
	return err
}
