package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getPrecedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "^":
		return 3
	}
	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	line := scanner.Text()
	tokens := strings.Fields(line)

	var output []string
	var operator []string

	for _, token := range tokens {
		prec := getPrecedence(token)

		if prec > 0 {
			for len(operator) > 0 {
				topOperator := operator[len(operator)-1]
				if getPrecedence(topOperator) >= prec {
					output = append(output, topOperator)
					operator = operator[:len(operator)-1]
				} else {
					break
				}
			}
			operator = append(operator, token)

		} else {
			output = append(output, token)
		}
	}
	for i := len(operator) - 1; i >= 0; i-- {
		output = append(output, operator[i])
	}

	fmt.Println(strings.Join(output, " "))

}
