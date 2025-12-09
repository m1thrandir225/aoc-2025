package main

import (
	"strconv"
	"strings"
)

type Tile [2]int

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (t *Tile) Area(corner *Tile) int {
	w := abs(t[0]-(*corner)[0]) + 1
	h := abs(t[1]-(*corner)[1]) + 1
	return w * h
}

func part1(input string) any {
	lines := strings.Split(input, "\n")
	tiles := make([]*Tile, len(lines))
	for i, line := range strings.Split(input, "\n") {
		ints := strings.Split(line, ",")
		x1, _ := strconv.Atoi(ints[0])
		x2, _ := strconv.Atoi(ints[1])
		tiles[i] = &Tile{x1, x2}
	}

	largest := 0
	for i := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			largest = max(largest, tiles[i].Area(tiles[j]))
		}
	}

	return largest
}

func part2(input string) any {
	lines := strings.Split(input, "\n")
	tiles := make([]*Tile, len(lines))
	for i, line := range strings.Split(input, "\n") {
		ints := strings.Split(line, ",")
		x1, _ := strconv.Atoi(ints[0])
		x2, _ := strconv.Atoi(ints[1])
		tiles[i] = &Tile{x1, x2}
	}

	return nil
}
