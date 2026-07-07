package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func create(parts *[]string) *Node {
	if len(*parts) == 0 {
		return nil
	}

	valStr := (*parts)[0]
	*parts = (*parts)[1:]

	if valStr == "#" {
		return nil
	}

	val, _ := strconv.Atoi(valStr)
	node := &Node{Value: val}

	node.Left = create(parts)
	node.Right = create(parts)

	return node
}

func BShow(node *Node, history string) {
	if node != nil && (node.Left != nil || node.Right != nil) {
		BShow(node.Left, history+"l")
	}

	for i := 0; i < len(history)-1; i++ {
		if history[i] != history[i+1] {
			
			fmt.Print("│   ")
		} else {
			
			fmt.Print("    ")
		}
	}
	if history != "" {
		if history[len(history)-1] == 'l' {
			fmt.Print("╭───")
		} else {
			fmt.Print("╰───")
		}
	}
	if node == nil {
		fmt.Println("#")
		return
	}
	fmt.Println(node.Value)
	if node.Left != nil || node.Right != nil {
		BShow(node.Right, history+"r")
	}
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var parts []string

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)
		parts = append(parts, tokens...)
	}


	for len(parts) > 0 {
		root := create(&parts)
		BShow(root, "")
	}
}