package main

import (
	"fmt"
	"strings"
)

type Grid [81]int

func index(row, col int) int {
	return col + row*9
}

func (g Grid) ElementAt(row, col int) int {
	ix := index(row, col)
	return g[ix]
}

func (g Grid) WithElementAt(row, col int, val int) Grid {
	ix := index(row, col)
	g[ix] = val
	return g
}

func ParseGrid(data []byte) (Grid, error) {
	g := Grid{}

	i := 0
	for _, d := range data {
		v := int(d) - int('0')

		if v >= 0 && v < 10 {
			g[i] = v
			i++
		}
	}

	var e error
	if i != 81 {
		e = fmt.Errorf("Expected a puzzle length of %d but received %d", 80, i)
	}

	return g, e
}

// Make printing our grid pretty!
func (g Grid) String() string {
	var sb strings.Builder
	for r := 0; r < 9; r++ {
		fmt.Fprintf(&sb, "%d %d %d | %d %d %d | %d %d %d\n",
			g.ElementAt(r, 0), g.ElementAt(r, 1), g.ElementAt(r, 2),
			g.ElementAt(r, 3), g.ElementAt(r, 4), g.ElementAt(r, 5),
			g.ElementAt(r, 6), g.ElementAt(r, 7), g.ElementAt(r, 8))
		if r == 2 || r == 5 {
			sb.WriteString("------+-------+------\n")
		}
	}
	return sb.String()
}

// Neighbours returns the values of the cells in the same row, column, or subgrid
// as the target cell.
func (g Grid) Neighbours(row, col int) IntSet {
	S := IntSet{}

	// Row
	for i := 0; i < 9; i++ {
		if v := g.ElementAt(row, i); v > 0 {
			S.Insert(v)
		}
	}

	// Col
	for i := 0; i < 9; i++ {
		if v := g.ElementAt(i, col); v > 0 {
			S.Insert(v)
		}
	}

	// Subgrid
	subrow := row / 3
	subcol := col / 3

	for r := 3 * subrow; r < 3*(subrow+1); r++ {
		for c := 3 * subcol; c < 3*(subcol+1); c++ {
			if v := g.ElementAt(r, c); v > 0 {
				S.Insert(v)
			}
		}
	}

	return S
}

// FirstEmptyCell returns the array index of the first empty cell in the
// grid, or -1 if there are no unfilled cells
func (g Grid) FirstEmptyCell() int {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if g.ElementAt(row, col) == 0 {
				return index(row, col)
			}
		}
	}

	return -1
}

func (g Grid) Candidates(row, col int) []int {
	neighbours := g.Neighbours(row, col)
	var candidates []int
	for _, v := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		if !neighbours.Contains(v) {
			candidates = append(candidates, v)
		}
	}
	return candidates
}

func (g Grid) Solve() (Grid, error) {
	i := g.FirstEmptyCell()

	if i == -1 {
		return g, nil
	}

	row := i / 9
	col := i % 9

	candidates := g.Candidates(row, col)

	for _, v := range candidates {
		result, err := g.WithElementAt(row, col, v).Solve()
		if err == nil {
			return result, nil
		}
	}

	return g, fmt.Errorf("No solutions found")
}
