package main

import "fmt"

type pr struct {
	quociente int
	resto     int
}

func main() {
	var numb int
	fmt.Scan(&numb)

	var princ func(vet int) []pr
	princ = func(vet int) []pr {
		if vet == 0 {
			return []pr{}
		}

		atual := pr{quociente: vet / 2, resto: vet % 2}

		return append(princ(vet/2), atual)
	}

	vals := princ(numb)

	for i := 0; i < len(vals); i++ {
		fmt.Printf("%d %d\n", vals[i].quociente, vals[i].resto)
	}
}
