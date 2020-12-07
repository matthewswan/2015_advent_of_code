package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day-6-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var (
		scanner       = bufio.NewScanner(file)
		grid          [1000][1000]bool
		lightsOnCount int
		turnOnRegex   = regexp.MustCompile("turn on")
		turnOffRegex  = regexp.MustCompile("turn off")
		toggleRegex   = regexp.MustCompile("toggle")
	)

	for scanner.Scan() {
		instruction := scanner.Text()
		startInstruction, finishInstruction := findInstructions(instruction)
		start, finish := convertInstruction(startInstruction), convertInstruction(finishInstruction)
		for x := start[0]; x <= finish[0]; x++ {
			for y := start[1]; y <= finish[1]; y++ {
				if match := turnOnRegex.FindString(instruction); match != "" {
					grid[x][y] = true
				}
				if match := turnOffRegex.FindString(instruction); match != "" {
					grid[x][y] = false
				}
				if match := toggleRegex.FindString(instruction); match != "" {
					grid[x][y] = !grid[x][y]
				}
			}
		}
	}
	for i, _ := range grid[0] {
		for j, _ := range grid[1] {
			if grid[i][j] {
				lightsOnCount++
			}
		}
	}
	fmt.Printf("The number of lights on: %d\n", lightsOnCount)
}

func convertInstruction(str []string) []int {
	strSlice := make([]int, 0)
	for _, s := range str {
		i, _ := strconv.Atoi(s)
		strSlice = append(strSlice, i)
	}
	return strSlice
}

func findInstructions(instr string) ([]string, []string) {
	fromRegex := regexp.MustCompile("\\d+,\\d+")
	toRegex := regexp.MustCompile("\\d+,\\d+$")
	return strings.Split(fromRegex.FindString(instr), ","), strings.Split(toRegex.FindString(instr), ",")
}
