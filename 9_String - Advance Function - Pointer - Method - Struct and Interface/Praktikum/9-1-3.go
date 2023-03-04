package main

import (
	"fmt"
)

func Compare(a, b string) string {
	var minLen int
	if len(a) < len(b) {
		minLen = len(a)
	} else {
		minLen = len(b)
	}

	for i := 0; i < minLen; i++ {
		if a[i:i+1] != b[i:i+1] {
			return a[0:i]
		}
	}

	return a[0:minLen]
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}
