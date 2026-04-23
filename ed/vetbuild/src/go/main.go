package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
}

func (v *Vector) Reserve(newCapacity int) {
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

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	var v *Vector

	// v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
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
			capac, _ := strconv.Atoi(parts[1])
			v = NewVector(capac)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)

				if v.size == v.capacity {
					v.Reserve(v.capacity * 2)
				}

				v.data[v.size] = value
				v.size++

			}
		case "show":
			fatia := v.data[:v.size]
			conte := Join(fatia, ", ")
			fmt.Printf("[%s]\n", conte)
		case "status":
			fmt.Print("size:")
			fmt.Print(v.size)
			fmt.Print(" ")
			fmt.Print("capacity:")
			fmt.Print(v.capacity)
			fmt.Print("\n")
		case "pop":
			if v.size-1 == -1 {
				fmt.Printf("vector is empty\n")
			} else {
				v.size = v.size - 1
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])

			if v.size == v.capacity {
				v.Reserve(v.capacity * 2)
			}

			for i := v.size; i > index; i-- {
				v.data[i] = v.data[i-1]
			}

			v.data[index] = value
			v.size++
		case "erase":
			index, _ := strconv.Atoi(parts[1])

			if index >= 0 && index < v.size {
				for i := index; i < v.size-1; i++ {
					v.data[i] = v.data[i+1]
				}

				v.size = v.size - 1
			} else {
				fmt.Println("index out of range")
			}
		case "indexOf":
			index, _ := strconv.Atoi(parts[1])
			acho := false
			for i := 0; i < v.size; i++ {
				temp := v.data[i]

				if temp == index {
					fmt.Print(i)
					fmt.Printf("\n")
					acho = true
					break
				}
			}
			if acho == false {
				fmt.Printf("-1\n")
			}

		case "contains":
			index, _ := strconv.Atoi(parts[1])
			for i := 0; i < v.size; i++ {
				temp := v.data[i]
				if temp == index {
					fmt.Print(true, "\n")
				}
			}

		case "clear":
			v.size = 0
		case "capacity":
			fmt.Println(v.capacity)
		case "get":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)

				if value >= 0 && value < v.size {
					temp := v.data[value]
					fmt.Print(temp)
					fmt.Printf("\n")
				} else {
					fmt.Println("index out of range")
				}

			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])

			if index >= 0 && index < v.size {
				v.data[index] = value
			} else {
				fmt.Println("index out of range")
			}
		case "reserve":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Reserve(value)
			}
		case "slice":
			start, _ := strconv.Atoi(parts[1])
			end, _ := strconv.Atoi(parts[2])

			if start < 0 {
				start = v.size + start
			}
			if end < 0 {
				end = v.size + end
			}

			if start >= 0 && end <= v.size && start <= end {
				fatia := v.data[start:end]
				fmt.Printf("[%s]\n", Join(fatia, ", "))
			} else {
				fmt.Println("fail: indice invalido")
			}
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
