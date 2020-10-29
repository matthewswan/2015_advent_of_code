package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("day-5-input.txt")
	if err != nil {
		panic(err)
	}
	inputSlice := strings.Fields(string(input))
	niceStrCount := countNiceStrings(inputSlice)
	fmt.Printf("You have %d nice strings!\n", niceStrCount)
}

func countNiceStrings(input []string) int {
	niceStrings := make([]string, 0, len(input))
	for _, str := range input {
		switch {
		case duplicatePairMet(str) && duplicateLetterWithGap(str):
			niceStrings = append(niceStrings, str)
		}
	}
	return len(niceStrings)
}

func duplicatePairMet(str string) bool {
	for i := 0; i < len(str)-2; i++ {
		if strings.Count(str, str[i:i+2]) >= 2 {
			return true
		}
	}
	return false
}

func duplicateLetterWithGap(str string) bool {
	for i := 0; i < len(str)-2; i++ {
		if str[i] == str[i+2] {
			return true
		}
	}
	return false
}
