package main

import "fmt"

func binaryRepresentation(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = i ^ (i >> 1)
	}
	return ans
}

func main() {

	fmt.Println(binaryRepresentation(2)) // output: [0 1 2]
	fmt.Println(binaryRepresentation(3)) // output: [0 1 2 3]
}
