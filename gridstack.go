package main

import "fmt"

type GridStack struct {
	grids []Grid
}

func (s *GridStack) Push(g Grid) {
	s.grids = append(s.grids, g)
}

func (s *GridStack) Pop() (Grid, error) {
	lastIdx := len(s.grids) - 1

	if lastIdx < 0 {
		return Grid{}, fmt.Errorf("No more grids in GridStack")
	}

	g := s.grids[lastIdx]
	s.grids = s.grids[0:lastIdx]

	return g, nil
}
