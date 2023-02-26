package main

import "fmt"

func Mapping(slice []string) map[string]int {

	counts := make(map[string]int)

	for _, a := range slice {

		if _, ok := counts[a]; ok {
			counts[a]++
		} else {
			counts[a] = 1
		}
	}

	return counts
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})) //map[adi:1 asd:2 qwe:3]
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))                      //map[asd:2 qwe:1]
	fmt.Println(Mapping([]string{}))                                         //map[]
}
