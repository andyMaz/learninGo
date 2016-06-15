package main

import "fmt"

func main() {
	for i := 1776; i < 1776+200; i++ {
		fmt.Printf("%4d \t %4b \t %4x \t %4q\n", i, i, i, i)
	}

}
