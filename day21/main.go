package main

import "fmt"

func main() {
	buf := []rune{}
	output := func(o int64) {
		if o == 10 && len(buf) > 0 { // newline
			fmt.Println(string(buf))
			buf = []rune{}
		} else if o > 255 {
			fmt.Println(o)
		} else {
			buf = append(buf, rune(o))
		}
	}

	codes := LoadIntcode("day21.txt")
	RunProgram(codes, part1(), output, 5000)
	RunProgram(codes, part2(), output, 5000)
}

func part1() func() int64 {
	inputString := `NOT A T
NOT C J
OR T J
AND D J
WALK
`
	inputs := []rune(inputString)
	return func() int64 {
		i := inputs[0]
		if len(inputs) > 1 {
			inputs = inputs[1:]
		}
		return int64(i)
	}
}

func part2() func() int64 {
	inputString := `NOT H T
OR C T
AND B T
AND A T
NOT T J
AND D J
RUN
`
	inputs := []rune(inputString)
	return func() int64 {
		i := inputs[0]
		if len(inputs) > 1 {
			inputs = inputs[1:]
		}
		return int64(i)
	}
}
