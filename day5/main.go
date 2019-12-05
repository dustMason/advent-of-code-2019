package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("day5.txt")
	strCodes := strings.Split(string(dat), ",")
	codes := []int{}
	for _, code := range strCodes {
		i, _ := strconv.Atoi(code)
		codes = append(codes, i)
	}

	runProgram(1, codes)
	runProgram(5, codes)
}

func runProgram(input int, program []int) {
	fmt.Printf("running with input %v\n", input)
  codes := append([]int{}, program...)
	pointer := 0

	for {
		rawCode := codes[pointer]
		digits := intToSlice(rawCode, []int{})
		reverse(digits)

		opcode := digits[0]
		mode1 := getMode(digits, 2)
		mode2 := getMode(digits, 3)

		if opcode == 9 { // we're cheating, this should be 99
			break
		} else if opcode == 1 {
			a := getVal(codes, codes[pointer+1], mode1)
			b := getVal(codes, codes[pointer+2], mode2)
			c := codes[pointer+3]
			codes[c] = a + b
			pointer += 4
		} else if opcode == 2 {
			a := getVal(codes, codes[pointer+1], mode1)
			b := getVal(codes, codes[pointer+2], mode2)
			c := codes[pointer+3]
			codes[c] = a * b
			pointer += 4
		} else if opcode == 3 { // STDIN
			a := codes[pointer+1]
			codes[a] = input
			pointer += 2
		} else if opcode == 4 { // STDOUT
			a := codes[pointer+1]
			fmt.Printf("%v ", codes[a])
			pointer += 2
		} else if opcode == 5 { // jump-if-true
			a := getVal(codes, codes[pointer+1], mode1)
			if a != 0 {
				pointer = getVal(codes, codes[pointer+2], mode2)
			} else {
				pointer += 3
			}
		} else if opcode == 6 { // jump-if-false
			a := getVal(codes, codes[pointer+1], mode1)
			if a == 0 {
				pointer = getVal(codes, codes[pointer+2], mode2)
			} else {
				pointer += 3
			}
		} else if opcode == 7 { // less then
			a := getVal(codes, codes[pointer+1], mode1)
			b := getVal(codes, codes[pointer+2], mode2)
			c := codes[pointer+3]
			if a < b {
				codes[c] = 1
			} else {
				codes[c] = 0
			}
			pointer += 4
		} else if opcode == 8 { // equal
			a := getVal(codes, codes[pointer+1], mode1)
			b := getVal(codes, codes[pointer+2], mode2)
			c := codes[pointer+3]
			if a == b {
				codes[c] = 1
			} else {
				codes[c] = 0
			}
			pointer += 4
		} else {
			break
		}
	}
	fmt.Println()
}

func getMode(d []int, i int) int {
	if i > len(d) - 1 {
		return 0
	}
	return d[i]
}

func getVal(codes []int, address int, mode int) int {
	if mode == 1 {
		return address
	} else {
		return codes[address]
	}
}

func reverse(digits []int) {
	for i := len(digits)/2 - 1; i >= 0; i-- {
		opp := len(digits) - 1 - i
		digits[i], digits[opp] = digits[opp], digits[i]
	}
}

func intToSlice(n int, sequence []int) []int {
	if n != 0 {
		i := n % 10
		sequence = append([]int{i}, sequence...)
		return intToSlice(n/10, sequence)
	}
	return sequence
}
