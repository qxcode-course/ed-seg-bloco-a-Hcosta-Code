package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func getNeig(p Pos) []Pos {
	return []Pos{{p.l, p.c - 1}, {p.l - 1, p.c}, {p.l, p.c + 1}, {p.l + 1, p.c}}
}

func inside(grid [][]rune, p Pos) bool {
	return !(p.l < 0 || p.l >= len(grid) || p.c < 0 || p.c >= len(grid[0]))
}

func match(grid [][]rune, p Pos, value rune) bool {
	return inside(grid, p) && grid[p.l][p.c] == value
}

// Função recursiva que tenta encontrar o caminho do início ao fim
func search(grid [][]rune, current Pos, endPos Pos) bool {
    // Caso base: se a posição atual for o destino
    if current == endPos {
        grid[current.l][current.c] = '.'
        return true
    }

    // Se não for um espaço vazio (ou seja, se for parede '#' ou já visitado '.'), para
    if !match(grid, current, ' ') {
        return false
    }

    // Marca a posição atual como parte do caminho
    grid[current.l][current.c] = '.'

    // Tenta ir para cada um dos vizinhos
    for _, vizinho := range getNeig(current) {
        if search(grid, vizinho, endPos) {
            return true // Se algum vizinho chegar ao fim, retorna true
        }
    }

    // Se nenhum vizinho levar ao destino, desmarca o caminho (Backtracking)
    // Se o problema pedir apenas para marcar onde ele "tentou" ir, remova a linha abaixo
    grid[current.l][current.c] = ' ' 
    
    return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nl_nc := scanner.Text()
	var nl, nc int
	fmt.Sscanf(nl_nc, "%d %d", &nl, &nc)
	grid := make([][]rune, nl)

	// Lê a gridriz
	for i := range nl {
		scanner.Scan()
		grid[i] = []rune(scanner.Text())
	}

	// Procura posições de início e endPos e conserta para _
	var startPos, endPos Pos
	for l := range nl {
		for c := range nc {
			if grid[l][c] == 'I' {
				grid[l][c] = ' '
				startPos = Pos{l, c}
			}
			if grid[l][c] == 'F' {
				grid[l][c] = ' '
				endPos = Pos{l, c}
			}
		}
	}

	search(grid, startPos, endPos)

	// Imprime o labirinto final
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
