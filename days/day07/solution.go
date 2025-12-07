package main

import (
	"strings"
)

type Point struct {
	x, y int
}

func part1(input string) any {
	spliterrs := map[Point]struct{}{}
	var start Point
	last := 0

	for i, line := range strings.Split(input, "\n") {
		for j, lineT := range line {
			switch lineT {
			case 'S':
				start = Point{i, j}
			case '^':
				spliterrs[Point{i, j}] = struct{}{}
			}
		}
		last = i
	}

	beams := map[Point]struct{}{{start.x + 1, start.y}: {}}
	splits := 0

	for i := 2; i < last; i++ {
		nextBeams := map[Point]struct{}{}
		for beam := range beams {
			if _, ok := spliterrs[Point{i, beam.y}]; ok {
				splits++
				nextBeams[Point{i, beam.y - 1}] = struct{}{}
				nextBeams[Point{i, beam.y + 1}] = struct{}{}
			} else {
				nextBeams[Point{i, beam.y}] = struct{}{}
			}
		}
		beams = nextBeams
	}

	return splits
}

func part2(input string) any {
	spliterrs := map[Point]struct{}{}
	var start Point
	last := 0
	for i, line := range strings.Split(input, "\n") {
		for j, lineT := range line {
			switch lineT {
			case 'S':
				start = Point{i, j}
			case '^':
				spliterrs[Point{i, j}] = struct{}{}
			}
		}
		last = i
	}

	beams := map[Point]int{{start.x + 1, start.y}: 1}

	for i := 2; i < last; i++ {
		nextBeams := map[Point]int{}
		for beam, splits := range beams {
			if _, ok := spliterrs[Point{i, beam.y}]; ok {
				nextBeams[Point{i, beam.y - 1}] += splits
				nextBeams[Point{i, beam.y + 1}] += splits
			} else {
				nextBeams[Point{i, beam.y}] += splits
			}
		}
		beams = nextBeams
	}

	timelines := 0
	for _, splits := range beams {
		timelines += splits
	}
	return timelines
}
