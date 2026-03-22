package main

import "fmt"

func main() {
	var album_t, figs_baru int
	fmt.Scan(&album_t, &figs_baru)

	album_baru := make(map[int]bool)
	var reps []int
	var anterior int = -1

	//Verificar repetidas e adicionar os valores individuais
	for i := 0; i < figs_baru; i++ {
		var temp int
		fmt.Scan(&temp)

		if temp == anterior {
			reps = append(reps, temp)
		}

		album_baru[temp] = true
		anterior = temp
	}

	//printar as repetidas
	if len(reps) == 0 {
		fmt.Print("N\n")
	} else {
		for i, v := range reps {
			fmt.Printf("%d", v)
			if i < len(reps)-1 {
				fmt.Print(" ")
			}

		}
		fmt.Print("\n")
	}

	//Criando uma lista das que faltam
	var falt []int
	for i := 1; i <= album_t; i++ {
		if !album_baru[i] {
			falt = append(falt, i)
		}
	}

	//printar as que faltam
	if len(falt) == 0 {
		fmt.Print("N\n")
	} else {
		for i, v := range falt {
			fmt.Printf("%d", v)
			if i < len(falt)-1 {
				fmt.Print(" ")
			}

		}
		fmt.Print("\n")
	}
}
