package main

import "testing"

func TestPush(t *testing.T) {
	S := GridStack{}

	g1 := Grid{}
	g2 := Grid{}

	S.Push(g1)
	S.Push(g2)

	if size := len(S.grids); size != 2 {
		t.Errorf("TestPush(g) expected size %d but received %d", 2, size)
	}
}

func TestPop(t *testing.T) {
	S := GridStack{}

	g1 := Grid{}
	g2 := g1.WithElementAt(0, 0, 1)

	S.Push(g1)
	S.Push(g2)
	p, err := S.Pop()

	if err != nil {
		t.Errorf("TestPop() returned an unexpected error")
	}

	if p.ElementAt(0, 0) != g2.ElementAt(0, 0) {
		t.Errorf("TestPop() did not return last pushed grid")
	}

	if size := len(S.grids); size != 1 {
		t.Errorf("TestPop() expected size %d but received %d", 1, size)
	}
}

func TestPopError(t *testing.T) {
	S := GridStack{}

	_, err := S.Pop()

	if err == nil {
		t.Errorf("TestPop() did not return expected error")
	}
}
