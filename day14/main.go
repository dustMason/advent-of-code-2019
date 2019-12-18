package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type chem struct {
	name   string
	amount int
}

func main() {
	graph := make(map[string][]chem)
	dat, _ := ioutil.ReadFile("day14.txt")
	lines := strings.Split(string(dat), "\n")
	for _, line := range lines {
		sides := strings.Split(line, " => ")
		if len(sides) < 2 {
			continue
		}
		inputs := strings.Split(sides[0], ", ")
		var components []chem
		for _, input := range inputs {
			name, amount := ing(input)
			components = append(components, chem{name, amount})
		}
		outName, outAmount := ing(sides[1])
		components = append([]chem{chem{"", outAmount}}, components...)
		graph[outName] = components
	}

	// part 1
	extra := make(map[string]int)
	result := search("FUEL", 1, graph, extra)
	fmt.Println(result)

	// part 2
	extra = make(map[string]int)
	const Stock = 100000
	const OneTrillion = 1000000000000
	var fuel, lastStock int
	for oreConsumed := 0; oreConsumed < OneTrillion; fuel++ {
		lastStock = OneTrillion - oreConsumed
		oreConsumed += search("FUEL", Stock, graph, extra)
	}
	fuel = (fuel - 1) * Stock
	for oreConsumed := 0; oreConsumed < lastStock; fuel++ {
		oreConsumed += search("FUEL", 1, graph, extra)
	}
	fmt.Println(fuel - 1)
}

func search(name string, amount int, graph map[string][]chem, extra map[string]int) int {
	if name == "ORE" {
		return amount
	}
	list, amountProduced := graph[name][1:], graph[name][0].amount
	if extra[name] > 0 {
		if extra[name] >= amount {
			extra[name] -= amount
			return 0
		}
		amount -= extra[name]
		extra[name] = 0
		return search(name, amount, graph, extra)
	}
	needed := (amount-1)/amountProduced + 1
	var oreNeeded int
	for _, r := range list {
		oreNeeded += search(r.name, r.amount*needed, graph, extra)
	}
	if needed*amountProduced-amount > 0 {
		extra[name] += needed*amountProduced - amount
	}
	return oreNeeded
}

func ing(s string) (string, int) {
	p := strings.Split(s, " ")
	a, _ := strconv.Atoi(p[0])
	return p[1], a
}
