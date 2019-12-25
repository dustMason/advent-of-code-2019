package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/logrusorgru/aurora"
)

var deltas = []vec{
	vec{0, 1},
	vec{-1, 0},
	vec{0, -1},
	vec{1, 0},
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
func (t tile) neighbors(grid [][]byte, keysHeld map[byte]bool) []tile {
	out := []tile{}
	for _, delta := range deltas {
		o := t.add(delta)
		if o.x >= 0 && o.y >= 0 && o.x < len(grid[0]) && o.y < len(grid) {
			val := grid[o.y][o.x]

			// if its a door and we don't have the key, skip
			if _, ok := keysHeld[val+32]; isDoor(val) && !ok {
				continue
			}

			// if its a wall, skip
			if val == wall {
				continue
			}

			out = append(out, tile{vec{o.x, o.y}, val})
		}
	}
	return out
}

func (t tile) String() string {
	return fmt.Sprintf("{%v,%v : %v}", t.x, t.y, string(t.val))
}

func main() {
	grid := load()

	// as we perform the DFS, keep track of step counts for each KEY found. if we encounter a KEY again with a
	// better step count, replace the path

	loc := tile{vec{40, 40}, entrance}
	keys := findKeys(grid)
	seen := make(map[vec]bool)
	keysHeld := map[byte]bool{} // (add 32 to lowerCase byte to get upperCase)
	neighbs := loc.neighbors(grid, keysHeld)
	for {
		if len(neighbs) == 0 {
			break
		}

		printGrid(grid, seen, keysHeld)

		n := neighbs[0]
		neighbs = neighbs[1:]
		for _, t := range n.neighbors(grid, keysHeld) {
			if _, ok := seen[t.vec]; !ok {
				neighbs = append(neighbs, t)
				seen[t.vec] = true
				if _, ok := keys[t.vec]; ok {
					keysHeld[t.val] = true
				}
			}
		}
	}
}

func findKeys(grid [][]byte) map[vec]byte {
	keys := make(map[vec]byte)
	for y, row := range grid {
		for x, t := range row {
			if t >= 'a' && t <= 'z' {
				keys[vec{x, y}] = t
			}
		}
	}
	return keys
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

func printGrid(grid [][]byte, seen map[vec]bool, keysHeld map[byte]bool) {
	fmt.Printf("\033[0;0H") // set cursor to terminal position 0,0
	for y, row := range grid {
		fmt.Println()
		for x, b := range row {
			if _, ok := seen[vec{x, y}]; ok {
				fmt.Print(Cyan("+"))
			} else if b == space {
				fmt.Print(" ")
			} else if b == wall {
				fmt.Print(string(b))
			} else {
				fmt.Print(Bold(Red(string(b))))
			}
		}
	}

	keysSorted := []string{}
	for b, _ := range keysHeld {
		keysSorted = append(keysSorted, string(b))
	}
	sort.Strings(keysSorted)
	fmt.Println(keysSorted)
}
