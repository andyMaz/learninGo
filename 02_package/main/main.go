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
	//fmt.Printf("%d\n", factorial(10))
	//fmt.Printf("%d\n", fact(10))
	fmt.Printf("\n-----\n")
	num_funcs.Fibonacci(20)
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
