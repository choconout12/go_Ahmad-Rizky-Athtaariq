package main

import (
	"fmt"
)

func printMultiples3(ch chan int) {
	for i := 1; ; i++ {
		if i%3 == 0 {
			ch <- i
		}
	}
}

func main() {
	ch := make(chan int)
	go printMultiples3(ch)
	for {
		fmt.Println(<-ch)
	}
}
