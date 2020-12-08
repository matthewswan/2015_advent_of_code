package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var circuitMap = make(map[string]uint16)

func parse(str string) (uint16, error) {
	number, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		if _, exists := circuitMap[str]; !exists {
			return uint16(number), err
		}
		return circuitMap[str], nil
	}
	return uint16(number), err
}

func main() {
	for {
		f, _ := os.Open("input.txt")
		defer f.Close()
		s := bufio.NewScanner(f)
		for s.Scan() {
			line := s.Text()
			re1 := regexp.MustCompile("^(NOT )?([a-z0-9]+) -> ([a-z]+)$")
			parts1 := re1.FindStringSubmatch(line)
			if len(parts1) > 0 {
				if _, exists := circuitMap[parts1[3]]; exists {
					continue
				}
				value, err := parse(parts1[2])
				if err != nil {
					continue
				}
				if parts1[1] == "NOT " {
					value = ^value
				}
				circuitMap[parts1[3]] = value
			}
			re2 := regexp.MustCompile("^([a-z0-9]+) (AND|OR|LSHIFT|RSHIFT) ([a-z0-9]+) -> ([a-z]+)$")
			parts2 := re2.FindStringSubmatch(line)
			if len(parts2) > 0 {
				if _, exists := circuitMap[parts2[4]]; exists {
					continue
				}
				value1, err := parse(parts2[1])
				if err != nil {
					continue
				}
				value2, err := parse(parts2[3])
				if err != nil {
					continue
				}
				switch parts2[2] {
				case "AND":
					circuitMap[parts2[4]] = value1 & value2
				case "OR":
					circuitMap[parts2[4]] = value1 | value2
				case "LSHIFT":
					circuitMap[parts2[4]] = value1 << value2
				case "RSHIFT":
					circuitMap[parts2[4]] = value1 >> value2

				}
			}
		}
		if val, done := circuitMap["a"]; done {
			fmt.Println(val)
			break
		}
	}
}
