package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	fila := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&fila[i])
	}

	var M int
	fmt.Scan(&M)
	deix_fila := make(map[int]bool)
	for i := 0; i < M; i++ {
		var temp int
		fmt.Scan(&temp)
		deix_fila[temp] = true
	}

	var final []int
	for i := 0; i < N; i++ {
		if deix_fila[fila[i]] {
			continue
		} else {
			final = append(final, fila[i])
		}

	}

	for i := 0; i < len(final); i++ {
		fmt.Print(final[i], " ")
		if i == len(final)-1 {
			fmt.Println()
		}
	}
}
