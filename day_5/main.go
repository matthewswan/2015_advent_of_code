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
	inputSlice := strings.Split(string(input), "\n")
	niceStrCount := countNiceStrings(inputSlice)
	fmt.Printf("You have %d nice strings!\n", niceStrCount)
}

func countNiceStrings(input []string) int {
	niceStrings := make([]string, 0, len(input))
	for _, str := range input {
		switch {
		case strings.Contains(str, "ab"), strings.Contains(str, "cd"), strings.Contains(str, "pq"), strings.Contains(str, "xy"):
			continue
		case vowelCountMet(str) && duplicateLetterMet(str):
			niceStrings = append(niceStrings, str)
		}
	}
	return len(niceStrings)
}

func vowelCountMet(str string) bool {
	vowelCount := 0
	chars := strings.Split(str, "")
	for _, char := range chars {
		switch char {
		case "a", "e", "i", "o", "u":
			vowelCount++
		}
	}
	return vowelCount >= 3
}

func duplicateLetterMet(str string) bool {
	chars := strings.Split(str, "")
	for i := 0; i < len(chars)-1; i++ {
		if chars[i] == chars[i+1] {
			return true
		}
	}
	return false
}
