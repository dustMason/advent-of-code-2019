package main

import (
	"bufio"
	"fmt"
	"os"
)

var deltas = []vec{
	{0, 1},
	{-1, 0},
	{0, -1},
	{1, 0},
}

const (
	space    = byte('.')
	wall     = byte('#')
	entrance = byte('@')
)

type vec struct {
	x, y int
}

func (v vec) add(other vec) vec {
	return vec{v.x + other.x, v.y + other.y}
}

type tile struct {
	vec
	val byte
}

func isDoor(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func isKey(b byte) bool {
	return b >= 'a' && b <= 'z'
}

// all legal (in-bounds and non-wall) adjacent tiles
func (t tile) neighbors(grid [][]byte) []tile {
	out := []tile{}
	for _, delta := range deltas {
		o := t.add(delta)
		if o.x >= 0 && o.y >= 0 && o.x < len(grid[0]) && o.y < len(grid) && grid[o.y][o.x] != wall {
			out = append(out, tile{vec{o.x, o.y}, grid[o.y][o.x]})
		}
	}
	return out
}

func (t tile) String() string {
	return fmt.Sprintf("{%v,%v : %v}", t.x, t.y, string(t.val))
}

type node struct {
	tile
	dist int
}

func neighbors(grid [][]byte, loc tile) map[byte]node {
	neighbs := make(map[byte]node)
	var first node
	queue := []node{{tile: loc, dist: 0}}
	visited := map[tile]bool{loc: true}
	for len(queue) > 0 {
		first, queue = queue[0], queue[1:]
		for _, adj := range first.neighbors(grid) {
			if visited[adj] {
				continue
			}
			visited[adj] = true
			nd := node{tile: adj, dist: first.dist + 1}
			if isDoor(adj.val) || isKey(adj.val) {
				neighbs[adj.val] = nd
			} else {
				queue = append(queue, nd)
			}
		}
	}
	return neighbs
}

func find(b byte, grid [][]byte) (vec, bool) {
	for y, row := range grid {
		for x, t := range row {
			if t == b {
				return vec{x, y}, true
			}
		}
	}
	return vec{0, 0}, false
}

func copied(x []byte) []byte {
	y := make([]byte, len(x))
	for i, v := range x {
		y[i] = v
	}
	return y
}

func remove(bytes []byte, i int) []byte {
	return append(copied(bytes[:i]), bytes[i+1:]...)
}

func copiedPts(ps []tile) []tile {
	qs := make([]tile, len(ps))
	copy(qs, ps)
	return qs
}

func replace(ps []tile, i int, p tile) []tile {
	ps = copiedPts(ps)
	ps[i] = p
	return ps
}

func copyGrid(grid [][]byte) [][]byte {
	duplicate := make([][]byte, len(grid))
	for i := range grid {
		duplicate[i] = make([]byte, len(grid[i]))
		copy(duplicate[i], grid[i])
	}
	return duplicate
}

var memo = make(map[string]int)

func cacheKey(pos []tile, need []byte) string {
	return fmt.Sprintf("%v-%s", pos, string(need))
}

func multiShortestPath(pos []tile, grid [][]byte, need []byte) int {
	mk := cacheKey(pos, need)
	if d, ok := memo[mk]; ok {
		return d
	}
	if len(need) == 0 {
		return 0
	}
	shortest := 0
	for j, from := range pos {
		neighbs := neighbors(grid, from)
		for i, key := range need {
			nd, ok := neighbs[key]
			if !ok {
				continue
			}

			grid2 := copyGrid(grid)

			grid2[nd.y][nd.x] = space
			if doorVec, ok := find(nd.val-32, grid); ok {
				grid2[doorVec.y][doorVec.x] = space
			}
			subSteps := multiShortestPath(replace(pos, j, tile{vec: vec{nd.x, nd.y}, val: 0}), grid2, remove(need, i)) + nd.dist
			if subSteps >= 0 && (shortest == 0 || subSteps < shortest) {
				shortest = subSteps
			}
		}
	}
	if shortest > 0 {
		memo[mk] = shortest
		return shortest
	}
	return -1
}

func findKeys(grid [][]byte) []byte {
	k := []byte{}
	for _, row := range grid {
		for _, t := range row {
			if t >= 'a' && t <= 'z' {
				k = append(k, t)
			}
		}
	}
	return k
}

func load() [][]byte {
	file, _ := os.Open("day18.txt")
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

func updateGridForPart2(grid [][]byte, start tile) [][]byte {
	pos := vec{start.x, start.y}
	grid2 := copyGrid(grid)
	grid2[pos.y][pos.x] = wall
	for _, delta := range deltas {
		step := pos.add(delta)
		grid2[step.y][step.x] = wall
	}
	return grid2
}

func main() {
	grid := load()

	loc := tile{vec{40, 40}, entrance}
	part1 := multiShortestPath([]tile{loc}, grid, findKeys(grid))
	fmt.Println(part1)

	grid2 := updateGridForPart2(grid, loc)
	startingPositions := []tile{
		{vec: loc.add(vec{-1, -1}), val: 1},
		{vec: loc.add(vec{-1, 1}), val: 2},
		{vec: loc.add(vec{1, -1}), val: 3},
		{vec: loc.add(vec{1, 1}), val: 4},
	}
	part2 := multiShortestPath(startingPositions, grid2, findKeys(grid2))
	fmt.Println(part2)
}
