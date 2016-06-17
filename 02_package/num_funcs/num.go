package num_funcs

import "fmt"

func Factorial(n int) int {
	var r int = 1
	if n > 1 {
		r = 1
		for i := 1; i <= n; i++ {
			r = r * i
		}
	}
	return r
}


func Fact(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * Fact(n - 1)
	}
}

func Fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fib(n-2) + Fib(n-1)
	}
}

func Fibonacci(n int) {
	for i := 0; i <= n; i++ {
		fmt.Printf("%5d\t", i)
	}
	fmt.Println()
	for i := 0; i <= n; i++ {
		fmt.Printf("%5d\t", Fib(i))
	}
	fmt.Println()
}
