package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Set {
	return &Set{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (v *Set) Reserve(newCapacity int) {
	if newCapacity <= v.capacity {
		return
	}

	newData := make([]int, newCapacity)
	for i := 0; i < v.size; i++ {
		newData[i] = v.data[i]
	}

	v.data = newData
	v.capacity = newCapacity
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	var v *Set

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)

				existe := false
				for i := 0; i < v.size; i++ {
					if v.data[i] == value {
						existe = true
						break
					}
				}

				if !existe {
					if v.size == v.capacity {
						novaCap := v.capacity * 2
						if novaCap == 0 {
							novaCap = 1
						}
						v.Reserve(novaCap)
					}

					posicao := 0
					for posicao < v.size && v.data[posicao] < value {
						posicao++
					}

					for i := v.size; i > posicao; i-- {
						v.data[i] = v.data[i-1]
					}

					v.data[posicao] = value
					v.size++
				}
			}
		case "show":
			fmt.Print("[")
			for i := 0; i < v.size; i++ {
				fmt.Print(v.data[i])
				if i < v.size-1 {
					fmt.Print(", ")
				}
			}
			fmt.Println("]")
		case "erase":
			alvo, _ := strconv.Atoi(parts[1])

			indiceEncontrado := -1
			for i := 0; i < v.size; i++ {
				if v.data[i] == alvo {
					indiceEncontrado = i
					break 
				}
			}

			if indiceEncontrado != -1 {
				for i := indiceEncontrado; i < v.size-1; i++ {
					v.data[i] = v.data[i+1]
				}

				v.size--
			} else {
				fmt.Println("value not found")
			}
	
		case "contains":
			valorProcurado, _ := strconv.Atoi(parts[1])
			achou := false

			for i := 0; i < v.size; i++ {
				if v.data[i] == valorProcurado {
					achou = true
					break
				}
			}
			fmt.Println(achou)
		case "clear":
			v.size = 0
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
