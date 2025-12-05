package main

import (
	"strconv"
	"strings"
)

func part1(input string) any {
	splitContent := strings.Split(input, "\n\n")
	freshDatesStr := splitContent[0]
	availableIdsStr := splitContent[1]

	availableIds := strings.Split(availableIdsStr, "\n")
	freshDates := strings.Split(freshDatesStr, "\n")
	freshIds := make(map[int]bool, 0)
	for _, value := range availableIds {
		id, _ := strconv.Atoi(value)
		for _, vDate := range freshDates {
			rangeSplit := strings.Split(vDate, "-")
			start, _ := strconv.Atoi(rangeSplit[0])
			end, _ := strconv.Atoi(rangeSplit[1])

			if id >= start && id <= end {
				if freshIds[id] {
					continue
				} else {
					freshIds[id] = true
				}
			}
		}
	}

	return len(freshIds)
}

func part2(input string) any {
	splitContent := strings.Split(input, "\n\n")
	freshDatesStr := splitContent[0]

	freshDates := strings.Split(freshDatesStr, "\n")

	ranges := make([][2]int, 0, len(freshDates))

	for _, vDate := range freshDates {
		rangeSplit := strings.Split(vDate, "-")
		start, _ := strconv.Atoi(rangeSplit[0])
		end, _ := strconv.Atoi(rangeSplit[1])
		ranges = append(ranges, [2]int{start, end})
	}

	merged := make([]bool, len(ranges))
	for i := 1; i < len(ranges); i++ {
		for j := 0; j < i; j++ {
			if merged[j] {
				continue
			}
			merged[j] = merge(&ranges[i], ranges[j])
		}
	}

	ans := 0
	for i, v := range ranges {
		if !merged[i] {
			ans += v[1] - v[0] + 1
		}
	}

	return ans
}

func merge(r1 *[2]int, r2 [2]int) bool {
	if r1[1] < r2[0] || r1[0] > r2[1] {
		return false
	}
	(*r1)[0] = min((*r1)[0], r2[0])
	(*r1)[1] = max((*r1)[1], r2[1])
	return true
}
