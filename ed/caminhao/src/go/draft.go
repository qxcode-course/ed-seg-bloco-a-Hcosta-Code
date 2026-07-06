package main

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	gas := make([]int, n)
	dist := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&gas[i], &dist[i])
	}

	start := 0
	current_surplus := 0
	total_surplus := 0

	for i := 0; i < n; i++ {
		diff := gas[i] - dist[i]

		total_surplus += diff

		current_surplus += diff

		if current_surplus < 0 {
			start = i + 1
			current_surplus = 0
		}
	}

	if total_surplus >= 0 {
		fmt.Println(start)
	} else {
		fmt.Println("-1")
	}
}
