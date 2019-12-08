package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("day8.txt")
	r := bufio.NewReader(file)
	width := 25
	height := 6
	layers := getPixels(width, height, r)

	fewest := (25 * 6) + 1
	part1 := 0
	output := make([]string, width*height)

	for i := len(layers)-1; i >= 0; i -= 1 {
		counts := make(map[int]int)
		layer := layers[i]
		for j, pixel := range layer {
			counts[pixel] += 1
			if pixel != 2 {
				output[j] = pxl(pixel)
			}
		}
		if counts[0] < fewest {
			fewest = counts[0]
			part1 = counts[1] * counts[2]
		}
	}
	fmt.Println(part1)

	// display the image
	for y := 0; y < height; y += 1 {
		for x := 0; x < width; x += 1 {
			i := (y * width) + x
			fmt.Print(output[i])
		}
		fmt.Println()
	}
}

func pxl(i int) string {
	if i == 1 {
		return "#"
	} else {
		return " "
	}
}

func getPixels(width int, height int, r *bufio.Reader) [][]int {
	layers := [][]int{}
	for {
		layer := []int{}
		for p := 0; p < (width * height); p++ {
			if c, _, err := r.ReadRune(); err != nil {
				if err == io.EOF {
					return layers
				}
			} else {
				pix, _ := strconv.Atoi(string(c))
				layer = append(layer, pix)
			}
		}
		layers = append(layers, layer)
	}
}
