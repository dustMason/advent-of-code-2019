package main

import (
	"fmt"
	"strings"
)

const input = `#.##.
###.#
#...#
##..#
.#...`

const (
	bug   = '#'
	space = '.'
)

var deltas = []vec{
	{0, 1},
	{-1, 0},
	{0, -1},
	{1, 0},
}

type vec struct {
	x, y int
}

func (v vec) add(o vec) vec {
	return vec{v.x + o.x, v.y + o.y}
}

func (v vec) adjacent(w int, h int) []vec {
	out := []vec{}
	for _, delta := range deltas {
		o := v.add(delta)
		if o.x >= 0 && o.y >= 0 && o.x < w && o.y < h {
			out = append(out, o)
		}
	}
	return out
}

func countBugs(vecs []vec, grid map[vec]rune) int {
	count := 0
	for _, v := range vecs {
		if grid[v] == bug {
			count++
		}
	}
	return count
}

func part1() {
	grid := make(map[vec]rune)
	lines := strings.Split(input, "\n")
	var w, h int
	for y, line := range lines {
		h = y + 1
		for x, b := range line {
			w = x + 1
			grid[vec{x, y}] = b
		}
	}

	history := make(map[string]bool)

	for {
		// printGrid(grid, w, h)
		newGrid := make(map[vec]rune)
		for v, r := range grid {
			adj := v.adjacent(w, h)
			bugs := countBugs(adj, grid)
			if r == bug && bugs != 1 {
				newGrid[v] = space
			} else if r == space && (bugs == 1 || bugs == 2) {
				newGrid[v] = bug
			} else {
				newGrid[v] = r
			}
		}
		grid = newGrid

		gridStr := fmt.Sprint(grid)
		if _, ok := history[gridStr]; ok {
			// printGrid(grid, w, h)
			fmt.Println(biodiversity(grid, w, h))
			break
		}
		history[fmt.Sprint(grid)] = true
	}
}

func printGrid(grid map[vec]rune, w int, h int) {
	for y := 0; y < h; y++ {
		fmt.Println()
		for x := 0; x < w; x++ {
			fmt.Print(string(grid[vec{x, y}]))
		}
	}
	fmt.Println()
}

func biodiversity(grid map[vec]rune, w int, h int) int64 {
	var total int64
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if grid[vec{x, y}] == bug {
				total += pow(2, int64((w*y)+x))
			}
		}
	}
	return total
}

func pow(a, b int64) int64 {
	var result int64 = 1
	for 0 != b {
		if 0 != (b & 1) {
			result *= a
		}
		b >>= 1
		a *= a
	}
	return result
}
