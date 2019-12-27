package main

import (
	"bufio"
	"fmt"
	"os"
)

var deltas = []vec{
	{0, 1, 0},
	{-1, 0, 0},
	{0, -1, 0},
	{1, 0, 0},
}

type vec struct {
	x, y, z int
}

func (v vec) add(other vec) vec {
	return vec{v.x + other.x, v.y + other.y, v.z + other.z}
}

// all legal (in-bounds and non-wall) adjacent tiles
func (v vec) neighbors(grid [][]byte) []vec {
	out := []vec{}
	for _, delta := range deltas {
		o := v.add(delta)
		if o.x >= 0 && o.y >= 0 && o.x < len(grid[0]) && o.y < len(grid) && walkable(grid[o.y][o.x]) {
			out = append(out, vec{o.x, o.y, o.z})
		}
	}
	return out
}

type node struct {
	pos  vec
	dist int
}

type portal struct {
	name  string
	outer bool
}

func (p portal) to() string {
	if p.name[2] == '1' {
		return p.name[0:2] + "2"
	} else {
		return p.name[0:2] + "1"
	}
}

func main() {
	part1()
	part2()
}

func part1() {
	grid := load()
	portals, start, end := portalLocations(grid)
	// graph := portalGraph(grid, portals)

	var visited = map[vec]bool{start: true}
	var queue = []node{{pos: start}}
	var n node
	for {
		n, queue = queue[0], queue[1:]
		for _, next := range n.pos.neighbors(grid) {
			if next == end {
				fmt.Println(n.dist + 1)
				return
			}

			if !visited[next] {
				visited[next] = true
				steps := 1
				p, ok := portals[next]
				if ok {
					steps++ // use a step to cross to the other portal
					for v, s := range portals {
						if s.name == p.to() {
							next = v
							break
						}
					}
				}

				queue = append(queue, node{next, n.dist + steps})
			}
		}
	}
}

func part2() {
	grid := load()
	portals, start, end := portalLocations(grid)

	var visited = map[vec]bool{start: true}
	var queue = []node{{pos: start}}
	var n node
	for {
		n, queue = queue[0], queue[1:]
		for _, next := range n.pos.neighbors(grid) {
			if next == end {
				fmt.Println(n.dist + 1)
				return
			}

			if !visited[next] {
				visited[next] = true
				steps := 1
				p, ok := portals[vec{next.x, next.y, 0}]

				if ok && (n.pos.z > 0 || !p.outer) {
					steps++ // use a step to cross to the other portal
					for v, s := range portals {
						if s.name == p.to() {
							next = vec{v.x, v.y, n.pos.z}
							if p.outer {
								next.z--
							} else {
								next.z++
							}
							visited[next] = true
							break
						}
					}
				}

				queue = append(queue, node{next, n.dist + steps})
			}
		}
	}
}

func portalLocations(grid [][]byte) (map[vec]portal, vec, vec) {
	portals := make(map[vec]portal)
	seen := make(map[string]bool)
	var start, end vec
	for y, row := range grid {
		for x, tile := range row {
			if isMapOrSpace(tile) {
				continue
			}
			if y == len(grid)-1 || x == len(grid[y])-1 {
				continue
			}

			secondChar := grid[y+1][x]
			pDelta := vec{0, 1, 0}
			if isMapOrSpace(secondChar) {
				secondChar = grid[y][x+1]
				pDelta = vec{1, 0, 0}
			}
			if isMapOrSpace(secondChar) {
				continue
			}
			p := closestWalkable([2]vec{{x, y, 0}, vec{x, y, 0}.add(pDelta)}, grid)

			name := string([]byte{tile, secondChar})
			if _, ok := seen[name]; ok {
				name = name + "2"
			} else {
				seen[name] = true
				name = name + "1"
			}

			if name == "AA1" {
				start = p
			}
			if name == "ZZ1" {
				end = p
			}

			portals[p] = portal{name, isOuter(p, grid)}
		}
	}
	return portals, start, end
}

func load() [][]byte {
	file, _ := os.Open("day20.txt")
	defer file.Close()
	grid := [][]byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r := scanner.Text()
		row := make([]byte, len(r))
		for i, b := range []byte(r) {
			row[i] = b
		}
		grid = append(grid, row)
	}
	return grid
}

func isMapOrSpace(b byte) bool {
	return b == ' ' || b == '.' || b == '#'
}

func isOuter(p vec, grid [][]byte) bool {
	return p.y == 2 || p.x == 2 || p.y == len(grid)-3 || p.x == len(grid[0])-3
}

func closestWalkable(portalLocs [2]vec, grid [][]byte) vec {
	dot := vec{}
	for _, loc := range portalLocs {
		for _, delta := range deltas {
			look := loc.add(delta)
			if look.x < 0 || look.y < 0 || look.y > len(grid)-1 || look.x > len(grid[0])-1 {
				continue
			}
			if grid[look.y][look.x] == '.' {
				dot = look
			}
		}
	}
	return dot
}

func walkable(b byte) bool {
	return b != ' ' && b != '#'
}
