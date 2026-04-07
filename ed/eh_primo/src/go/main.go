package main

import "fmt"

func eh_primo(x int, i int) bool {
	if x <= 1 {
		return false
	}
	if i == 1 {
		return true // Chegamos ao 1 sem encontrar divisores, então é primo
	}

	if x%i == 0 {
		return false
	}

	return eh_primo(x, i-1)
}

func main() {
	var x int
	fmt.Scan(&x)

	if eh_primo(x, x-1) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
