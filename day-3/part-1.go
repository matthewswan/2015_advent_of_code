package main

import (
	"io/ioutil"
	"strings"
)

type Coor struct {
	x, y int
}

func main() {
	input, err := ioutil.ReadFile("./day-3-input.txt")
	if err != nil {
		panic(err)
	}

	homes := make(map[Coor]bool)
	coor := Coor{0, 0}

	homes[coor] = true

	for _, dir := range strings.Split(string(input), "") {
		switch dir {
		case ">":
			coor = Coor{coor.x + 1, coor.y}
		case "<":
			coor = Coor{coor.x - 1, coor.y}
		case "^":
			coor = Coor{coor.x, coor.y + 1}
		case "v":
			coor = Coor{coor.x, coor.y - 1}
		}

		homes[coor] = true
	}

	println(len(homes))
}
