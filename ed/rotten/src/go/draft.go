package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct{
    r, c int
}

func orangesRotting(grid [][]int) int{
    if len(grid) == 0{
        return 0
    }

    m,n := len(grid), len(grid[0])
    var queue []Pos
    freshCout := 0

    for i := 0; i < m;i++{
        for j := 0; j <n;j++{
            if grid[i][j] == 2{
                queue = append(queue, Pos{r: i,c: j})
            } else if grid[i][j] == 1{
                freshCout++
            }
        }
    }
    if freshCout == 0{
        return 0
    }

    minutes := 0
    dirs := []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

    for len(queue) > 0 && freshCout > 0{
        minutes++

        size := len(queue)

        for i := 0; i < size; i++{
            curr := queue[0]
            queue = queue[1:]

            for _, d := range dirs {
                nr, nc := curr.r+d.r, curr.c+d.c

                if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == 1 {
					grid[nr][nc] = 2               
					freshCout--                  
					queue = append(queue, Pos{nr, nc})
				}
            }
        }
    }
    if freshCout == 0{
        return minutes
    }
    return -1
}

func main(){
    scanner := bufio.NewScanner(os.Stdin)
    if !scanner.Scan(){
        return
    }
    parts := strings.Fields(scanner.Text())
    if len(parts) < 2{
        return
    }

    nl, _ := strconv.Atoi(parts[0])
    nc, _ := strconv.Atoi(parts[1])

    grid := make([][]int, nl)
    for i := 0; i < nl; i++ {
        if !scanner.Scan() {
            return
        }
        tokens := strings.Fields(scanner.Text())
        row := make([]int, nc)
        for j := 0; j < nc && j < len(tokens); j++{
            row[j], _ = strconv.Atoi(tokens[j])
        }
        grid[i] = row
    }
    fmt.Println(orangesRotting(grid))

}





