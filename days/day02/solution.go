package main

import (
	"strconv"
	"strings"
)

func part1(input string) any {
	count := 0
	for _, line := range strings.Split(input, ",") {
		line = strings.Trim(line, "\n")
		ranges := strings.Split(line, "-")
		r1, _ := strconv.Atoi(ranges[0])
		r2, _ := strconv.Atoi(ranges[1])

		count += countNumbers(r1, r2)
	}

	return count
}

func countNumbers(start, end int) int {
	sum := 0

	for num := start; num <= end; num++ {
		s := strconv.Itoa(num)
		length := len(s)

		if length%2 != 0 {
			continue
		}

		half := length / 2
		first := s[:half]
		second := s[half:]

		if first == second {
			sum += num
		}
	}
	return sum
}

func part2(input string) any {
	count := 0

	for _, line := range strings.Split(input, ",") {
		line = strings.Trim(line, "\n")
		ranges := strings.Split(line, "-")
		r1, _ := strconv.Atoi(ranges[0])
		r2, _ := strconv.Atoi(ranges[1])

		count += countPart2(r1, r2)
	}

	return count
}

func countPart2(start, end int) int {
	sum := 0

	for num := start; num <= end; num++ {
		s := strconv.Itoa(num)
		n := len(s)

		for block := 1; block <= n/2; block++ {
			if n%block != 0 {
				continue // divide full len
			}

			candidate := s[:block] // 123123123, canditate: 1 -> 12 -> 123...
			repeats := n / block
			valid := true

			// check if entire string is candidate repeated
			for i := 1; i < repeats; i++ {
				if s[i*block:(i+1)*block] != candidate {
					valid = false
					break
				}
			}

			if valid && repeats >= 2 {
				sum += num
				break
			}
		}
	}

	return sum
}
