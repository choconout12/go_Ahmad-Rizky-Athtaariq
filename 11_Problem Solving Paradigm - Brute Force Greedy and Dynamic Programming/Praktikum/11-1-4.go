package main

import (
	"fmt"
	"math"
)

func SimpleEquations(a, b, c int) {
	// Menggunakan rumus untuk mencari nilai x, y, dan z
	// dari tiga hubungan yang diberikan
	x := math.Pow(float64(c)-float64(a)*float64(a)/3.0-float64(b), 1.0/3.0)
	y := (float64(a) - x*x - float64(b)/x) / 2.0
	z := float64(a) - x - y

	// Mengecek apakah solusi bilangan bulat ada
	if math.Round(x)+math.Round(y)+math.Round(z) == float64(a) && math.Round(x)*math.Round(y)*math.Round(z) == float64(b) && math.Round(x*x)+math.Round(y*y)+math.Round(z*z) == float64(c) {
		fmt.Println(int(math.Round(x)), int(math.Round(y)), int(math.Round(z)))
	} else {
		fmt.Println("no solution")
	}
}

func main() {
	SimpleEquations(1, 2, 3)  //no solution
	SimpleEquations(6, 6, 14) //1 2 3
}
