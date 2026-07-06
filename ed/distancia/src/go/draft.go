package main

import (
	"fmt"
)

func solve(s []byte, L int, index int) bool {
	if index == len(s) {
		return true
	}

	if s[index] != '.' {
		return solve(s, L, index+1)
	}

	for d := 0; d <= L; d++ {
		charD := byte('0' + d)
		valid := true

		for i := 1; i <= L; i++ {
			if index-i >= 0 && s[index-i] == charD {
				valid = false
				break
			}
		}

		if valid {
			for i := 1; i <= L; i++ {
				if index+i < len(s) && s[index+i] == charD {
					valid = false
					break
				}
			}
		}

		if valid {
			s[index] = charD

			if solve(s, L, index+1){
				return true
			}

			s[index] = '.'
		}

	}

	return false

}

func main() {
	var seq string
	var L int

	if _, err := fmt.Scan(&seq, &L); err != nil {
		return
	}

	s := []byte(seq)

	solve(s, L, 0)

	fmt.Println(string(s))

}
