package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("day6.txt")
	orbitsTxt := strings.Split(string(dat), "\n")
	graph := make(map[string]string)

	for _, s := range orbitsTxt {
		pair := strings.Split(s, ")")
		graph[pair[1]] = pair[0]
	}

	// part 1
	total := 0
	for _, v := range graph {
		ok := true
		for ok {
			v, ok = graph[v]
			total += 1
		}
	}
	fmt.Println(total)

	// part 2
	you := []string{graph["YOU"]}
	san := []string{graph["SAN"]}
	for {
		u := intersection(you, san)
		if len(u) > 0 {
			commonNode := u[0]
			fmt.Println(indexOf(commonNode, san) + indexOf(commonNode, you))
			break
		}
		you = append(you, graph[you[len(you)-1]])
		san = append(san, graph[san[len(san)-1]])
	}
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func intersection(a []string, b []string) []string {
	set := make([]string, 0)
	hash := make(map[string]bool)

	for i := 0; i < len(a); i++ {
		el := a[i]
		hash[el] = true
	}

	for i := 0; i < len(b); i++ {
		el := b[i]
		if _, found := hash[el]; found {
			set = append(set, el)
		}
	}

	return set
}
