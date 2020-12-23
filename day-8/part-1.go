package main

import (
	"bufio"
	"fmt"
	"log"
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
		StringCharCount += len(line)
		memString, err := strconv.Unquote(line)
		if err != nil {
			log.Fatal(err)
		}
		MemoryCharCount += len(fmt.Sprintf("%s", memString))
	}
	fmt.Printf("Final count: %d\n", StringCharCount-MemoryCharCount)
}
