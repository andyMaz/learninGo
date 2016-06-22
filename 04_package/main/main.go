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


}
