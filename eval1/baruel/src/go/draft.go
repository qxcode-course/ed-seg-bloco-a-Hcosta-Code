package main

import "fmt"

func main() {
	var quant_t, quant_b int
	fmt.Scan(&quant_t, &quant_b)

	album_baru := make(map[int]bool)
	var reps []int
	var anterior = -1

	for i := 0; i < quant_b; i++ {
		var temp int
		fmt.Scan(&temp)

		if temp == anterior {
			reps = append(reps, temp)
		}

		album_baru[temp] = true
		anterior = temp

	}

	if len(reps) == 0 {
		fmt.Print("N\n")
	} else {
		for i := 0; i < len(reps); i++ {
			if i == len(reps)-1 {
				fmt.Print(reps[i])
				continue
			} else {
				fmt.Print(reps[i])
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}

	var falt []int
	for i := 1; i <= quant_t; i++ {
		if !album_baru[i] {
			falt = append(falt, i)
		}
	}

	if len(falt) == 0 {
		fmt.Print("N\n")
	} else {
		for i := 0; i < len(falt); i++ {
			if i == len(falt)-1 {
				fmt.Print(falt[i])
				continue
			} else {
				fmt.Print(falt[i])
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")

	}

}
