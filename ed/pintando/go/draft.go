package main

import (
	"fmt"
	"math"
)

func main() {
	var l3, l1, l2 float64
	fmt.Scan(&l1, &l2, &l3)

	var p float64
	p = (l1 + l2 + l3) / 2

	var result float64
	result = math.Sqrt(p * (p - l1) * (p - l2) * (p - l3))

	fmt.Printf("%.2f\n", result)
}
