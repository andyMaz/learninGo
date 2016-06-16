package main

import (
	//"fmt"
	"github.com/andyMaz/learninGo/02_package/stringutil"
	"fmt"
)

func main() {
	stringutil.Name()
	fmt.Println(stringutil.Reverse("Andy"))
}