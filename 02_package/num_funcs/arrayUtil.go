package num_funcs

import (
	"fmt"
	"math/rand"
)

func RandArray(n int) []int {
	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(10000)%100 + 1
	}
	return a
}

func PrintArray(a []int) {
	for i := 0; i < len(a); i++ {
		if i%10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%5d", a[i])

	}
}
