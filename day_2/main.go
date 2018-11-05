package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s := `20x3x11
15x27x5
14x12x8`
	toys := strings.Split(s, "\n")
	sum := 0
	for _, toy := range toys {
		var intSizes []int
		sizes := strings.Split(toy, "x")
		for _, size := range sizes {
			intSize, _ := strconv.Atoi(size)
			intSizes = append(intSizes, intSize)
		}
		sort.Ints(intSizes)
		sum += 2*intSizes[0]*intSizes[1] + 2*intSizes[0]*intSizes[2] + 2*intSizes[1]*intSizes[2] + intSizes[0]*intSizes[1]
	}
	fmt.Println(sum)
}
