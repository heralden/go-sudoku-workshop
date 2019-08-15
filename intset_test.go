package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestContains(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	x := rand.Intn(9) + 1

	S := IntSet{x}

	if contains := S.Contains(x); !contains {
		t.Errorf("TestContains(%d) = false, expected true", x)
	}
}

func TestInsert(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	x := rand.Intn(9) + 1

	S := IntSet{}

	S.Insert(x)

	if got := S[0]; got != x {
		t.Errorf("TestInsert(%d) did not add %d to the set.", x, x)
	}
}
