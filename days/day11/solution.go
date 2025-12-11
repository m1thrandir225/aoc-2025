package main

import "strings"

func buildGraph(lines []string) map[string]map[string]struct{} {
	graph := make(map[string]map[string]struct{})

	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		i := parts[0]
		v := parts[1]

		if _, exists := graph[i]; !exists {
			graph[i] = make(map[string]struct{})
		}

		for _, j := range strings.Fields(v) {
			if j != "" {
				graph[i][j] = struct{}{}
			}
		}
	}

	return graph
}

func countPaths(graph map[string]map[string]struct{}, start string) int {
	cache := make(map[string]int)

	var walker func(string) int

	walker = func(current string) int {
		if v, ok := cache[current]; ok {
			return v
		}

		if current == "out" {
			return 1
		}

		ans := 0
		for next := range graph[current] {
			ans += walker(next)
		}
		cache[current] = ans
		return ans
	}

	return walker(start)
}

func countMaskedPaths(
	graph map[string]map[string]struct{},
	masks map[string]int,
	start string,
) int {
	type key struct {
		node string
		mask int
	}

	cache := make(map[key]int)

	var walker func(string, int) int
	walker = func(curr string, mask int) int {
		k := key{curr, mask}
		if v, ok := cache[k]; ok {
			return v
		}

		if curr == "out" {
			if mask == 3 { // must have fft + dac
				cache[k] = 1
				return 1
			}
			cache[k] = 0
			return 0
		}

		newMask := mask
		if m, ok := masks[curr]; ok {
			newMask |= m
		}

		ans := 0
		for nxt := range graph[curr] {
			ans += walker(nxt, newMask)
		}

		cache[k] = ans
		return ans
	}

	return walker(start, 0)
}

func part1(input string) any {
	lines := strings.Split(input, "\n")
	graph := buildGraph(lines)

	return countPaths(graph, "you")
}

func part2(input string) any {
	lines := strings.Split(input, "\n")
	graph := buildGraph(lines)

	masks := map[string]int{
		"fft": 1,
		"dac": 2,
	}

	return countMaskedPaths(graph, masks, "svr")
}
