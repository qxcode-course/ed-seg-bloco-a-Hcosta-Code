package main

import (
	"bufio"
	"fmt"
	"os"
)

func ehPar(abertura, fechamento rune) bool {
	if abertura == '(' && fechamento == ')' {
		return true
	}
	if abertura == '[' && fechamento == ']' {
		return true
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	texto := scanner.Text()

	var pilha []rune

	for _, char := range texto {
		if char == '(' || char == '[' {
			pilha = append(pilha, char)
		} else if char == ')' || char == ']' {
			if len(pilha) == 0 {
				fmt.Println("nao balanceado")
				return
			}

			topo := pilha[len(pilha)-1]

			if ehPar(topo, char) {
				pilha = pilha[:len(pilha)-1]
			} else {
				fmt.Println("nao balanceado")
				return
			}
		}
	}

	if len(pilha) == 0 {
		fmt.Println("balanceado")
	} else {
		fmt.Println("nao balanceado")
	}
}
