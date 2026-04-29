package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	var ms int
	var slicev []int

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = value
			slicev = make([]int, 0, ms)
		case "insert":
			//inserindo
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				slicev = append(slicev, value)
			}
			sort.Ints(slicev)

		case "show":
			fmt.Print("[")
			for index, value := range slicev {
				if index == len(slicev)-1 {
					fmt.Print(value)
				} else {
					fmt.Print(value, ", ")
				}
			}
			fmt.Println("]")
		case "erase":
			value, _ := strconv.Atoi(args[1])
			for indice, number := range slicev {
				if number == value {
					slicev = append(slicev[:indice], slicev[indice+1:]...)
					break	
				} else if indice == len(slicev)-1{
					fmt.Println("value not found")
				}
			}

		case "contains":
			value, _ := strconv.Atoi(args[1])
			temp := false
			for _, number := range slicev {
				if number == value {
					temp = true
					break
				}
			}
			fmt.Println(temp)
		case "count":
			value, _ := strconv.Atoi(args[1])
			var count int
			for _, number := range slicev {
				if number == value {
					count++
				}
			}
			fmt.Println(count)
		case "unique":
			var count int
			var number int
			for index, _ := range slicev {
				if slicev[index] == number {
					continue
				} else{
					count ++
					number = slicev[index]
				}
			}
			fmt.Println(count)

		case "clear":
			NewSliceV := make([]int, 0, ms)
			slicev = NewSliceV

		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
