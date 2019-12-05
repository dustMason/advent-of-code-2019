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

		if opcode == 9 { // we're cheating, this should be 99
			break
		} else if opcode == 1 {
			args := getArgs(codes, pointer, digits, 2)
			c := codes[pointer+3]
			codes[c] = args[0] + args[1]
			pointer += 4
		} else if opcode == 2 {
			args := getArgs(codes, pointer, digits, 2)
			c := codes[pointer+3]
			codes[c] = args[0] * args[1]
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
			args := getArgs(codes, pointer, digits, 2)
			if args[0] != 0 {
				pointer = args[1]
			} else {
				pointer += 3
			}
		} else if opcode == 6 { // jump-if-false
			args := getArgs(codes, pointer, digits, 2)
			if args[0] == 0 {
				pointer = args[1]
			} else {
				pointer += 3
			}
		} else if opcode == 7 { // less then
			args := getArgs(codes, pointer, digits, 2)
			c := codes[pointer+3]
			if args[0] < args[1] {
				codes[c] = 1
			} else {
				codes[c] = 0
			}
			pointer += 4
		} else if opcode == 8 { // equal
			args := getArgs(codes, pointer, digits, 2)
			c := codes[pointer+3]
			if args[0] == args[1] {
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

func getArgs(codes []int, pointer int, opcode []int, n int) []int {
	result := []int{}
	for i := 0; i < n; i++ {
		mode := getMode(opcode, i+2)
		val := getVal(codes, codes[pointer+i+1], mode)
		result = append(result, val)
	}
	return result
}

func getMode(d []int, i int) int {
	if i > len(d)-1 {
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
