package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func RunProgram(program []int64, input func() int64, output func(int64), memSize int) {
	codes := make([]int64, memSize)
	copy(codes, program)
	pointer := int64(0)
	relativeBase := int64(0)

	for {
		cmd := codes[pointer]
		opcode := cmd % 100

		mode := func(offset int64) int64 {
			return cmd / pow(10, offset+1) % 10
		}

		read := func(offset int64) int64 {
			p := codes[pointer+offset]
			m := mode(offset)
			if m == 0 {
				return codes[p]
			} else if m == 2 {
				return codes[p+relativeBase]
			} else {
				return p
			}
		}

		write := func(offset int64, value int64) {
			m := mode(offset)
			v := codes[pointer+offset]
			if m == 2 {
				codes[relativeBase+v] = value
			} else {
				codes[v] = value
			}
		}

		if opcode == 99 {
			break
		} else if opcode == 1 {
			a, b := read(1), read(2)
			write(3, a+b)
			pointer += 4
		} else if opcode == 2 {
			a, b := read(1), read(2)
			write(3, a*b)
			pointer += 4
		} else if opcode == 3 { // STDIN
			v := input()
			write(1, v)
			pointer += 2
		} else if opcode == 4 { // STDOUT
			a := read(1)
			output(a)
			pointer += 2
		} else if opcode == 5 { // jump-if-true
			a, b := read(1), read(2)
			if a != 0 {
				pointer = b
			} else {
				pointer += 3
			}
		} else if opcode == 6 { // jump-if-false
			a, b := read(1), read(2)
			if a == 0 {
				pointer = b
			} else {
				pointer += 3
			}
		} else if opcode == 7 { // less than
			a, b := read(1), read(2)
			if a < b {
				write(3, 1)
			} else {
				write(3, 0)
			}
			pointer += 4
		} else if opcode == 8 { // equal
			a, b := read(1), read(2)
			if a == b {
				write(3, 1)
			} else {
				write(3, 0)
			}
			pointer += 4
		} else if opcode == 9 {
			relativeBase += read(1)
			pointer += 2
		} else {
			panic("bad instruction")
		}
	}
}

func LoadIntcode(filename string) []int64 {
	dat, _ := ioutil.ReadFile(filename)
	strCodes := strings.Split(string(dat), ",")
	codes := []int64{}
	for _, code := range strCodes {
		i, _ := strconv.ParseInt(code, 10, 64)
		codes = append(codes, i)
	}
	return codes
}

// Source: https://groups.google.com/d/msg/golang-nuts/PnLnr4bc9Wo/z9ZGv2DYxXoJ
func pow(a, b int64) int64 {
	p := int64(1)
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}
