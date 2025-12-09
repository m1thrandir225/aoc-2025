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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
		line = strings.TrimSpace(line)

		ints := strings.Split(line, ",")
		x1, _ := strconv.Atoi(strings.TrimSpace(ints[0]))
		x2, _ := strconv.Atoi(strings.TrimSpace(ints[1]))
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

	tiles := make([]Tile, len(lines))
	for i, line := range lines {
		line = strings.TrimSpace(line)

		ints := strings.Split(line, ",")
		x1, _ := strconv.Atoi(strings.TrimSpace(ints[0]))
		x2, _ := strconv.Atoi(strings.TrimSpace(ints[1]))
		tiles[i] = Tile{x1, x2}
	}

	ans := 0
	n := len(tiles)

	for i := range n {
		for j := i + 1; j < n; j++ {

			a := tiles[i]
			b := tiles[j]

			w := abs(a[0]-b[0]) + 1
			h := abs(a[1]-b[1]) + 1
			area := w * h

			if area <= ans {
				continue
			}

			possible := true
			xmin := min(a[0], b[0])
			xmax := max(a[0], b[0])
			ymin := min(a[1], b[1])
			ymax := max(a[1], b[1])

			for k := 0; k < len(tiles); k++ {
				l := (k + 1) % n

				x1, y1 := tiles[k][0], tiles[k][1]
				x2, y2 := tiles[l][0], tiles[l][1]

				outside := (x1 <= xmin && x2 <= xmin) || // fully left
					(x1 >= xmax && x2 >= xmax) || // fully right
					(y1 <= ymin && y2 <= ymin) || // fully below
					(y1 >= ymax && y2 >= ymax) // fully above

				if !outside {
					possible = false
					break
				}
			}

			if possible {
				ans = area
			}
		}
	}

	return ans
}
