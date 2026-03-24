package main

import "fmt"

func imprimir_fila(fila []int, espada int) {
	fmt.Print("[ ")
	for i := 0; i < len(fila); i++ {
		if i == espada {
			fmt.Printf("%v>", fila[i])
			if i < len(fila)-1 {
				fmt.Print(" ")
			}
			continue
		} else {
			fmt.Printf("%v", fila[i])
		}
		if i < len(fila)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print(" ]\n")
}

func tirar_morto(fila []int, morto int) []int {
	return append(fila[:morto], fila[morto+1:]...)
}

func main() {
	var N, E int
	fmt.Scan(&N, &E)

	var fila []int
	for i := 0; i < N; i++ {
		fila = append(fila, i+1)
	}

	espada := E - 1

	for len(fila) > 1 {
		imprimir_fila(fila, espada)

		morto := (espada + 1) % len(fila)
		fila = tirar_morto(fila, morto)

		espada = morto % len(fila)
	}

	imprimir_fila(fila, espada)
}
