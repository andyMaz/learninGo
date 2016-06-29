package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
)


func main() {
	echo()
}

func echo() {
	fileName := "main.go"
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
