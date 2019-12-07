package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("day7.txt")
	strCodes := strings.Split(string(dat), ",")
	codes := []int{}
	for _, code := range strCodes {
		i, _ := strconv.Atoi(code)
		codes = append(codes, i)
	}

	// part 1
	best := 0
	for _, values := range permutations([]int{0, 1, 2, 3, 4}) {
		result := cycle(values, codes)
		if result > best {
			best = result
		}
	}
	fmt.Println(best)

	// part 2
	best = 0
	for _, values := range permutations([]int{5, 6, 7, 8, 9}) {
		result := cycle(values, codes)
		if result > best {
			best = result
		}
	}
	fmt.Println(best)
}

func cycle(phases []int, program []int) int {
	done := make(chan bool)

	ea := make(chan int, 1)
	ab := make(chan int)
	bc := make(chan int)
	cd := make(chan int)
	de := make(chan int)

	go runProgram(program, ea, ab, done)
	go runProgram(program, ab, bc, done)
	go runProgram(program, bc, cd, done)
	go runProgram(program, cd, de, done)
	go runProgram(program, de, ea, done)

	ea <- phases[0]
	ab <- phases[1]
	bc <- phases[2]
	cd <- phases[3]
	de <- phases[4]

	ea <- 0 // first input
	for i := 0; i < 5; i++ {
		<-done
	}

	return <-ea
}

func runProgram(program []int, input <-chan int, output chan<- int, done chan<- bool) {
	codes := append([]int{}, program...)
	pointer := 0

	for {
		cmd := codes[pointer]
		opcode := cmd % 100

		arg := func(offset int) int {
			parameter := codes[pointer+offset]
			mode := cmd / pow(10, offset+1) % 10
			if mode == 0 {
				return codes[parameter]
			} else {
				return parameter
			}
		}

		if opcode == 99 {
			done <- true
			return
		} else if opcode == 1 {
			a, b, c := arg(1), arg(2), codes[pointer+3]
			codes[c] = a + b
			pointer += 4
		} else if opcode == 2 {
			a, b, c := arg(1), arg(2), codes[pointer+3]
			codes[c] = a * b
			pointer += 4
		} else if opcode == 3 { // STDIN
			a := codes[pointer+1]
			codes[a] = <-input
			pointer += 2
		} else if opcode == 4 { // STDOUT
			a := arg(1)
			output <- a
			pointer += 2
		} else if opcode == 5 { // jump-if-true
			a, b := arg(1), arg(2)
			if a != 0 {
				pointer = b
			} else {
				pointer += 3
			}
		} else if opcode == 6 { // jump-if-false
			a, b := arg(1), arg(2)
			if a == 0 {
				pointer = b
			} else {
				pointer += 3
			}
		} else if opcode == 7 { // less than
			a, b, c := arg(1), arg(2), codes[pointer+3]
			if a < b {
				codes[c] = 1
			} else {
				codes[c] = 0
			}
			pointer += 4
		} else if opcode == 8 { // equal
			a, b, c := arg(1), arg(2), codes[pointer+3]
			if a == b {
				codes[c] = 1
			} else {
				codes[c] = 0
			}
			pointer += 4
		} else {
			panic("bad instruction")
		}
	}
}

func permutations(values []int) (result [][]int) {
	if len(values) == 1 {
		result = append(result, values)
		return
	}
	for i, current := range values {
		others := make([]int, 0, len(values)-1)
		others = append(others, values[:i]...)
		others = append(others, values[i+1:]...)
		for _, step := range permutations(others) {
			result = append(result, append(step, current))
		}
	}
	return
}

// Source: https://groups.google.com/d/msg/golang-nuts/PnLnr4bc9Wo/z9ZGv2DYxXoJ
func pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}
