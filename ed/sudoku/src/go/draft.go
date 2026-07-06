package main

import "fmt"

func resolver(matriz [][]rune, index int) bool {
	N := len(matriz)

	if index == N*N {
		return true
	}

	l := index / N
	c := index % N

	if matriz[l][c] != '.' {
		return resolver(matriz, index+1)
	}

	limite := rune('0' + N)

	for val := rune('1'); val <= limite; val++ {
		if isValid(matriz, l, c, val) {
			matriz[l][c] = val

			if resolver(matriz, index+1) {
				return true
			}
			matriz[l][c] = '.'
		}

	}
	return false
}

func isValid(matriz [][]rune, r int, c int, val rune) bool {
	N := len(matriz)

	for i := 0; i < N; i++ {
		if matriz[r][i] == val || matriz[i][c] == val {
			return false
		}
	}

	blockSize := 2
	if N == 9 {
		blockSize = 3
	}

	boxRow := (r / blockSize) * blockSize
	boxCol := (c / blockSize) * blockSize

	for i := 0; i < blockSize; i++ {
		for j := 0; j < blockSize; j++ {
			if matriz[boxRow+i][boxCol+j] == val {
				return false
			}
		}
	}

	return true
}

func main() {
	var N int
	if _, err := fmt.Scan(&N); err != nil {
		return
	}
	matriz := make([][]rune, N)
	for i := 0; i < N; i++ {
		matriz[i] = make([]rune, N)
	}

	for i := 0; i < N*N; {
		var s string
		if _, err := fmt.Scan(&s); err != nil {
			break
		}
		for _, char := range s {
			r := i / N
			c := i % N
			matriz[r][c] = char
			i++
		}
	}
	resolver(matriz, 0)

	for i := 0; i < N; i++ {
		for J := 0; J < N; J++ {
			fmt.Print(string(matriz[i][J]))
		}
		fmt.Println()
	}
	
}
