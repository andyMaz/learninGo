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
	fmt.Println(i)
	fmt.Println(s)
	fmt.Println(f)
	fmt.Println(q)
	stringutil.Name()
	fmt.Println(stringutil.Reverse("Andy"))
}