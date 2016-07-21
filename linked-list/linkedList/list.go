package linkedList

import (
	"fmt"
	"strconv"
)

type List struct {
	data int
	next *List
}

func ConsBack(e int, n *List) *List {
	return &List {data : e, next : n}
}

func ConsFront(i int, n *List) *List {
	if n == nil {
		return &List {data : i, next : n}
	}
	e := n
	for ; e.next != nil; e = e.next {}
	e.next = &List{data : i, next : nil}
	return n
}

func ConsFt(i int, n *List) *List {
	if n == nil {
		return &List {data : i, next : n}
	}
	e := n
	for ; e.next != nil; e = e.next {}
	e.next = &List{data : i, next : nil}
	return n
}


func Concat(a, b *List) *List {
	if b == nil {
		return a
	}
	return Concat(ConsFront(b.data, a), b.next)
}

func PrintLst(a *List) {
	for e := a; e != nil; e = e.next {
		fmt.Printf("%d\t", e.data)
	}
	fmt.Println()
}

func LstToSting(a *List) string {
	temp := ""
	for e := a; e != nil; e = e.next {
		temp += strconv.Itoa(e.data) + " "
	}
	return temp
}
