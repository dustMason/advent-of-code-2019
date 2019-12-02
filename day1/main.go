package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("day1.txt")
	inputs := strings.Split(string(dat), "\n")
	part1 := 0
	part2 := 0
	for _, input := range inputs {
		mass, _ := strconv.Atoi(input)
		part1 += (mass / 3) - 2
		part2 += totalFuel(mass)
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func totalFuel(mass int) int {
	fuel := (mass / 3) - 2
	if fuel <= 0 {
		return 0
	}
	return totalFuel(fuel) + fuel
}
