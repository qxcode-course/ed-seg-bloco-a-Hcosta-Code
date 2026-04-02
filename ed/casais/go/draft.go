package main

import (
	"fmt"
)

func main() {
	var Q int
	fmt.Scan(&Q)

	estoque := make(map[int]int)
	casais := 0

	for i := 0; i < Q; i++ {
		var bota int
		fmt.Scan(&bota)

		parEsperado := -bota

		if estoque[parEsperado] > 0 {
			// Achei um par!
			casais++
			estoque[parEsperado]--
		} else {
			estoque[bota]++
		}
	}

	fmt.Printf("%d\n", casais)
}
