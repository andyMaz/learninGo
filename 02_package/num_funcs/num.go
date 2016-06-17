package stringutil

import "fmt"

func factorial(n int) int {
	var r int = 1
	if n > 1 {
		r = 1
		for i := 1; i <= n; i++ {
			r = r * i
		}
	}
	return r
}


func fact(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * fact(n - 1)
	}
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-2) + fib(n-1)
	}
}

func Fibonacci(n int) {
	for i := 0; i <= n; i++ {
		fmt.Printf("%5d\t", i)
	}
	fmt.Println()
	for i := 0; i <= n; i++ {
		fmt.Printf("%5d\t", fib(i))
	}
	fmt.Println()
}
