package main

import (
	"strconv"
	"strings"
)

func part1(input string) any {
	dial := 50

	passwordPointer := 0

	for _, line := range strings.Split(input, "\n") {
		toDial, _ := strconv.Atoi(line[1:])
		rotation := []rune(line[:1])[0]
		switch rotation {
		case 'L':
			dial -= toDial
		case 'R':
			dial += toDial
		}

		dial = dial % 100

		if dial == 0 {
			passwordPointer++
		}
	}

	return passwordPointer
}

func part2(input string) any {
	dial := 50
	passwordPointer := 0

	for _, line := range strings.Split(input, "\n") {
		toDial, _ := strconv.Atoi(line[1:])
		for range toDial {
			switch line[0] {
			case 'L':
				dial--
			case 'R':
				dial++
			}
			if dial%100 == 0 {
				passwordPointer++
			}
		}
	}

	return passwordPointer
}
