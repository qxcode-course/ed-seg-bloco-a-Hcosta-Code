package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {

	if len(vet) == 0 {
		return "[]"
	}

	var proces func(temp []int) string
	proces = func(temp []int) string {
		if len(temp) == 1 {
			return strconv.Itoa(temp[0])
		}

		return strconv.Itoa(temp[0]) + ", " + proces(temp[1:])

	}

	return "[" + proces(vet) + "]"
}

func tostrrev(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}

	var proces func(temp []int) string
	proces = func(temp []int) string {
		if len(temp) == 1 {
			return strconv.Itoa(temp[0])
		}

		return strconv.Itoa(temp[len(temp)-1]) + ", " + proces(temp[:len(temp)-1])
	}
	return "[" + proces(vet) + "]"
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	var opost []int
	for i := len(vet) - 1; i >= 0; i-- {
		opost = append(opost, vet[i])
	}
	copy(vet, opost)
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	var temp int
	for _, val := range vet {
		temp += val
	}

	return temp
}

// mult: produto dos elementos do slice
func mult(vet []int) int {

	if len(vet) == 0 {
		return 1 // Ou 1, dependendo da regra de negócio (geralmente 0 para vazio)
	}

	var proces func(temp []int) int
	proces = func(temp []int) int {
		if len(temp) == 1 {
			return temp[0]
		}
		return temp[0] * proces(temp[1:])
	}
	return proces(vet)
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	// Caso para lista vazia (conforme sua imagem, retorna -1)
	if len(vet) == 0 {
		return -1
	}

	var proces func(inicio int) int
	proces = func(inicio int) int {
		// CASO BASE: Se chegamos no último elemento,
		// o índice do menor é ele mesmo.
		if inicio == len(vet)-1 {
			return inicio
		}

		// 1. Pega o índice do menor elemento que está lá na frente
		indiceMenorDoResto := proces(inicio + 1)

		// 2. Compara o VALOR do elemento atual com o VALOR do menor encontrado lá na frente
		// Usamos vet[inicio] e vet[indiceMenorDoResto] para comparar os números reais
		if vet[inicio] <= vet[indiceMenorDoResto] {
			return inicio
		}
		return indiceMenorDoResto
	}

	return proces(0)
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
