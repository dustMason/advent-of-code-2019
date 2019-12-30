package main

import (
	"fmt"
	"strings"
)

type vec3 struct {
	x, y, z int
}

var deltas3 = []vec3{
	{0, 1, 0},
	{-1, 0, 0},
	{0, -1, 0},
	{1, 0, 0},
}

func (v vec3) add(o vec3) vec3 {
	return vec3{v.x + o.x, v.y + o.y, v.z + o.z}
}

func (v vec3) adjacent(w int, h int) []vec3 {
	out := []vec3{}
	for _, delta := range deltas3 {
		o := v.add(delta)

		// if o is a center tile, adjacent includes a whole set of tiles at z-1:
		if o.x == 2 && o.y == 2 {
			if delta.x == -1 {
				out = append(out,
					vec3{4, 0, v.z - 1},
					vec3{4, 1, v.z - 1},
					vec3{4, 2, v.z - 1},
					vec3{4, 3, v.z - 1},
					vec3{4, 4, v.z - 1},
				)
			} else if delta.x == 1 {
				out = append(out,
					vec3{0, 0, v.z - 1},
					vec3{0, 1, v.z - 1},
					vec3{0, 2, v.z - 1},
					vec3{0, 3, v.z - 1},
					vec3{0, 4, v.z - 1},
				)
			} else if delta.y == -1 {
				out = append(out,
					vec3{0, 4, v.z - 1},
					vec3{1, 4, v.z - 1},
					vec3{2, 4, v.z - 1},
					vec3{3, 4, v.z - 1},
					vec3{4, 4, v.z - 1},
				)
			} else if delta.y == 1 {
				out = append(out,
					vec3{0, 0, v.z - 1},
					vec3{1, 0, v.z - 1},
					vec3{2, 0, v.z - 1},
					vec3{3, 0, v.z - 1},
					vec3{4, 0, v.z - 1},
				)
			}
			// if o is out of bounds, adjacent includes tiles at z+1:
		} else if o.x < 0 {
			out = append(out, vec3{1, 2, v.z + 1})
		} else if o.x == w {
			out = append(out, vec3{3, 2, v.z + 1})
		} else if o.y < 0 {
			out = append(out, vec3{2, 1, v.z + 1})
		} else if o.y == h {
			out = append(out, vec3{2, 3, v.z + 1})
		} else {
			out = append(out, o)
		}
	}
	return out
}

func part2() {
	grid := make(map[vec3]rune)
	lines := strings.Split(input, "\n")
	var w, h int
	for y, line := range lines {
		h = y + 1
		for x, b := range line {
			w = x + 1
			grid[vec3{x, y, 0}] = b
		}
	}

	for i := 0; i < 200; i++ {
		newGrid := make(map[vec3]rune)

		// make sure the grid has an extra blank level on each end of the z
		grid = fill(grid, w, h)

		for v, r := range grid {
			if v.x == 2 && v.y == 2 {
				continue
			}
			bugs := countBugs3(v, w, h, grid)
			if r == bug && bugs != 1 {
				newGrid[v] = space
			} else if r == space && (bugs == 1 || bugs == 2) {
				newGrid[v] = bug
			} else {
				newGrid[v] = r
			}
		}
		grid = newGrid
	}

	total := 0
	for _, r := range grid {
		if r == bug {
			total++
		}
	}
	fmt.Println(total)
}

func countBugs3(v vec3, w int, h int, grid map[vec3]rune) int {
	vecs := v.adjacent(w, h)
	count := 0
	for _, v := range vecs {
		if grid[v] == bug {
			count++
		}
	}
	return count
}

func fill(grid map[vec3]rune, w int, h int) map[vec3]rune {
	minZ, maxZ, counts := 500, 0, make(map[int]int)
	for v, r := range grid {
		if v.z < minZ {
			minZ = v.z
		}
		if v.z > maxZ {
			maxZ = v.z
		}
		if r == bug {
			counts[v.z]++
		}
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if counts[minZ] > 0 {
				grid[vec3{x, y, minZ - 1}] = space
			}
			if counts[maxZ] > 0 {
				grid[vec3{x, y, maxZ + 1}] = space
			}
		}
	}

	return grid
}
