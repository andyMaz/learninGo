package tree

import (
	"fmt"
)

type Tree struct {
	left *Tree
	data int
	right *Tree
}

func Cons(l *Tree, d int, r *Tree) *Tree {
	return &Tree {left : l, data : d, right : r}
}

func  Insert(e int, t *Tree) *Tree {
	if t == nil {
		return Cons(nil, e, nil)
	} else if e <= t.data {
		return Cons(Insert(e, t.left), t.data, t.right)
	} else {
		return &Tree{left: t.left, data: t.data, right: Insert(e, t.right)}
	}
}

func In_order(t *Tree) {
	if t != nil {
		In_order(t.left)
		fmt.Printf("%d\t", t.data)
		In_order(t.right)
	}
}

