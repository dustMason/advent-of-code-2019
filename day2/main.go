package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("day2.txt")
	strCodes := strings.Split(string(dat), ",")
	codes := []int{}
	for _, code := range strCodes {
		i, _ := strconv.Atoi(code)
		codes = append(codes, i)
	}

	// part 1: modify the program:
	// codes[1] = 12
	// codes[2] = 2

	// part 2: search combinations of 1 and 2 to get magic number:
	answer := 19690720

	for a := 0; a < 100; a += 1 {
		for b := 0; b < 100; b += 1 {
			newCodes := append([]int{}, codes...)
			ans := runProgram(a, b, newCodes)
			if ans == answer {
				panic(100 * a + b)
			}
		}
	}
}

func runProgram(a int, b int, codes []int) int {
	pointer := 0
	codes[1] = a
	codes[2] = b
	for {
		opcode := codes[pointer]
		if opcode == 1 {
			a := codes[pointer+1]
			b := codes[pointer+2]
			c := codes[pointer+3]
			codes[c] = codes[a] + codes[b]
		} else if opcode == 2 {
			a := codes[pointer+1]
			b := codes[pointer+2]
			c := codes[pointer+3]
			codes[c] = codes[a] * codes[b]
		} else { // includes opcode 99
			break
		}
		pointer += 4
	}
	return codes[0]
}
