package main

import (
	"fmt"
	"sync"
)

func factorial(n int, wg *sync.WaitGroup, mu *sync.Mutex, result *int) {
	mu.Lock()
	defer mu.Unlock()
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	*result *= fact
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	result := 1
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go factorial(i, &wg, &mu, &result)
	}
	wg.Wait()
	fmt.Println("Result:", result)
}
