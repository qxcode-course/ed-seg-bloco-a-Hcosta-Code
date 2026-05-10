package main

import "fmt"

func eh_primo(n int, i int) bool {
	if n <= 2 {
		return n == 2
	}
	if n%i == 0 {
		return false
	}
	if i*i > n {
		return true
	}
	return eh_primo(n, i+1)
}

func buscarNPrime(n, atual, encontrados int) int {
	if eh_primo(atual, 2) {
		if encontrados+1 == n {
			return atual
		}
		return buscarNPrime(n, atual+1, encontrados+1)
	}

	return buscarNPrime(n, atual+1, encontrados)
}

func main() {
	var num int
	fmt.Scan(&num)

	fmt.Println(buscarNPrime(num, 2, 0))
}
