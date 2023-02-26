package main

import "fmt"

func ArrayMerge(arayA, arrayB []string) []string {
	result := append(arayA, arrayB...)
	gabung := make(map[string]bool)
	var merged []string

	for _, nama := range result {
		if _, b := gabung[nama]; !b {
			gabung[nama] = true
			merged = append(merged, nama)
		}
	}

	return merged
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	//["sergei", "jin", "steve", "bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	//["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	//["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	//["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))
	//[]
}
