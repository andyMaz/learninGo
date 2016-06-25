package main

import "fmt"

const (
	pi =  3.14
	Language = "Go"
)

const (
	A = iota
	B = iota
	C = iota
	D = iota
)

const (
	_ = iota
	X = 1 << (iota * 10)
	Y = 1 << (iota * 10)
)

func main() {
	fmt.Printf("%3.2f\n", 3.14)
	fmt.Printf("%s\n", Language)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println()
	fmt.Println(X)
	fmt.Println(Y)

	as := []int{4,6,3,10,16,1,8,23,15}
	b := isIn2(21, as)
	fmt.Println(b)
	print(as)
	loopy(as)

}

func print(ls []int)  {
	for i := range ls {
		fmt.Printf("%d\t", ls[i])
	}
	fmt.Println()
}

func isIn2(e int, ls []int) bool {
	for i := range ls {
		if ls[i] == e {
			return true
		}
	}
	return false
}

func isIn(e int, ls []int) bool {
	for _, y := range ls {
		if y == e {
			return true
		}
	}
	return false
}

func loopy(ls []int) {
	for x, y := range ls {
		fmt.Print(x, "\t")
		fmt.Println(y)
	}

}