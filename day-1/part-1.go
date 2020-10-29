package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "()()))(())(())"
	fmt.Println(countFloors(s))
}

func countFloors(a string) int {
	var sum int
	for _, c := range a {
		if strings.Compare(string(c), "(") == 0 {
			sum += 1
		} else if strings.Compare(string(c), ")") == 0 {
			sum -= 1
		}
	}
	return sum
}

// func countFloors(a string) int {
// 	var sum int
// 	runes := []rune(a)
// 	for _, c := range runes {
// 		if c == 40 {
// 			sum += 1
// 		} else if c == 41 {
// 			sum -= 1
// 		}
// 	}
// 	return sum
// }
