package main

import (
	"io/ioutil"
	"strings"
)

type Coor struct {
	x, y int
}

type Santa struct {
	coor Coor
}

func (s *Santa) gift(homes map[Coor]bool) {
	homes[s.coor] = true
}
func (s *Santa) up() {
	coor := s.coor
	s.coor = Coor{coor.x, coor.y + 1}
}

func (s *Santa) down() {
	coor := s.coor
	s.coor = Coor{coor.x, coor.y - 1}
}

func (s *Santa) left() {
	coor := s.coor
	s.coor = Coor{coor.x - 1, coor.y}
}

func (s *Santa) right() {
	coor := s.coor
	s.coor = Coor{coor.x + 1, coor.y}
}

func main() {
	input, err := ioutil.ReadFile("day-3-input.txt")
	if err != nil {
		panic(err)
	}

	directions := strings.Split(string(input), "")
	santas := []*Santa{new(Santa), new(Santa)}

	homes := make(map[Coor]bool)
	homes[Coor{0, 0}] = true

	for i, dir := range directions {
		santa := santas[i%2]

		switch dir {
		case ">":
			santa.right()
		case "<":
			santa.left()
		case "^":
			santa.up()
		case "v":
			santa.down()
		}

		santa.gift(homes)
	}

	println(len(homes))
}
