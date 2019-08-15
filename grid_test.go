package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestElementAt(t *testing.T) {
	G := Grid{}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			v := G.ElementAt(row, col)
			if v != 0 {
				t.Errorf("ElementAt(%d, %d) = %d, expected 0", row, col, v)
			}
		}
	}
}

// WithElementAt should return a new grid with the square set as specified.
func TestWithElementAt(t *testing.T) {
	// Initialize random seed.
	rand.Seed(time.Now().UnixNano())

	row := rand.Intn(9)
	col := rand.Intn(9)
	val := rand.Intn(9) + 1

	G := Grid{}
	newG := G.WithElementAt(row, col, val)

	if v := newG.ElementAt(row, col); v != val {
		t.Errorf("ElementAt(%d, %d) = %d, expected %d", row, col, v, val)
	}
}

func TestParseGrid(t *testing.T) {
	data := []byte(`010020706
700913040
380004001
000007010
500109003
090500000
200300094
040762005
105090070`)
	G, err := ParseGrid(data)
	if err != nil {
		t.Errorf("Unexpected error parsing grid: %v", err)
	}
	// Check that row 2 is as expected
	row := 2
	expected := []int{3, 8, 0, 0, 0, 4, 0, 0, 1}
	for col, want := range expected {
		got := G.ElementAt(row, col)
		if got != want {
			t.Errorf("ElementAt(%d,%d) = %d, expected %d", row, col, got, want)
		}
	}
}

func TestParseGridError(t *testing.T) {
	data := []byte(`01002070
700913040
380004001
000007010
500109003
090500000
200300094
040762005
105090070`)
	_, err := ParseGrid(data)
	if err == nil {
		t.Errorf("ParseGrid() did not return an expected error")
	}
}

func TestNeighbours(t *testing.T) {
	data := []byte(`010020706
700913040
380004001
000007010
500109003
090500000
200300094
040762005
105090070`)
	G, _ := ParseGrid(data)

	S := G.Neighbours(0, 0)

	expected := []int{1, 2, 7, 6, 3, 5, 8}
	for _, want := range expected {
		if !S.Contains(want) {
			t.Errorf("Neighbours(0,0) failed to return the neighbour %d", want)
		}
	}
	unexpected := []int{4, 9}
	for _, avoid := range unexpected {
		if S.Contains(avoid) {
			t.Errorf("Neighbours(0,0) returned invalid neighbour %d", avoid)
		}
	}
}

func TestFirstEmptyCell(t *testing.T) {
	data := []byte(`410020706
700913040
380004001
000007010
500109003
090500000
200300094
040762005
105090070`)

	G, _ := ParseGrid(data)

	idx := G.FirstEmptyCell()

	if idx != 2 {
		t.Errorf("FirstEmptyCell() failed to returned incorrect index %d", idx)
	}
}

func TestFirstEmptyCellFail(t *testing.T) {
	data := []byte(`435269781
682571493
197834562
826195347
374682915
951743628
519326874
248957136
763418259`)

	G, _ := ParseGrid(data)

	idx := G.FirstEmptyCell()

	if idx != -1 {
		t.Errorf("FirstEmptyCell() failed to return -1 for completed board")
	}
}

func TestCandidates(t *testing.T) {
	data := []byte(`010020706
700913040
380004001
000007010
500109003
090500000
200300094
040762005
105090070`)

	G, _ := ParseGrid(data)

	C := G.Candidates(0, 0)

	expected := []int{4, 9}

	sameLen := len(C) == len(expected)
	sameElems := true

	for i := range C {
		if C[i] != expected[i] {
			sameElems = false
		}
	}

	if !(sameLen && sameElems) {
		t.Errorf("Candidates(0,0) = %v, expected %v", C, expected)
	}
}

func TestSolve(t *testing.T) {
	data := []byte(`010020706
700913040
380004001
000007010
500109003
090500000
200300094
040762005
105090070`)

	G, _ := ParseGrid(data)

	solution, _ := G.Solve()

	expectedData := []byte(`419825736
756913248
382674951
634287519
527149863
891536427
278351694
943762185
165498372
`)

	expected, _ := ParseGrid(expectedData)

	sameLen := len(solution) == len(expected)
	sameElems := true

	for i := range solution {
		if solution[i] != expected[i] {
			sameElems = false
		}
	}

	if !(sameLen && sameElems) {
		fmt.Println(solution)
		fmt.Println(expected)
		t.Errorf("Solved puzzle does not match solution.")
	}
}
