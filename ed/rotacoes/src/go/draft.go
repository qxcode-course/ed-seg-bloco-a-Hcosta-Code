package main

import "fmt"

func main() {
	var t, r int
	fmt.Scan(&t, &r)

	vetor := make([]int, t)
	for i := 0; i < t; i++ {
		vetor[i] = i + 1
	}

	vetor2 := make([]int, t)

	//logica
	for i := 0; i < t; i++ {
		novaPo := (i + r) % t
		vetor2[novaPo] = vetor[i]
	}

	fmt.Printf("[")
	for i := 0; i < t; i++ {
		if i == t-1 {
			fmt.Printf(" ")
			fmt.Print(vetor2[i])
			fmt.Printf(" ")
			continue
		}
		fmt.Printf(" ")
		fmt.Print(vetor2[i])
	}
	fmt.Printf("]\n")

}
