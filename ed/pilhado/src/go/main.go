package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	r, c int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	if !scanner.Scan() {
		return
	}

	var R, C int
	fmt.Sscanf(scanner.Text(), "%d %d", &R, &C)

	maze := make([][]byte, R)
	var start Pos

	for i := 0; i < R; i++ {
		scanner.Scan()
		maze[i] = []byte(scanner.Text())

		for J := 0; J < len(maze[i]); J++ {
			if maze[i][J] == 'I' {
				start = Pos{i, J}
			}
		}
	}

	visited := make([][]bool, R)
	for i := range visited {
		visited[i] = make([]bool, C)
	}

	path := []Pos{start}
	visited[start.r][start.c] = true

	dr := []int{-1, 1, 0, 0}
	dc := []int{0, 0, -1, 1}

	for len(path) > 0 {
		curr := path[len(path)-1]

		if maze[curr.r][curr.c] == 'F' {
			break
		}

		foundValidNeighbor := false
		for i := 0; i < 4; i++ {
			nr, nc := curr.r+dr[i], curr.c+dc[i]

			if nr >= 0 && nr < R && nc >= 0 && nc < len(maze[nr]) {
				if !visited[nr][nc] && maze[nr][nc] != '#' {
					visited[nr][nc] = true
					path = append(path, Pos{nr, nc})
					foundValidNeighbor = true
					break
				}
			}

		}

		if !foundValidNeighbor {
			path = path[:len(path)-1]
		}

	}

	for _, p := range path {
		maze[p.r][p.c] = '.'

	}

	for i := 0; i < R; i++ {
		fmt.Println(string(maze[i]))
	}
}
