package main

import (
	"fmt"
)

func main() {
	var H, P, F, D int
	fmt.Scanln(&H, &P, &F, &D)

	for i := 0; i <= 15; i++ {
		if D == -1 {
			F++
		} else {
			F--
		}

		if F > 15 {
			F = 0
		} else if F < 0 {
			F = 15
		}

		if F == H {
			fmt.Print("N\n")
			break
		}
		if F == P {
			fmt.Print("S\n")
			break
		}
	}
}
