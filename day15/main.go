package main

import "fmt"

type vec struct {
	x, y int
}

func (v vec) add(o vec) vec {
	return vec{v.x + o.x, v.y + o.y}
}

const (
	wall  = 1
	open  = 2
	found = 3

	north = 1
	south = 2
	west  = 3
	east  = 4
)

var deltas = map[int]vec{
	south: vec{0, 1},
	west:  vec{-1, 0},
	north: vec{0, -1},
	east:  vec{1, 0},
}

func main() {
	codes := LoadIntcode("day15.txt")

	left := map[int]int{
		north: west,
		south: east,
		west:  south,
		east:  north,
	}

	right := map[int]int{
		north: east,
		south: west,
		west:  north,
		east:  south,
	}

	area := make(map[vec]int)
	loc := vec{0, 0}
	var target vec
	facing := south

	steps := 0

	input := func() int64 {
		result := int64(facing)
		return result
	}

	output := func(o int64) {
		c := int(o + 1)
		where := loc.add(deltas[facing])

		if where.x == 0 && where.y == 0 {
			// we looped all the way back to the start
			fmt.Printf("part 1: %v\n", steps)
			time := part2(area, target)
			fmt.Printf("part 2: %v\n", time)
			panic("done")
		}

		// simple exploration for graphs with no cycles: hit a wall? turn right. otherwise go fwd, turn left.
		if c == wall {
			facing = right[facing]
		} else {
			facing = left[facing]
			loc = where
			if target.x == 0 && target.y == 0 { // if oxygen hasn't been found
				_, k := area[where]
				if k {
					steps -= 1
				} else {
					steps += 1
				}
			}
			if c == found {
				target = where
			}
		}
		area[where] = c
	}

	RunProgram(codes, input, output)
}

func part2(area map[vec]int, oxygen vec) int {
	// do a BFS search starting from oxygen location. each round of possible single moves is an iteration of the outer
	// loop. we stop when the count of filled tiles equals the count of open tiles.

	var minutes int
	var okCount int
	for _, i := range area {
		if i == open || i == found {
			okCount++
		}
	}
	neighbors := []vec{oxygen}
	filled := map[vec]bool{oxygen: true}

	for {
		minutes += 1
		newNeighbors := map[vec]bool{}
		for {
			if len(neighbors) == 0 {
				break
			}
			loc := neighbors[len(neighbors)-1]
			neighbors = neighbors[:len(neighbors)-1]
			for _, delta := range deltas {
				neighbor := loc.add(delta)
				_, isFilled := filled[neighbor]
				_, isNewNeighbor := newNeighbors[neighbor]
				if area[neighbor] == open && !isFilled && !isNewNeighbor {
					filled[neighbor] = true
					newNeighbors[neighbor] = true
				}
			}
		}

		neighbors = []vec{}
		for v, _ := range newNeighbors {
			neighbors = append(neighbors, v)
		}

		if okCount == len(filled) {
			break
		}
	}
	return minutes
}

// print the maze for debugging
func printArea(grid map[vec]int, loc vec) {
	maxX, maxY, minX, minY := 0, 0, 1000000, 1000000
	for p, _ := range grid {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	for y := minY - 1; y <= maxY; y += 1 {
		fmt.Println()
		for x := minX - 1; x <= maxX; x += 1 {
			if x == 0 && y == 0 {
				fmt.Print("s")
			} else if loc.x == x && loc.y == y {
				fmt.Print("*")
			} else if grid[vec{x, y}] == wall {
				fmt.Print("■")
			} else if grid[vec{x, y}] == open {
				fmt.Print(" ")
			} else if grid[vec{x, y}] == found {
				fmt.Print("!")
			} else {
				fmt.Print("□")
			}
		}
	}
}
