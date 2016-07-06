package main

import (
	"fmt"
)

func main() {
	fmt.Println("03_package")
	Loop1(10)
	printNums(10)
	ary()
	x, y := two(10, "desk")
	fmt.Printf("%d\t%s\n", x, y)
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
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		case 6:
			fmt.Println("Six")
		case 7:
			fmt.Println("Seven")
		case 8:
			fmt.Println("Eight")
		case 9:
			fmt.Println("Nine")
		default:
			{
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

func two(x int, y string) (int, string) {
	return x, y
}
