package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type point [2]int
func (p point) left() point {
	return point{p[0]-1, p[1]}
}
func (p point) right() point {
	return point{p[0]+1, p[1]}
}
func (p point) up() point {
	return point{p[0], p[1]-1}
}
func (p point) down() point {
	return point{p[0], p[1]+1}
}

func main() {
	dat, _ := ioutil.ReadFile("day11.txt")
	strCodes := strings.Split(string(dat), ",")
	codes := []int64{}
	for _, code := range strCodes {
		i, _ := strconv.ParseInt(code, 10, 64)
		codes = append(codes, i)
	}

	grid := make(map[point]int64)
	pos := point{0, 0}
	dir := 0

	// part 1
	grid[pos] = 0

	// part 2
	grid[pos] = 1

	index := 0

	input := func() int64 {
		p, ok := grid[pos]
		if ok {
			return p
		} else {
			return 0
		}
	}

	output := func(result int64) {
		if index % 2 == 0 {
			// paint
			grid[pos] = result
		} else {
			// move
			if result == 0 {
				dir = dir - 1
				if dir == -1 {
					dir = 3
				}
			} else {
				dir = dir + 1
				if dir == 4 {
					dir = 0
				}
			}
			if dir == 0 {
				pos = pos.up()
			} else if dir == 1 {
				pos = pos.right()
			} else if dir == 2 {
				pos = pos.down()
			} else if dir == 3 {
				pos = pos.left()
			}
		}

		index += 1
	}

	RunProgram(codes, input, output)

	fmt.Println(len(grid)) // part 1
	print(grid) // part 2
}

func print(grid map[point]int64) {
	maxX, maxY, minX, minY := 0, 0, 1000000, 1000000
	for p, _ := range grid {
		if p[0] < minX {
			minX = p[0]
		}
		if p[0] > maxX {
			maxX = p[0]
		}
		if p[1] < minY {
			minY = p[1]
		}
		if p[1] > maxY {
			maxY = p[1]
		}
	}

	for y := minY; y <= maxY; y += 1 {
		fmt.Println()
		for x := minX; x <= maxX; x += 1 {
			if grid[point{x, y}] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("X")
			}
		}
	}
}

