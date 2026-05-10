package main

import "fmt"

func recur(n int) {
	if n == 1 {
		fmt.Println("1^2 = 1")
		return
	}

	fmt.Print(n, "^2 = ", n-1, "^2 + 2*", n-1, " + 1 = ?")
	fmt.Println()
	recur(n - 1)

	fmt.Print(n, "^2 = ", n-1, "^2 + 2*", n-1, " + 1 = ")
	fmt.Print(n * n)
	fmt.Println()

}

func main() {
	var n int
	fmt.Scan(&n)
	recur(n)
}
