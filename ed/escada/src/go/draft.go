package main

import "fmt"

func calculando(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	return calculando(n-1) + calculando(n-3)
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(calculando(n))
}
