package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l int
	c int
}

func burnTrees(grid [][]rune, l, c int) {
	nl := len(grid)
	if nl == 0 {
		return
	}
	nc := len(grid[0])

	stack := []Pos{{l, c}}

	for len(stack) > 0 {
		atual := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if atual.l >= 0 && atual.l < nl && atual.c >= 0 && atual.c < nc {
			if grid[atual.l][atual.c] == '#' {

				grid[atual.l][atual.c] = 'o'

				stack = append(stack, Pos{atual.l - 1, atual.c}) // Vizinho de Cima
				stack = append(stack, Pos{atual.l + 1, atual.c}) // Vizinho de Baixo
				stack = append(stack, Pos{atual.l, atual.c - 1}) // Vizinho da Esquerda
				stack = append(stack, Pos{atual.l, atual.c + 1}) // Vizinho da Direita
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
