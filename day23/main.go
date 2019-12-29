package main

import (
	"fmt"
	"sync"
)

type packet [3]int64 // 0: addr, 1: X, 2: Y

type machine struct {
	in chan int64
}

type router struct {
	machines []machine
	bus      chan packet
}

func (r router) route() {
	for {
		packet := <-r.bus
		addr := packet[0]
		fmt.Printf("[%v] <- %v %v\n", addr, packet[1], packet[2])
		r.machines[int(addr)].in <- packet[1]
		r.machines[int(addr)].in <- packet[2]
	}
}

func (r router) run(codes []int64) {
	var wg sync.WaitGroup
	go r.route()
	for addr, machine := range r.machines {
		wg.Add(1)
		go RunProgram(codes, machine.in, r.bus, &wg, 500000)
		machine.in <- int64(addr)
	}
	wg.Wait()
}

func newRouter() router {
	r := router{
		machines: make([]machine, 50),
		bus:      make(chan packet),
	}
	for i := 0; i < 50; i++ {
		r.machines[i] = machine{
			in: make(chan int64),
		}
	}
	return r
}

func main() {
	codes := LoadIntcode("day23.txt")
	newRouter().run(codes)

	// 8668 is too low
	// 2251799813563177 is too high
}
