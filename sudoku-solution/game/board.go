package game

import (
	"encoding/csv"
	"io"
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

type Board struct {
	Puzzle [9][9]int
}

func ReadPuzzle(fileName string) Board {
	b := new(Board)
	f, _ := os.Open(fileName)

	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		for value := range record {
			cell := strings.Split(record[value], "=")
			c , _ := strconv.Atoi(strings.TrimSpace(cell[0]))
			v , _ := strconv.Atoi(strings.TrimSpace(cell[1]))
			x := c / 10
			y := c - (x * 10)
			b.Puzzle[x - 1][y - 1] = v
		}
	}
	return *b
}

func (b Board) String() string {
	temp := ""
	for i, v := range b.Puzzle {
		if i % 3 == 0 {
			temp += "\n"
		}
		for j, c := range v {
			if j % 3 == 0 {
				temp += "\t"
			}
			sc := strconv.Itoa(c)
			temp += sc + "  "
		}
		temp += "\n"
	}
	return fmt.Sprintf(temp)
}


