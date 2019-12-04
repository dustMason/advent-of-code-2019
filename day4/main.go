package main

import (
	"fmt"
	"sort"
)

func main() {
	min := 265275
	max := 781584
	count1 := 0
	count2 := 0

	// parts 1 and 2
	for i := min; i < max; i += 1 {
		s := IntToSlice(i, []int{})
		if !sort.IntsAreSorted(s) {
			continue
		}
		last := -1
		for _, n := range s {
			if n == last {
				count1 += 1
				if hasAnyDoubles(s) {
					count2 += 1
				}
				break
			}
			last = n
		}
	}

	fmt.Println(count1)
	fmt.Println(count2)
}

func IntToSlice(n int, sequence []int) []int {
	if n != 0 {
		i := n % 10
		sequence = append([]int{i}, sequence...)
		return IntToSlice(n/10, sequence)
	}
	return sequence
}

func hasAnyDoubles(s []int) bool {
	counts := make(map[int]int)
	for _, n := range s {
		counts[n] += 1
	}
	for _, c := range counts {
		if c == 2 {
			return true
		}
	}
	return false
}
