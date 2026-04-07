package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

func formatar(vetor []int) string {
	var s []string
	for _, v := range vetor {
		s = append(s, strconv.Itoa(v))
	}
	
	return "[ " + strings.Join(s, " ") + " ]"
}
func processa(vetor []int) {

	if len(vetor) == 1 {
		fmt.Println(formatar(vetor))
		return
	}

	temp := make([]int, len(vetor)-1)

	for i := 0; i < len(vetor)-1; i++ {
		temp[i] = vetor[i] + vetor[i+1]
	}

	processa(temp)

	fmt.Println(formatar(vetor))
}

func main() {
	var vetor []int
	var n int

	for {
		_, err := fmt.Scan(&n)
		if err == io.EOF {
			break
		}
		vetor = append(vetor, n)
	}

	// Só processa se o vetor não estiver vazio
	if len(vetor) > 0 {
		processa(vetor)
	}
}
