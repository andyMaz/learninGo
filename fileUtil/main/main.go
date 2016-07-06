package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
)

const dir = "C:/Users/andrei/Documents/GitHub/Goworkspace/src/github.com/andyMaz/learninGo/02_package/num_funcs/"

func main() {
	echo()
}

func echo() {
	fileName := dir + "sort.go"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	input := bufio.NewScanner(f)
	for input.Scan() {
		fmt.Printf("%s\n", input.Text())
	}
}
