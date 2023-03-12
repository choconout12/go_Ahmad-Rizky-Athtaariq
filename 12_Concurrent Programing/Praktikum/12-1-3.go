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
	ch := make(chan int, 10) // buat channel dengan buffer size 10
	go printMultiples3(ch)
	for i := 1; i <= 20; i++ {
		fmt.Println(<-ch)
	}
}
