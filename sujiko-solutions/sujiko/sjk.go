package sujiko

import (
	"fmt"
	"strconv"
)

const board_size = 9

var board = make([]int, board_size)
var corners []int
var givens map[int]int

func SujikoCons(c []int, g map[int]int) {
	corners = c
	givens = g
}

/*
	This function fills the board with the initial given values.
	Notice, it relies on go's automatic initialisation to
	fill the rest of the board with default value zero
*/
func Bord() {
	for k, v := range givens {
		board[k] = v
	}
}

func ToString() string {
	temp := ""
	for n := 0; n < board_size; n++ {
		if n%3 == 0 {
			temp += "\n"
		}
		temp += "  " + strconv.Itoa(board[n])
	}
	return fmt.Sprintf("%s", temp)
}

/*
	this function checks if condition for solving the puzzle is met
*/
func solved() bool {
	condition := board[0]+board[1]+board[3]+board[4] == corners[0]
	condition = condition && (board[1]+board[2]+board[4]+board[5] == corners[1])
	condition = condition && (board[3]+board[4]+board[6]+board[7] == corners[2])
	condition = condition && (board[4]+board[5]+board[7]+board[8] == corners[3])
	return condition
}

func contains(e int, s []int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

/*
	this function computes what number(s) needs to
	be placed, given the current state of the board
*/
func to_be_placed() []int {
	var temp = make([]int, 0)
	for x := 1; x <= board_size; x++ {
		if contains(x, board) == false {
			temp = append(temp, x)
		}
	}
	return temp
}

/*
	this function placed numbers in the board and
	backtracks when reached the end and conditions are not met.
*/
func Sujiko(n int) {
	if n < board_size {
		if board[n] == 0 {
			sf := to_be_placed()
			for _, x := range sf {
				board[n] = x // tentative placing of value x at position n
				Sujiko(n + 1)
				// checking if the board is full and a solution is found
				if n+1 == board_size && solved() {
					fmt.Println(ToString())
				}
				board[n] = 0 // backtracking
			}
		} else {
			// checking if the board is full and a solution is found
			if n+1 == board_size && solved() {
				fmt.Println(ToString())
			}
			Sujiko(n + 1)
		}
	}
}
