package main

import (
	"fmt"
)

func calc(n, k int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return calc(n-1, k) + k*calc(n-2, k)
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	fmt.Println(calc(n, k))
}