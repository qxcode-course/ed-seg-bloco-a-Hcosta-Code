package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanln(&N)

	type Player struct {
		A int
		B int
	}

	pontos := make([]Player, N)

    for var i int = 0; i < N; i++ {
        fmt.Scanln(pontos[i].A, pontos[i].B)
    }

}
