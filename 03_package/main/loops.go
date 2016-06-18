package main

import "fmt"

func loop1(n int) {
	i := 1
	for i <= 10 {
		fmt.Printf("%5d\t", i)
	}
	fmt.Println()
}

