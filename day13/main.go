package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type vec2 struct {
	x int64
	y int64
}

func main() {
	part1()
	part2()
}

func part1() {
	codes := load()
	pixels := make(map[vec2]int64)
	buf := []int64{}
	input := func() int64 { return 0 }
	output := func(o int64) {
		buf = append(buf, o)
		if len(buf) == 3 {
			pixels[vec2{buf[0], buf[1]}] = buf[2]
			buf = []int64{}
		}
	}

	RunProgram(codes, input, output)
	blocks := 0
	for _, t := range pixels {
		if t == 2 {
			blocks += 1
		}
	}
	fmt.Println(blocks)
}

func part2() {
	codes := load()
	pixels := make(map[vec2]int64)
	codes[0] = 2 // insert quarters!
	var score, ballX, paddleX int64

	buf := []int64{}
	input := func() int64 {
		if ballX > paddleX {
			return 1
		}
		if ballX < paddleX {
			return -1
		}
		return 0
	}
	output := func(o int64) {
		buf = append(buf, o)
		if len(buf) == 3 {
			x, y, t := buf[0], buf[1], buf[2]
			if x == -1 && y == y {
				score = t
			} else {
				pixels[vec2{x, y}] = t
				if t == 3 {
					paddleX = x
				}
				if t == 4 {
					ballX = x
				}
			}
			buf = []int64{}
		}
	}

	RunProgram(codes, input, output)
	fmt.Println(score)
}

func load() []int64 {
	dat, _ := ioutil.ReadFile("day13.txt")
	strCodes := strings.Split(string(dat), ",")
	codes := []int64{}
	for _, code := range strCodes {
		i, _ := strconv.ParseInt(code, 10, 64)
		codes = append(codes, i)
	}
	return codes
}
