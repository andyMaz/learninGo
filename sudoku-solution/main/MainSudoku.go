package main

import (

	"sudoku-solution/game"
	"fmt"
)

const DIR = "io/"

func main() {
	b := game.ReadPuzzle(DIR + "/sudoku_puzzle.txt")
	game.Solve(b)
	game.Help()
	fmt.Println()
	game.Sudoku(0, "solution.txt")
}
