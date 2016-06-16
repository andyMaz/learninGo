package main

import (
	"github.com/andyMaz/learninGo/02_package/stringutil"
	"fmt"
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
}