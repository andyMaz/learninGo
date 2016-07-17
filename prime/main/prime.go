package main

import (
	"container/list"
	"fmt"
	"math"
)

var initials = [11]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
var primes = list.New()

func main() {
	for _, v := range initials {
		primes.PushBack(v)
	}
	primeNumbers(1009)
	count := 0
	for e := primes.Front(); e != nil; e = e.Next() {
		if count%10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%d\t\t", e.Value)
		count++
	}
	fmt.Println()
}

func primeNumbers(num int) {
	for n := 33; n <= num; n += 2 {
		if isPrime(n) {
			primes.PushBack(n)
		}
	}
}

func isPrime(num int) bool {
	for n := primes.Front(); n != nil; n = n.Next() {
		sqrt := math.Sqrt(float64(num))
		if n.Value.(int) <= int(sqrt) {
			if num%n.Value.(int) == 0 {
				return false
			}
		} else {
			return true
		}
	}
	return false
}
