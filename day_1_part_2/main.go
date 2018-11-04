package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := `)`
	s2 := `()())`
	fmt.Println(findBasement(s1))
	fmt.Println(findBasement(s2))
}

func findBasement(s string) int {
	var floor int
	for i := 0; i < len(s); i++ {
		if strings.Compare(string(s[i]), "(") == 0 {
			floor += 1
		} else if strings.Compare(string(s[i]), ")") == 0 {
			floor -= 1
		}

		if floor < 0 {
			return i + 1
		}
	}
	return 0
}
