package main

import (
	"fmt"
	"github.com/andyMaz/learninGo/02_package/stringutil"
	"github.com/andyMaz/learninGo/02_package/num_funcs"
)
var s string

var f float64
func main() {
	q := "Danny"
	s = "Joey"
	i := 1
	fmt.Printf("%T\t%10d\n",i, i)
	fmt.Printf("%T\t%10s\n",s, s)
	fmt.Printf("%T\t%10f\n",f, f)
	fmt.Printf("%T\t%10s\n",q, q)
	stringutil.Name()
	fmt.Println(stringutil.Reverse("Andy"))
	fmt.Printf("\n-----\n")
	zero()
	fmt.Printf("\n-----\n")
	fmt.Printf("%d\n", num_funcs.Factorial(10))
	fmt.Printf("%d\n", num_funcs.Fact(10))
	fmt.Printf("\n-----\n")
	num_funcs.Fibonacci(20)
	stringArray()
	fmt.Printf("\n-----\n")
	fmt.Print("BubbleSort")
	var ary1 = num_funcs.RandArray(40)
	num_funcs.BubbleSort(ary1)
	num_funcs.PrintArray(ary1)
	fmt.Print("\nInsertSort")
	var ary2 = num_funcs.RandArray(40)
	num_funcs.InsertSort(ary2)
	num_funcs.PrintArray(ary2)
	test("Hi Danny")
	fmt.Printf("\n%s\n", stringutil.Test())

}

func test(n string) {
	fmt.Printf("\n%s", n)
}

func zero() {
	var a int
	var b string
	var c float64
	var d bool
	fmt.Printf("%T\t%10d\n",a, a)
	fmt.Printf("%T\t%10s\n",b, b)
	fmt.Printf("%T\t%10f\n",c, c)
	fmt.Printf("%T\t%10t\n",d, d)
}




func stringArray() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// here are how to invoke higher order functions
	// three different ways of calling function funcPara
	// that requires function as parameter
	g := func(x int) int {
		x++
		return x
	}
	fmt.Printf("%T\n",g)
	funcPara(g, primes)
	funcPara(func(x int) int {
		x--
		return x
	}, primes)
	funcPara(add2, primes)
	fmt.Printf("%d\n", funcResult(3)(4))
}


// the following two functions are higher order functions
// an example of func that accepts function as parameter
func funcPara(f func(int) int, a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%3d\t", f(a[i]))
	}
	fmt.Println()
}

// an example of func that returns a function as result
func funcResult(x int) func(int) int {
	return func(y int) int {
		return x+y
	}
}

func add2(x int) int {
	return x+2
}

