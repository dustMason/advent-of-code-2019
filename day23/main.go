package main

import (
	"fmt"
)

type vec struct {
	x, y int64
}

func makeMachines(codes []int64) []chan int64 {
	out := make([]chan int64, 50)
	for i := 0; i < 50; i++ {
		c := make(chan int64)
		go RunProgram(codes, c, 500000)
		c <- int64(i)
		out[i] = c
	}
	return out
}

func main() {
	// part1()
	part2()
}

func part1() {
	codes := LoadIntcode("day23.txt")
	machines := makeMachines(codes)
	queue := make([][]vec, len(machines))
	for i := 0; i < len(machines); i++ {
		queue[i] = make([]vec, 0)
	}

	for {
		for i, machine := range machines {
			var packet vec
			if len(queue[i]) > 0 {
				packet = queue[i][0]
			} else {
				packet = vec{-1, -1}
			}

			select {
			case dest := <-machine:
				x := <-machine
				y := <-machine
				if dest == 255 {
					fmt.Println(y)
					return
				}
				queue[dest] = append(queue[dest], vec{x, y})
			case machine <- packet.x:
				if packet.x != -1 {
					machine <- packet.y
					queue[i] = queue[i][1:]
				}
			}
		}
	}
}

func part2() {
	codes := LoadIntcode("day23.txt")
	machines := makeMachines(codes)
	var nat vec
	queue := make([][]vec, len(machines))
	emptyReceiveCounts := make([]int, len(machines))
	for i := 0; i < len(machines); i++ {
		queue[i] = make([]vec, 0)
	}

	for {
		for i, machine := range machines {
			var packet vec

			if len(queue[i]) > 0 {
				packet = queue[i][0]
			} else {
				packet = vec{-1, -1}
			}

			select {
			case dest := <-machine:
				x := <-machine
				y := <-machine
				if dest == 255 {
					nat = vec{x, y}
				} else {
					queue[dest] = append(queue[dest], vec{x, y})
				}
			case machine <- packet.x:
				if packet.x != -1 {
					machine <- packet.y
					queue[i] = queue[i][1:]
					emptyReceiveCounts[i] = 0
				} else {
					emptyReceiveCounts[i]++
				}
			}
		}

		idle := true
		for _, count := range emptyReceiveCounts {
			if count <= 2 {
				idle = false
				break
			}
		}
		if idle {
			machines[0] <- nat.x
			machines[0] <- nat.y
			fmt.Println(nat)
		}
	}
}
