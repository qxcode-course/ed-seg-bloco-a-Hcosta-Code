package main

import (
	"fmt"
)

type cod struct {
	str string
	qtd int
}

func main() {
	var entrada, verifica int
	fmt.Scan(&entrada)
	entrando := make(map[string]int)
	for i := 0; i < entrada; i++ {
		var temp string
		fmt.Scan(&temp)
		entrando[temp]++
	}

	fmt.Scan(&verifica)
	verificando := make([]cod, verifica)
	for i := 0; i < verifica; i++ {
		fmt.Scan(&verificando[i].str)
		verificando[i].qtd = 0
	}

	for i := 0; i < verifica; i++ {
		codigoProcurado := verificando[i].str
		verificando[i].qtd = entrando[codigoProcurado]
	}

	for i := 0; i < verifica; i++ {
		if i < verifica-1 {
			fmt.Print(verificando[i].qtd, " ")
		} else {
			fmt.Print(verificando[i].qtd)
		}
	}
	fmt.Println()
}
