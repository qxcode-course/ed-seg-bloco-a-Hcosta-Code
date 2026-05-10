package main

import "fmt"

func recur(n int, k, valor float64) {
	if n == 0 {
		fmt.Printf("%.2f\n", valor)
		return
	}

	valor = valor + k
	valor = valor / 2
	recur(n-1, k, valor)

}

func main() {
	var n int
	var valor, k float64
	fmt.Scan(&n, &k)
	recur(n, k, valor)
}
