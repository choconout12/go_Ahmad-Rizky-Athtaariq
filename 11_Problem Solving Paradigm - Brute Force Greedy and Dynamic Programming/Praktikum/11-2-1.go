package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	n := len(jumps)

	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = math.MaxInt32
	}
	p[0] = 0

	for i := 0; i < n-1; i++ {
		for j := i + 1; j <= i+2 && j < n; j++ {
			cost := int(math.Abs(float64(jumps[i] - jumps[j])))
			if p[i]+cost < p[j] {
				p[j] = p[i] + cost
			}
		}
	}
	return p[n-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         //30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) //40
}
