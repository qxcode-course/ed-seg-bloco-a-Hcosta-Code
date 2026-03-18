package main

import "fmt"

func main() {
	var G int
	var dir string
	fmt.Scan(&G, &dir)

	type loc struct {
		X int
		Y int
	}

	//Lendo posição de cada gomo
	gomo := make([]loc, G)
	for i := 0; i < G; i++ {
		fmt.Scan(&gomo[i].X, &gomo[i].Y)
	}

	//Mudando posições corpo
	for i := G - 1; i > 0; i-- {
		gomo[i].X = gomo[i-1].X
		gomo[i].Y = gomo[i-1].Y
	}

	//Mudando posições cabeça
	if dir == "L" {
		gomo[0].X = gomo[0].X - 1
	} else if dir == "R" {
		gomo[0].X = gomo[0].X + 1
	} else if dir == "U" {
		gomo[0].Y = gomo[0].Y - 1
	} else if dir == "D" {
		gomo[0].Y = gomo[0].Y + 1
	}

	//Printando os gomos
	for i := 0; i < G; i++ {
		fmt.Println(gomo[i].X, gomo[i].Y)
	}

}
