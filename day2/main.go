package main

import (
	"fmt"
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
	part1 := runProgram(12, 2, codes)
	fmt.Println(part1)

	// part 2: search combinations of 1 and 2 to get magic number:
	answer := 19690720
	for a := 0; a < 100; a += 1 {
		for b := 0; b < 100; b += 1 {
			ans := runProgram(a, b, codes)
			if ans == answer {
			  fmt.Println(100 * a + b)
				panic("found")
			}
		}
	}
}

func runProgram(a int, b int, program []int) int {
  codes := append([]int{}, program...)
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
