package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeX(n int) int {
	if n < 1 {
		return -1
	}
	count := 0
	for i := 2; ; i++ {
		if isPrime(i) {
			count++
			if count == n {
				return i
			}
		}
	}
}

func main() {
	fmt.Println(primeX(1))  //2
	fmt.Println(primeX(5))  //11
	fmt.Println(primeX(8))  //19
	fmt.Println(primeX(9))  //23
	fmt.Println(primeX(10)) //29
}
