package main

import (
	"fmt"
)

func main() {
	fmt.Println("03_package")
	Loop1(10)
	printNums(7)
	ary()
}

func Loop1(n int) {
	i := 1
	for i <= n {
		fmt.Printf("%5d\t", i)
		i++
	}
	fmt.Println()
}

func printNums(n int) {
	i := 0
	Loop:
	for i <= n {
		switch i {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		case 4: fmt.Println("Four")
		case 5: fmt.Println("Five")
		default: {
			fmt.Println("Unknown Number")
			break Loop
		}
		}
		i++
	}
}

func ary() {
	var x [5]int
	for i := 0; i < len(x); i++ {
		x[i] = 100 + 1
	}
	fmt.Println(x)
}

