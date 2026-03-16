package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Scan(&name, &age)
	var class string

	for range 5 {
		if age < 12 {
			class = "crianca"
		} else if age < 18 {
			class = "jovem"
		} else if age < 65 {
			class = "adulto"
		} else if age < 1000 {
			class = "idoso"
		} else {
			class = "mumia"
		}
	}

	fmt.Println(name, "eh", class)
}
