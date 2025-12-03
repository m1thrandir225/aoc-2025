package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part1(input string) any {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		joltage, _ := strconv.Atoi(findJoltage(line))

		sum += joltage
	}

	return sum
}

func findJoltage(s string) string {
	maxRight := -1
	best := -1

	for i := len(s) - 1; i >= 0; i-- {
		current := int(s[i] - '0')

		if maxRight != -1 {
			value := current*10 + maxRight
			if value > best {
				best = value
			}
		}

		if current > maxRight {
			maxRight = current
		}
	}

	if best == -1 {
		return ""
	}

	return fmt.Sprintf("%02d", best)
}

func part2(input string) any {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		joltage, _ := strconv.Atoi(largestSubsequenceK(line, 12))
		sum += joltage
	}

	return sum
}

func largestSubsequenceK(s string, k int) string {
	stack := make([]byte, 0)
	toRemove := len(s) - k

	for i := 0; i < len(s); i++ {
		c := s[i]

		for toRemove > 0 && len(stack) > 0 && stack[len(stack)-1] < c {
			stack = stack[:len(stack)-1]
			toRemove--
		}

		stack = append(stack, c)
	}

	// Keep only first k digits
	return string(stack[:k])
}
