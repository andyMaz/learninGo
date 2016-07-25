package tree


func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func Depthbal(t *Tree) bool {
	if t == nil {
		return true
	}
	return abs(depth(t.left) - depth(t.right)) <= 1 && (Depthbal(t.left) && Depthbal(t.right))
}

func slope(t *Tree) int {
	if t == nil {
		return 0
	}
	return depth(t.left) - depth(t.right)
}

func rotr(t *Tree) *Tree {
	return Cons(t.left.left, t.left.data,  Cons(t.left.right, t.data,  t.right))
}

func rotl(t *Tree) *Tree {
	return Cons(Cons(t.left, t.data, t.right.left), t.right.data, t.right.right)
}

func shiftl(t *Tree) *Tree {
	if slope(t.right) == 1 {
		return rotl(Cons(t.left, t.data, rotr(t.right)))
	}
	return rotl(t)
}

func shiftr(t *Tree) *Tree {
	if slope(t.left) == -1 {
		return rotr(Cons(rotl(t.left), t.data, t.right))
	}
	return rotr(t)
}

func rebal(t *Tree) *Tree {
	if slope(t) == 2 {
		return shiftr(t)
	} else if slope(t) == -2 {
		return shiftl(t)
	}
	return t
}

func Binsert(e int, t *Tree) *Tree {
	if t == nil {
		return Cons(nil, e, nil)
	} else if e < t.data {
		return rebal(Cons(Binsert(e, t.left), t.data, t.right))
	} else if e == t.data {
		return t
	} else {
		return rebal(Cons(t.left, t.data, Binsert(e, t.right)))
	}
}


