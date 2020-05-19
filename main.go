package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/garabais/kpmp/kpage"
)

func main() {
	err := solve(strings.NewReader("4 2 1 3 2 4"), os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

func solve(in io.Reader, out io.Writer) error {
	var s, sc *kpage.Solution

	// Generate initial solution
	s, err := generateInitialSolution(in, out)
	if err != nil {
		return err
	}

	sc, err = s.Copy()

	if err != nil {
		return fmt.Errorf("unable to copy %T", sc)
	}

	fmt.Fprintln(out, s)
	fmt.Fprintln(out, sc)
	// sc.Order[0] = 999
	// sc.Crossings = 999
	// sc.Pages = 999
	// sc.Edges[0].Page = 999
	// fmt.Fprintln(out, s)
	// fmt.Fprintln(out, sc)

	return nil
}

func generateInitialSolution(in io.Reader, out io.Writer) (*kpage.Solution, error) {
	var v, e uint
	fmt.Fscan(in, &v)
	fmt.Fscan(in, &e)

	edg := make([]*kpage.Edge, e)

	for i := uint(0); i < e; i++ {
		var src, dst uint

		fmt.Fscan(in, &src)
		fmt.Fscan(in, (&dst))

		edg[i] = kpage.NewEdge(src, dst)
	}

	s, err := kpage.Solve(edg, v, 1)

	if err != nil {
		return nil, err
	}

	return s, nil
}
