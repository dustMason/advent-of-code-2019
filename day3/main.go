package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("day3.txt")
	wiresTxt := strings.Split(string(dat), "\n")

	wires := make([][]string, 2)
	for i, wireTxt := range wiresTxt {
		wires[i] = strings.Split(wireTxt, ",")
	}

	points1, path1 := expandPoints(wires[0])
	points2, path2 := expandPoints(wires[1])

	intersects := make([][2]int, 0)

	// part1:
	bestDist := 10000000000000 // so lazy :D
	for loc, _ := range points1 {
		_, ok := points2[loc]
		if ok {
			intersects = append(intersects, loc)
			dist := Abs(loc[0]) + Abs(loc[1])
			if dist < bestDist && dist > 0 {
				bestDist = dist
				fmt.Printf("found new best dist %v\n", bestDist)
			}
		}
	}

	// part2:
	bestPath := 10000000000000 // so lazy :D
	for _, intersect := range intersects {
		dist1 := 0
		dist2 := 0
		for i, ints := range path1 {
			if ints == intersect {
				dist1 = i + 1
			}
		}
		for i, ints := range path2 {
			if ints == intersect {
				dist2 = i + 1
			}
		}
		dist := dist1 + dist2
		if dist < bestPath && dist > 0 {
			bestPath = dist
			fmt.Printf("found new best path %v\n", bestPath)
		}
	}
}

func expandPoints(wires []string) (map[[2]int]bool, [][2]int) {
	points := make(map[[2]int]bool, len(wires[0]))
	path := make([][2]int, 0)
	// NOTE: intentionally don't add the first point to path
	loc := [2]int{0, 0}
	points[loc] = true
	for _, dir := range wires {
		delta := [2]int{0, 0}

		if dir[0] == 'U' {
			delta[1] = -1
		}
		if dir[0] == 'R' {
			delta[0] = 1
		}
		if dir[0] == 'D' {
			delta[1] = 1
		}
		if dir[0] == 'L' {
			delta[0] = -1
		}

		steps, _ := strconv.Atoi(dir[1:])
		for i := 0; i < steps; i++ {
			loc[0] += delta[0]
			loc[1] += delta[1]
			points[loc] = true
			path = append(path, loc)
		}
	}
	return points, path
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
