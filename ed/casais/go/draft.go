package main

import (
	"fmt"
)

func main() {
	var Q int
	fmt.Scan(&Q)

	// O mapa guarda: [tamanho da bota] -> [quantidade disponível no estoque]
	estoque := make(map[int]int)
	casais := 0

	for i := 0; i < Q; i++ {
		var bota int
		fmt.Scan(&bota)

		// O par de uma bota é o valor oposto (ex: par de 37 é -37)
		parEsperado := -bota

		// Verificamos se o par já existe no estoque
		if estoque[parEsperado] > 0 {
			// Achei um par!
			casais++
			// Removemos uma unidade do par do estoque, pois ele já "casou"
			estoque[parEsperado]--
		} else {
			// Se não tem par esperando, guardamos esta bota no estoque
			estoque[bota]++
		}
	}

	fmt.Printf("%d\n", casais)
}
