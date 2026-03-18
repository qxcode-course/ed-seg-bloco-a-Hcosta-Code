package main

import "fmt"

func main() {
	qtd_album := 0
	qtd_fig := 0
	fmt.Scan(&qtd_album, &qtd_fig)
	album := make([]int, qtd_fig)

	unicos := make(map[int]bool)
	repetidos := make([]int, 0, qtd_fig)

	for i := range album {
		fmt.Scan(&album[i])
	}

	for _, fig := range album {
		if unicos[fig] {
			repetidos = append(repetidos, fig)
		} else {
			unicos[fig] = true
		}

	}

	for i := range repetidos {
		if i != 0 {
			fmt.print(" ")
		}
		fmt.Printf("%v ", valor)
	}
	fmt.Print("\n")

}
