package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s := `20x3x11
	4x8x78
	39x2x89`

	toys := strings.Split(s, "\n")
	var ribbonLength int
	for _, toy := range toys {
		var intSizes []int
		sizes := strings.Split(toy, "x")
		for _, size := range sizes {
			intSize, _ := strconv.Atoi(size)
			intSizes = append(intSizes, intSize)
		}
		sort.Ints(intSizes)
		ribbonLength += 2*(intSizes[0]+intSizes[1]) + (intSizes[0] * intSizes[1] * intSizes[2])
	}
	fmt.Println(ribbonLength)
}
