package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	
	adj := make([][]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue 
			}
			
			dx := int64(bombs[i][0] - bombs[j][0])
			dy := int64(bombs[i][1] - bombs[j][1])
			r := int64(bombs[i][2])
			
			if dx*dx + dy*dy <= r*r {
				adj[i] = append(adj[i], j)
			}
		}
	}

	maxDetonated := 0

	for i := 0; i < n; i++ {
		visited := make([]bool, n)
		
		count := dfs(i, adj, visited)
		
		if count > maxDetonated {
			maxDetonated = count
		}
	}

	return maxDetonated
}

func dfs(u int, adj [][]int, visited []bool) int {
	visited[u] = true
	count := 1 
	
	for _, v := range adj[u] {
		if !visited[v] {
			count += dfs(v, adj, visited)
		}
	}
	
	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	
	nl, _ := strconv.Atoi(parts[0]) 

	bombs := make([][]int, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			return
		}
		tokens := strings.Fields(scanner.Text())
		bomb := make([]int, 3)
		for j := 0; j < 3 && j < len(tokens); j++ {
			bomb[j], _ = strconv.Atoi(tokens[j])
		}
		bombs[i] = bomb
	}

	fmt.Println(maximumDetonation(bombs))
}