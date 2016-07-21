package main


import (
	"github.com/andyMaz/learninGo/linked-list/linkedList"
	"fmt"
)

func main() {
	numbers := []int{3, 2, 6, 4, 1, 10, 12, 7, 9, 5}
	var l *linkedList.List
	for _, v := range numbers {
		l = linkedList.ConsBack(v, l)
	}

	linkedList.PrintLst(l)

	var l2 *linkedList.List
	for _, v := range numbers {
		l2 = linkedList.ConsFront(v, l2)
	}
	linkedList.PrintLst(l2)
	fmt.Println("concat")
	linkedList.PrintLst(linkedList.Concat(l, l2))
	linkedList.PrintLst(l)
	linkedList.PrintLst(l2)
	l3 := linkedList.Concat(l, l2)
	linkedList.PrintLst(l3)
}