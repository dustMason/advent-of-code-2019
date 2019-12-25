package main

import (
	"fmt"
)

const (
	scaffold = '#'
	space    = '.'
	north    = '^'
	south    = 'v'
	west     = '<'
	east     = '>'
	comma    = ','
	L        = 'L'
	R        = 'R'
)

type vec struct {
	x, y int
	dir  rune
}

func (v vec) step() vec {
	switch v.dir {
	case north:
		return vec{v.x, v.y - 1, v.dir}
	case east:
		return vec{v.x + 1, v.y, v.dir}
	case south:
		return vec{v.x, v.y + 1, v.dir}
	case west:
		return vec{v.x - 1, v.y, v.dir}
	}
	return vec{}
}

func (v vec) turn(dir rune) vec {
	if dir == L {
		switch v.dir {
		case north:
			return vec{v.x, v.y, west}
		case east:
			return vec{v.x, v.y, north}
		case south:
			return vec{v.x, v.y, east}
		case west:
			return vec{v.x, v.y, south}
		}
	} else { // R
		switch v.dir {
		case north:
			return vec{v.x, v.y, east}
		case east:
			return vec{v.x, v.y, south}
		case south:
			return vec{v.x, v.y, west}
		case west:
			return vec{v.x, v.y, north}
		}
	}
	return vec{}
}

func main() {
	part1()
	buildPath()
	part2()
}

func part2() {
	codes := LoadIntcode("day17.txt")
	codes[0] = 2

	buf := []rune{}

	// hand-figured patterns:

	// A: R,6,R,6,R,8,L,10,L,4,
	//
	// B: R,6,L,10,R,8,
	// B: R,6,L,10,R,8,
	//
	// A: R,6,R,6,R,8,L,10,L,4,
	//
	// C: L,4,L,12,R,6,L,10,
	//
	// A: R,6,R,6,R,8,L,10,L,4,
	//
	// C: L,4,L,12,R,6,L,10,
	//
	// A: R,6,R,6,R,8,L,10,L,4,
	//
	// C: L,4,L,12,R,6,L,10,
	//
	// B: R,6,L,10,R,8

	inputString := "A,B,B,A,C,A,C,A,C,B\n"
	inputString += "R,6,R,6,R,8,L,10,L,4\n"
	inputString += "R,6,L,10,R,8\n"
	inputString += "L,4,L,12,R,6,L,10\n"
	inputString += "n\n"
	inputs := []rune(inputString)

	input := func() int64 { // noop
		i := inputs[0]
		if len(inputs) > 1 {
			inputs = inputs[1:]
		}
		return int64(i)
	}

	output := func(o int64) {
		fmt.Println(o)               // final print is the solution
		if o == 10 && len(buf) > 0 { // newline
			fmt.Println(string(buf))
			buf = []rune{}
		} else {
			buf = append(buf, rune(o))
		}
	}

	RunProgram(codes, input, output, 10000)
}

func buildPath() {
	codes := LoadIntcode("day17.txt")
	grid := [][]rune{}
	row := []rune{}

	input := func() int64 { // noop
		return 0
	}

	output := func(o int64) {
		if o == 10 && len(row) > 0 { // newline
			newRow := make([]rune, len(row))
			copy(newRow, row)
			grid = append(grid, newRow)
			row = []rune{}
		} else {
			row = append(row, rune(o))
		}
	}

	RunProgram(codes, input, output, 10000)

	look := func(v vec) (rune, bool) {
		if v.x < 0 || v.y < 0 || v.y >= len(grid) || v.x >= len(grid[v.y]) {
			return 'X', false
		}
		return grid[v.y][v.x], true
	}

	print := func(c []rune) {
		for _, r := range c {
			if r == L || r == R || r == comma {
				fmt.Print(string(r))
			} else {
				fmt.Print(r)
			}
		}
	}

	loc := vec{2, 0, north}

	// setup: dir = east
	loc = loc.turn(R)
	commands := []rune{R}
	for {
		// look forward until we hit space (inner loop)
		var steps int
		for {
			ahead, ok := look(loc.step())
			if ok && ahead != space {
				loc = loc.step() // keep moving
				steps++
			} else {
				commands = append(commands, comma, rune(steps))
				steps = 0
				break // we're at a turn
			}
		}
		// look for the one scaffold tile on either side
		leftHand, _ := look(loc.turn(L).step())
		rightHand, _ := look(loc.turn(R).step())
		if leftHand == scaffold && rightHand == scaffold {
			panic("found a turn with two options")
		}
		if leftHand != scaffold && rightHand != scaffold {
			print(commands)
			return
		}
		if leftHand == scaffold {
			commands = append(commands, comma, L)
			loc = loc.turn(L)
		} else {
			commands = append(commands, comma, R)
			loc = loc.turn(R)
		}
	}
}

func part1() {
	codes := LoadIntcode("day17.txt")
	grid := [][]rune{}
	row := []rune{}

	input := func() int64 { // noop
		return 0
	}

	output := func(o int64) {
		if o == 10 && len(row) > 0 { // newline
			newRow := make([]rune, len(row))
			copy(newRow, row)
			grid = append(grid, newRow)
			row = []rune{}
		} else {
			row = append(row, rune(o))
		}
	}

	RunProgram(codes, input, output, 10000)

	var sum int
	for y, row := range grid {
		for x, tile := range row {
			if x == 0 || x == len(row)-1 || y == 0 || y == len(grid)-1 {
				continue
			}
			if tile != scaffold {
				continue
			}
			l, r, u, d := row[x-1], row[x+1], grid[y-1][x], grid[y+1][x]
			if l == scaffold && r == scaffold && u == scaffold && d == scaffold {
				sum += x * y
			}
		}
	}
	fmt.Println(sum)
}
