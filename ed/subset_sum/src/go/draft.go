package main

import "fmt"

func isSubsetSum(corridas []int, meta int, i int) bool {
	if meta == 0 {
		return true
	}

	if i < 0 {
		return false
	}
	
	caminhoEsquerda := isSubsetSum(corridas, meta-corridas[i], i-1)
	
	caminhoDireita := isSubsetSum(corridas, meta, i-1)

	return caminhoEsquerda || caminhoDireita
}

func metaBatida(corridas []int, meta int) bool {
	ultimoIndice := len(corridas) - 1
	
	return isSubsetSum(corridas, meta, ultimoIndice)
}

func main() {
	var quantidade int
	var meta int

	fmt.Scan(&quantidade, &meta)

	corridas := make([]int, quantidade)

	for i := 0; i < quantidade; i++ {
		fmt.Scan(&corridas[i])
	}

	fmt.Println(metaBatida(corridas, meta))
}