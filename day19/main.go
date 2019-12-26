package main

import "fmt"

func main() {
	part1()
	part2()
}

func part1() {
	codes := LoadIntcode("day19.txt")
	count := 0

	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			inputs := []int{x, y}
			input := func() int64 {
				i := inputs[len(inputs)-1]
				inputs = inputs[:len(inputs)-1]
				return int64(i)
			}

			output := func(o int64) {
				if o == 1 {
					count++
				}
			}

			RunProgram(codes, input, output, 500)
		}
	}
	fmt.Println(count)
}

func part2() {
	codes := LoadIntcode("day19.txt")

	look := func(x, y int) bool {
		result := false
		inputs := []int{x, y}
		index := 0
		input := func() int64 {
			i := inputs[index]
			index++
			return int64(i)
		}
		output := func(o int64) {
			if o == 1 {
				result = true
			}
		}
		RunProgram(codes, input, output, 500)
		return result
	}

	x, y := 0, 0
	for !look(x+99, y) {
		y++
		for !look(x, y+99) {
			x++
		}
	}
	fmt.Println(x*10000 + y)
}
