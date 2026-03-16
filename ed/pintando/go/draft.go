package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scanf("%f\n%f\n%f\n", &a, &b, &c)
	var result float64

	var p = (a + b + c) / 2
	result = math.Sqrt(p * (p - a) * (p - b) * (p - c))

	fmt.Printf("%.2f\n", result)
}
