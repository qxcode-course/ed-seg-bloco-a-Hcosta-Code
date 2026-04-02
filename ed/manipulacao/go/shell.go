package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	var homens []int

	for _, valor := range vet {
		if valor > 0 {
			homens = append(homens, valor)
		}
	}

	return homens
}

func getCalmWomen(vet []int) []int {
	var mulheres []int

	for _, valor := range vet {
		if valor < 0 {
			tempo := valor * -1
			if tempo < 10 {
				mulheres = append(mulheres, valor)
			}
		}
	}

	return mulheres
}

func sortVet(vet []int) []int {
	sort.Ints(vet)
	return vet
}

func sortStress(vet []int) []int {
	sort.Slice(vet, func(i, j int) bool {
		a1 := vet[i]
		if a1 < 0 {
			a1 = a1 * -1
		}

		a2 := vet[j]
		if a2 < 0 {
			a2 = a2 * -1
		}

		return a1 < a2
	})
	return vet
}

func reverse(vet []int) []int {
	var opost []int
	for i := len(vet) - 1; i >= 0; i-- {
		opost = append(opost, vet[i])
	}
	return opost
}

func repeated(vet []int) []int {
	unico := make(map[int]bool)
	var clone []int

	for _, val := range vet {

		if !unico[val] {
			unico[val] = true
			continue
		} else {
			clone = append(clone, val)
		}
	}
	return clone
}

func unique(vet []int) []int {
	unico := make(map[int]bool)
	var clone []int

	for _, val := range vet {

		if !unico[val] {
			unico[val] = true
			clone = append(clone, val)
		}
	}
	return clone
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}
