package main

import (
		"fmt"
		"runtime"
	)

func main() {
	fmt.Println("Hello World!")
	test()
}

func test() {
	fmt.Printf("%s%s\n", "Go runs on ", runtime.GOOS)
	x := 5
	y := 10
	fmt.Printf("x = %d\ty = %d\n", x, y)
	x, y = swap(x, y)
	fmt.Printf("x = %d\ty = %d\n", x, y)
	fmt.Printf("%x\t%x\t%x\t%x\n", 85, 165, 34, 50)
}


func swap(x, y int) (int, int) {
	return y, x
}
