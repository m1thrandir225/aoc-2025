package main

import (
	"bufio"
	"strconv"
	"strings"
)

func part1(input string) any {
	reader := strings.NewReader(input)

	var lines []string
	scanner := bufio.NewScanner(reader)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		l := scanner.Text()
		lines = append(lines, l)
	}

	var builder strings.Builder

	inPatterns := true

	var patterns []int
	var regionLines []string

	for _, line := range lines {
		if line == "" {
			if inPatterns && builder.Len() > 0 {
				patterns = append(patterns, strings.Count(builder.String(), "#"))
				builder.Reset()
			}
			continue
		}
		if inPatterns {
			if strings.Contains(line, ": ") {
				inPatterns = false
				if builder.Len() > 0 {
					patterns = append(patterns, strings.Count(builder.String(), "#"))
					builder.Reset()
				}
				regionLines = append(regionLines, line)
			} else {
				if builder.Len() > 0 {
					builder.WriteString("\n")
				}
				builder.WriteString(line)
			}
		} else {
			regionLines = append(regionLines, line)
		}
	}
	if inPatterns && builder.Len() > 0 {
		patterns = append(patterns, strings.Count(builder.String(), "#"))
	}
	regions := 0
	for _, line := range regionLines {
		colonIndex := strings.Index(line, ": ")
		if colonIndex == -1 {
			continue
		}

		areaStr := line[:colonIndex]
		numsStr := line[colonIndex+2:]

		xIndex := strings.Index(areaStr, "x")
		if xIndex == -1 {
			continue
		}

		width, _ := strconv.Atoi(areaStr[:xIndex])
		height, _ := strconv.Atoi(areaStr[xIndex+1:])

		area := width * height

		fields := strings.Fields(numsStr)
		size := 0

		for i, field := range fields {
			if i >= len(patterns) {
				break
			}
			if num, err := strconv.Atoi(field); err == nil {
				size += patterns[i] * num
			}
		}

		if size > area {
			continue
		}

		if float64(size)*(1.2) < float64(area) {
			regions++
		}
	}
	return strconv.Itoa(regions)
}

func part2() any {
	final := "Yay"

	return final
}
