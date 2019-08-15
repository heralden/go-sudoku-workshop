package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please specify only the file of a sudoku puzzle")
		return
	}

	file := args[0]
	data, readErr := ioutil.ReadFile(file)

	if readErr != nil {
		fmt.Println(readErr)
		return
	}

	g, err := ParseGrid(data)

	if err != nil {
		fmt.Println(err)
		return
	}

	solved, solveErr := g.Solve()

	if solveErr != nil {
		fmt.Println(solveErr)
	} else {
		fmt.Println(solved)
	}
}
