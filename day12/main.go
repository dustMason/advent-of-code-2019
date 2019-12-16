package main

import (
	"fmt"
	"os"
)

type vec3 struct {
	x int
	y int
	z int
}

type moon struct {
	position vec3
	velocity vec3
}

func (m *moon) addVelocity(x, y, z int) {
	m.velocity.x += x
	m.velocity.y += y
	m.velocity.z += z
}

func (m *moon) applyGravity(other *moon) {
	x, y, z := m.position.x, m.position.y, m.position.z
	ox, oy, oz := other.position.x, other.position.y, other.position.z
	dx, dy, dz := delta(x, ox), delta(y, oy), delta(z, oz)
	m.addVelocity(-dx, -dy, -dz)
	other.addVelocity(dx, dy, dz)
}

func (m *moon) applyVelocity() {
	m.position.x += m.velocity.x
	m.position.y += m.velocity.y
	m.position.z += m.velocity.z
}

func (m *moon) totalEnergy() int {
	potential := abs(m.position.x) + abs(m.position.y) + abs(m.position.z)
	kinetic := abs(m.velocity.x) + abs(m.velocity.y) + abs(m.velocity.z)
	return potential * kinetic
}

func main() {
	// part 1
	moons := loadMoons()
	combos := [][2]int{}
	comb(len(moons), 2, func(pair []int) {
		combos = append(combos, [2]int{pair[0], pair[1]})
	})
	for i := 0; i < 1000; i++ {
		tick(combos, moons)
	}
	totalEnergy := 0
	for _, m := range moons {
		totalEnergy += m.totalEnergy()
	}
	fmt.Println(totalEnergy)

	// part 2
	moons = loadMoons()
	period, found := vec3{}, 0
	for i := 1; found < 3; i++ {
		tick(combos, moons)

		if period.x == 0 {
			inv := false
			for m := range moons {
				if moons[m].velocity.x != 0 {
					inv = true
					break
				}
			}

			if !inv {
				period.x = i * 2
				found++
			}
		}

		if period.y == 0 {
			inv := false
			for m := range moons {
				if moons[m].velocity.y != 0 {
					inv = true
					break
				}
			}

			if !inv {
				period.y = i * 2
				found++
			}
		}

		if period.z == 0 {
			inv := false
			for m := range moons {
				if moons[m].velocity.z != 0 {
					inv = true
					break
				}
			}

			if !inv {
				period.z = i * 2
				found++
			}
		}
	}

	fmt.Println(lcm(lcm(period.x, period.y), period.z))
}

func tick(combos [][2]int, moons []*moon) {
	for _, c := range combos {
		moons[c[0]].applyGravity(moons[c[1]])
	}
	for _, m := range moons {
		m.applyVelocity()
	}
}

func loadMoons() []*moon {
	input, _ := os.Open("day12.txt")
	moons := make([]*moon, 0)
	x, y, z := 0, 0, 0
	for {
		_, err := fmt.Fscanf(input, "<x=%d, y=%d, z=%d>", &x, &y, &z)
		if err != nil {
			break
		}
		moons = append(moons, &moon{position: vec3{x, y, z}, velocity: vec3{}})
		_, _ = fmt.Fscanln(input)
	}
	return moons
}

// https://rosettacode.org/wiki/Combinations#Go
func comb(n, m int, emit func([]int)) {
	s := make([]int, m)
	last := m - 1
	var rc func(int, int)
	rc = func(i, next int) {
		for j := next; j < n; j++ {
			s[i] = j
			if i == last {
				emit(s)
			} else {
				rc(i+1, j+1)
			}
		}
		return
	}
	rc(0, 0)
}

func delta(l, r int) int {
	if l == r {
		return 0
	}
	if l > r {
		return 1
	}
	return -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
