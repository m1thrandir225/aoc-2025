package main

import (
	"strings"
)

type point struct {
	x, y int
}

func part1(input string) any {
	spliterrs := map[point]struct{}{}
	var start point
	last := 0

	for i, line := range strings.Split(input, "\n") {
		for j, lineT := range line {
			switch lineT {
			case 'S':
				start = point{i, j}
			case '^':
				spliterrs[point{i, j}] = struct{}{}
			}
		}
		last = i
	}

	beams := map[point]struct{}{{start.x + 1, start.y}: {}}
	splits := 0

	for i := 2; i < last; i++ {
		nextBeams := map[point]struct{}{}
		for beam := range beams {
			if _, ok := spliterrs[point{i, beam.y}]; ok {
				splits++
				nextBeams[point{i, beam.y - 1}] = struct{}{}
				nextBeams[point{i, beam.y + 1}] = struct{}{}
			} else {
				nextBeams[point{i, beam.y}] = struct{}{}
			}
		}
		beams = nextBeams
	}

	return splits
}

func part2(input string) any {
	spliterrs := map[point]struct{}{}
	var start point
	last := 0
	for i, line := range strings.Split(input, "\n") {
		for j, lineT := range line {
			switch lineT {
			case 'S':
				start = point{i, j}
			case '^':
				spliterrs[point{i, j}] = struct{}{}
			}
		}
		last = i
	}

	beams := map[point]int{{start.x + 1, start.y}: 1}

	for i := 2; i < last; i++ {
		nextBeams := map[point]int{}
		for beam, splits := range beams {
			if _, ok := spliterrs[point{i, beam.y}]; ok {
				nextBeams[point{i, beam.y - 1}] += splits
				nextBeams[point{i, beam.y + 1}] += splits
			} else {
				nextBeams[point{i, beam.y}] += splits
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
