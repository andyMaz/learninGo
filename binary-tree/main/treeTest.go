package main

import (
	"github.com/andyMaz/learninGo/binary-tree/tree"
	"fmt"
	"github.com/andyMaz/learninGo/linked-list/linkedList"
)

func main() {
	numbers := []int{3, 2, 6, 4, 1, 10, 12, 7, 9, 5}
	var t *tree.Tree
	for _, v := range numbers {
		t = tree.Insert(v, t)
	}

	tree.In_order(t)
	fmt.Println()
	tree.Level_by_level(t)
	fmt.Println()
	l := tree.Flat(t)
	s := linkedList.LstToSting(l)
	fmt.Println(s)

}
