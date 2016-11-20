package main

import (
	"github.com/andyMaz/backtracking/sujiko-solutions/sujiko"
)

func main() {
	corners := make([]int, 4)
	corners[0] = 17
	corners[1] = 24
	corners[2] = 15
	corners[3] = 21

	givens := make(map[int]int)
	givens[2] = 5
	givens[7] = 3

	sujiko.SujikoCons(corners, givens)
	sujiko.Bord()
	sujiko.Sujiko(0)

}
