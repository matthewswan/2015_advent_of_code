package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	StringCharCount = 0
	MemoryCharCount = 0
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		MemoryCharCount += len(line)
		encodedLine := strconv.Quote(line)
		StringCharCount += len(encodedLine)
	}
	fmt.Printf("Final count: %d\n", StringCharCount-MemoryCharCount)
}
