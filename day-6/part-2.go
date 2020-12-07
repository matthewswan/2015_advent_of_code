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
		scanner         = bufio.NewScanner(file)
		grid            [1000][1000]int
		lightBrightness int
		turnOnRegex     = regexp.MustCompile("turn on")
		turnOffRegex    = regexp.MustCompile("turn off")
		toggleRegex     = regexp.MustCompile("toggle")
	)

	for scanner.Scan() {
		instruction := scanner.Text()
		startInstruction, finishInstruction := findInstructions(instruction)
		start, finish := convertInstruction(startInstruction), convertInstruction(finishInstruction)
		for x := start[0]; x <= finish[0]; x++ {
			for y := start[1]; y <= finish[1]; y++ {
				if match := turnOnRegex.FindString(instruction); match != "" {
					grid[x][y]++
					lightBrightness++
				}
				if match := turnOffRegex.FindString(instruction); match != "" {
					if grid[x][y] == 0 {
						continue
					}
					grid[x][y]--
					lightBrightness--
				}
				if match := toggleRegex.FindString(instruction); match != "" {
					grid[x][y] += 2
					lightBrightness += 2
				}
			}
		}
	}
	fmt.Printf("The number of lights on: %d\n", lightBrightness)
}

func convertInstruction(strSlice []string) (intSlice []int) {
	for _, str := range strSlice {
		i, _ := strconv.Atoi(str)
		intSlice = append(intSlice, i)
	}
	return intSlice
}

func findInstructions(instr string) ([]string, []string) {
	fromRegex, toRegex := regexp.MustCompile("\\d+,\\d+"), regexp.MustCompile("\\d+,\\d+$")
	return strings.Split(fromRegex.FindString(instr), ","), strings.Split(toRegex.FindString(instr), ",")
}
