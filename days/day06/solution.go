package main

import (
	"strconv"
	"strings"
)

func part1(input string) any {
	answer := 0
	numberSlots := make([][]int, 0)
	operationSlots := make([]string, 0)

	lines := strings.Split(input, "\n")

	for _, line := range lines {

		if strings.Contains(line, "*") {
			parts := strings.Split(line, " ")

			for _, part := range parts {
				p := strings.TrimSpace(part)
				if p != "+" && p != "*" {
					continue
				}
				operationSlots = append(operationSlots, p)
			}
			continue
		}
		parts := strings.Split(line, " ")

		nums := make([]int, 0)

		for _, part := range parts {
			if num, err := strconv.Atoi(strings.TrimSpace(part)); err == nil {
				nums = append(nums, num)
			}
		}

		numberSlots = append(numberSlots, nums)
	}

	for i := 0; i < len(operationSlots); i += 1 {
		operation := operationSlots[i]

		total := 0
		for rowIndex, row := range numberSlots {
			if rowIndex == 0 {
				total = row[i]
				continue
			}

			if operation == "+" {
				total += row[i]
			} else {
				total *= row[i]
			}
		}
		answer += total

	}

	return answer
}

func part2(input string) any {
	answer := 0
	numberSlots := make([]int, 0)

	lines := strings.Split(input, "\n")
	lineLength := len(lines[0])

	for col := lineLength - 1; col >= -1; col -= 1 {
		total := 0
		if col >= 0 {
			for i := 0; i < len(lines)-1; i += 1 {
				char := lines[i][col]
				if char == ' ' {
					continue
				}
				num, _ := strconv.Atoi(string(char))
				if total == 0 {
					total = num
				} else {
					total *= 10
					total += num
				}
			}
		}
		if total == 0 {
			localTotal := numberSlots[0]

			if lines[len(lines)-1][col+1] != '+' {
				for i := 1; i < len(numberSlots); i += 1 {
					localTotal *= numberSlots[i]
				}
			} else {
				for i := 1; i < len(numberSlots); i += 1 {
					localTotal += numberSlots[i]
				}
			}
			answer += localTotal
			numberSlots = make([]int, 0)
			continue
		}
		numberSlots = append(numberSlots, total)
	}

	return answer
}
