package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

const deckSize int = 10007

func main() {
	part1()
	part2()
}

func part1() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	deck := [deckSize]int64{}
	for i := 0; i < deckSize; i++ {
		deck[i] = int64(i)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		deck = shuffle(scanner.Text(), deck)
	}

	for i, card := range deck {
		if card == 2019 {
			fmt.Printf("%v ", i)
			break
		}
	}
}

func part2() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	n, iter := big.NewInt(119315717514047), big.NewInt(101741582076661)
	offset, increment := big.NewInt(0), big.NewInt(1)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		op := scanner.Text()
		switch {
		case op == "deal into new stack":
			increment.Mul(increment, big.NewInt(-1))
			offset.Add(offset, increment)
		case op[:3] == "cut":
			i, _ := strconv.Atoi(op[4:])
			offset.Add(offset, big.NewInt(0).Mul(big.NewInt(int64(i)), increment))
		case op[:19] == "deal with increment":
			i, _ := strconv.Atoi(op[20:])
			increment.Mul(increment, big.NewInt(0).Exp(big.NewInt(int64(i)), big.NewInt(0).Sub(n, big.NewInt(2)), n))
		}
	}

	finalIncr := big.NewInt(0).Exp(increment, iter, n)
	finalOffs := big.NewInt(0).Exp(increment, iter, n)
	finalOffs.Sub(big.NewInt(1), finalOffs)
	invmod := big.NewInt(0).Exp(big.NewInt(0).Sub(big.NewInt(1), increment), big.NewInt(0).Sub(n, big.NewInt(2)), n)
	finalOffs.Mul(finalOffs, invmod)
	finalOffs.Mul(finalOffs, offset)

	answer := big.NewInt(0).Mul(big.NewInt(2020), finalIncr)
	answer.Add(answer, finalOffs)
	answer.Mod(answer, n)

	fmt.Println(answer)
}

func shuffle(i string, deck [deckSize]int64) [deckSize]int64 {
	if strings.HasPrefix(i, "cut") {
		var v int
		_, _ = fmt.Sscanf(i, "cut %d", &v)
		if v > 0 { // take v cards off the top and put them on the bottom
			newDeck := append(deck[v:], deck[:v-1]...)
			copy(deck[:], newDeck)
		} else { // take v cards off the bottom and put them on the bottom
			v = v * -1
			newDeck := append(deck[deckSize-v:], deck[:deckSize-v]...)
			copy(deck[:], newDeck)
		}
	}

	if strings.HasPrefix(i, "deal with increment") {
		var v int
		_, _ = fmt.Sscanf(i, "deal with increment %d", &v)
		newDeck := [deckSize]int64{}
		index := 0
		for i := 0; i < deckSize; i++ {
			newDeck[index] = deck[i]
			index = (index + v) % deckSize
		}
		deck = newDeck
	}

	if strings.HasPrefix(i, "deal into new") { // aka reverse
		for i, j := 0, deckSize-1; i < j; i, j = i+1, j-1 {
			deck[i], deck[j] = deck[j], deck[i]
		}
	}
	return deck
}
