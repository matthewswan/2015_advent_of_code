package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findSmallestIntFor00000PrefixMD5("iwrupvqb"))
}

func findSmallestIntFor00000PrefixMD5(str string) int {
	for i := 1; i > 0; i++ {
		s := strconv.Itoa(i)
		input := []byte(str + s)
		mdSum := md5.Sum(input)
		if fmt.Sprintf("%x", mdSum)[:5] == "00000" {
			return i
		}
	}
	return 0
}
