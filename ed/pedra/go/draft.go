package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanln(&N)
	var desclass = 0
	var indice int = 0

	type Player struct {
		A   int
		B   int
		Dif int
	}

	pontos := make([]Player, N)

	for i := 0; i < N; i++ {
		fmt.Scanln(&pontos[i].A, &pontos[i].B)
		if pontos[i].A >= 10 && pontos[i].B >= 10 {
			pontos[i].Dif = pontos[i].A - pontos[i].B
		} else if pontos[i].A < 10 || pontos[i].B < 10 {
			desclass++
			pontos[i].Dif = 10000
		} else {
			pontos[i].Dif = 10000
		}
	}

	for i := 0; i < N; i++ {
		if pontos[i].Dif < 0 {
			pontos[i].Dif = pontos[i].Dif * -1
		}
	}

	for i := 1; i < N; i++ {
		if pontos[indice].Dif == pontos[i].Dif {
			continue
		} else if pontos[indice].Dif > pontos[i].Dif {
			indice = i
		}
	}

	if desclass == N {
		fmt.Print("sem ganhador\n")
	} else {
		fmt.Printf("%v\n", indice)
	}
}
