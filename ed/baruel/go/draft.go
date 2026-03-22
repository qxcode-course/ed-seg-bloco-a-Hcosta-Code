package main

import "fmt"

func main() {
	var total, possui int
	fmt.Scan(&total, &possui)

	// Usaremos um slice para as repetidas e um mapa para marcar o que ele TEM
	album := make(map[int]bool)
	var repetidas []int

	var anterior int = -1 // Começamos com -1 pois as figurinhas são de 1 em diante

	for i := 0; i < possui; i++ {
		var atual int
		fmt.Scan(&atual)

		// Lógica das Repetidas:
		// Como a entrada é crescente, se o atual for igual ao anterior, é repetida!
		if atual == anterior {
			repetidas = append(repetidas, atual)
		}

		// Marca no nosso mapa (álbum) que ele tem essa figurinha
		album[atual] = true
		anterior = atual
	}

	// --- SAÍDA 1: Repetidas ---
	if len(repetidas) == 0 {
		fmt.Println("N")
	} else {
		for i, v := range repetidas {
			fmt.Print(v)
			if i < len(repetidas)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	// --- SAÍDA 2: Faltando ---
	var faltando []int
	// Percorremos de 1 até o Total do álbum
	for i := 1; i <= total; i++ {
		if !album[i] { // Se NÃO está no mapa, está faltando
			faltando = append(faltando, i)
		}
	}

	if len(faltando) == 0 {
		fmt.Println("N")
	} else {
		for i, v := range faltando {
			fmt.Print(v)
			if i < len(faltando)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
