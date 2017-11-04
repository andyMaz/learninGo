package game

import (
	"fmt"
	"strconv"
)


var board Board

func Solve(b Board) {
	board = b
}

func row(i int) []int {
	r := make([]int, 0)
	for j := 0; j < 9; j++ {
		if board.Puzzle[i][j] != 0 {
			r = append(r, board.Puzzle[i][j])
		}
	}
	return r
}

func column(j int) []int {
	r := make([]int, 0)
	for i := 0; i < 9; i++ {
		if board.Puzzle[i][j] != 0 {
			r = append(r, board.Puzzle[i][j])
		}
	}
	return r
}

func findBoxStart(n int) int {
	return 3 * (n / 3)
}

func box(i, j int) []int {
	si := findBoxStart(i)
	sj := findBoxStart(j)
	r := make([]int, 0)
	for p := si; p < si + 3; p++ {
		for q := sj; q < sj + 3; q++ {
			if board.Puzzle[p][q] != 0 {
				r = append(r, board.Puzzle[p][q])
			}
		}
	}
	return r
}

func contains(e int, s []int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Possibilities(i, j int) []int {
	s := make([]int, 0)
	r := make([]int, 0)
	if board.Puzzle[i][j] == 0 {
		s = append(s, row(i)...)
		s = append(s, column(j)...)
		s = append(s, box(i, j)...)
		for p := 1; p < 10; p++ {
			if contains(p, s) == false {
				r = append(r, p)
			}
		}
	}
	return r
}

func isUnique(r, c int) bool {
	if !rowUnique(r, c) {
		return false
	}
	if !columUnique(r, c) {
		return false
	}
	return boxUnique(r, c)
}

func rowUnique(r, c int) bool {
	e := board.Puzzle[r][c]
	for j := 0; j < 9; j++ {
		if j != c && board.Puzzle[r][j] == e {
			return false
		}
	}
	return true
}

func columUnique(r, c int) bool {
	e := board.Puzzle[r][c]
	for i := 0; i < 9; i++ {
		if i != r && board.Puzzle[i][c] == e {
			return false
		}
	}
	return true
}

func boxUnique(r, c int) bool {
	si := findBoxStart(r)
	sj := findBoxStart(c)
	e := board.Puzzle[r][c]
	for p := si; p < si + 3; p++ {
		for q := sj; q < sj + 3; q++ {
			if p != r && q != c && board.Puzzle[p][q] == e {
				return false
			}
		}
	}
	return true
}

//solved is not needed
func Solved() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			r := isUnique(i, j)
			if r == false {
				return false
			}
		}
	}
	return true
}

func printLst(ls []int)  string {
	temp := ""
	for _, l := range ls {
		temp += strconv.Itoa(l)
	}
	return temp
}


func Help()  {
	fmt.Println()
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board.Puzzle[i][j] == 0 {
				fmt.Printf("%8s",  printLst(Possibilities(i, j))+ " ")
			} else {
				fmt.Printf("%8s",  strconv.Itoa(board.Puzzle[i][j]) + " ")
			}
			if j%3 == 2  && j != 8 {
				fmt.Print("  |")
			}
		}
		fmt.Println()
		if i%3 == 2 && i != 8 {
			fmt.Println("   --------------------------------------------------------------------------")
		}
	}

}

func Sudoku(n int, outFile string)  {
		if n < 81 {
			r := n / 9
			c := n % 9
			if board.Puzzle[r][c] == 0 {
				ps := Possibilities(r, c)
				for _, x := range ps {
					board.Puzzle[r][c] = x // tentative placing of value x at position n
					Sudoku(n+1, outFile)
					//add()
					// checking if the game is full and a solution is found
					if n+1 == 81 { // && Solved() {
						fmt.Println(Board{board.Puzzle})
						//WriteSolution(outFile)
					}
					board.Puzzle[r][c] = 0 // backtracking
				}
			} else {
				// checking if the game is full and a solution is found
				if n+1 == 81 { //&& Solved() {
					fmt.Println(Board{board.Puzzle})
					//WriteSolution(outFile);
				}
				Sudoku(n+1, outFile)
			}
		}
}
