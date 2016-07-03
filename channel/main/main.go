package main

import (
	"fmt"
	"math/rand"
	"time"
)

const max = 10000000
func sum(s []int, ch chan<- int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	ch <- sum // send sum to ch
}

func add(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}


func printAry(a [max]int) {
	for x, y := range a {
		if x%20 == 0 {
			fmt.Println()
		}
		fmt.Printf("%6d", y)
	}
	fmt.Println()
}

func buildAry() [max]int {
	var r [max]int
	i := 0
	for i < max {
		r[i] = rand.Intn(10000)%1000 + 1 //i+1
		i++
	}
	return r
}

func main() {
	r := buildAry()
	ch := make(chan int)
	s := len(r)
	start1 := time.Now()
	fmt.Printf("sum = %4d\n", add(r[:]))
	secs1 := time.Since(start1).Seconds()
	fmt.Printf("time = %f\n", secs1)
	fmt.Println("\n-------\n")
	start2 := time.Now()
	go sum(r[:s/4], ch)
	go sum(r[s/4:s/2], ch)
	go sum(r[s/2:(3*s)/4], ch)
	go sum(r[(3*s)/4:], ch)
	i := 1
	sm := 0
	for i <= 4 {
		sm += <-ch
		i++
	}
	fmt.Printf("sum = %4d\n", sm)
	secs2 := time.Since(start2).Seconds()
	fmt.Printf("time = %f\n", secs2)
	//x, y, z, w := <-ch, <-ch, <-ch, <-ch

	//	fmt.Printf("sum = %4d\n", x+y+z+w)

}
