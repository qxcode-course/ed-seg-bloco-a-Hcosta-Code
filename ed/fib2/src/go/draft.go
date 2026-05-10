package main

import (
	"fmt"
)

func calc(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return calc(n-1) + calc(n-2)
	}
	return calc(n-2) + calc(n-3)
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(calc(n))
}
