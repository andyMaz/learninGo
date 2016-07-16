package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("\n")
	result := isbn(6162754)
	for _, v := range result {
		fmt.Printf("%d  ", v)
	}
	fmt.Printf("\n")
	rs := isbnStr("0191012203")

	fmt.Printf("%d", rs)

}

func isbn(n int) []int {
	result := make([]int, 7)
	i := 0
	for n > 1 {
		result[i] = n%10
		i++
		n = n/10
	}
	return result
}

func isbnStr(ns string) int {
	result := 0
	m := len(ns)
	for _, v :=  range ns {
		i, _ := strconv.Atoi(string(v))
		result += i * m
		m--
	}
	return result

}