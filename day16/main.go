package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	// part 1
	list1 := transform(load(), 100)
	fmt.Println(list1[:8])

	// part 2
	dat := load()
	offset := dat[0]*1000000 + dat[1]*100000 + dat[2]*10000 + dat[3]*1000 + dat[4]*100 + dat[5]*10 + dat[6]
	list2 := fastTransform(dat)
	fmt.Println(list2[offset : offset+8])
}

func load() []int {
	dat, _ := ioutil.ReadFile("day16.txt")
	ints := make([]int, len(dat))
	for i, c := range dat {
		n, _ := strconv.Atoi(string(c))
		ints[i] = n
	}
	return ints
}

func transform(list []int, phases int) []int {
	mask := []int{0, 1, 0, -1}
	for i := 0; i < phases; i++ {
		newlist := make([]int, len(list))
		for j := 0; j < len(list); j++ {
			total := 0
			for k := j; k < len(list); k++ {
				total += list[k] * mask[((k+1)/(j+1))%4]
			}
			newlist[j] = abs(total % 10)
		}
		list = newlist
	}
	return list
}

func fastTransform(list []int) []int {
	newlist := make([]int, len(list)*10000)
	for i := 0; i < len(newlist); i += len(list) {
		copy(newlist[i:], list)
	}
	list = newlist
	newlist = make([]int, len(list))
	for i := 0; i < 100; i++ {
		newlist[len(list)-1] = list[len(list)-1]
		for j := len(list) - 2; j >= 0; j-- {
			newlist[j] = (list[j] + newlist[j+1]) % 10
		}
		list = newlist
	}
	return list
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
