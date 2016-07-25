package tree

import (
	"fmt"
	"strconv"
	"github.com/andyMaz/learninGo/linked-list/linkedList"
	"container/list"
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func depth(t *Tree) int {
	if t == nil {
		return 0
	}
	return 1 + max(depth(t.left), depth(t.right))
}


func depthN(n int, t *Tree)  string {
	if n == 0 || t == nil {
		return ""
	} else if n == 1 && t != nil {
		return strconv.Itoa(t.data) + " "
	}
	return depthN(n-1, t.left) +  depthN(n-1, t.right)
}

func depthNlist(n int, t *Tree)   *linkedList.List {
	if n == 0 || t == nil {
		return nil
	} else if n == 1 && t != nil {
		return linkedList.ConsBack(t.data, nil)
	}
	return linkedList.Concat(depthNlist(n-1, t.left),  depthNlist(n-1, t.right))
}

func size(t *Tree) int {
	if t == nil {
		return 0
	}
	return 1 + size(t.left) + size(t.right)
}

func Level_by_level(t *Tree){
	var n int
	if t  != nil {
		n = depth(t)
	}
	for i := 1; i < n+1; i++ {
		s := linkedList.LstToSting(depthNlist(i, t))
		fmt.Printf("%d -> %s\n", i, s)
	}
}

func Flat(t *Tree) *linkedList.List {
	if t == nil {
		return nil
	}
	return linkedList.Concat(linkedList.ConsFront(t.data, Flat(t.left)), Flat(t.right))
}

func  Bfs(t *Tree) {
	var que = list.New()
	que.PushBack(t)
	for que.Len() > 0 {
		v := que.Front()
		child := [] *Tree{0: v.Value.(*Tree).left, 1: v.Value.(*Tree).right}
		fmt.Printf("%d\t", v.Value.(*Tree).data)
		que.Remove(v)
		for i := range child {
			if child[i] != nil {
				que.PushBack(child[i])
			}
		}
	}
}

