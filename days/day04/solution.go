package main

import (
	"strings"

	"github.com/m1thrandir225/aoc_2025/util"
)

func part1(input string) any {
	grid := make([][]string, 0)
	for _, line := range strings.Split(input, "\n") {
		lineChars := strings.Split(line, "")
		grid = append(grid, lineChars)
	}

	valid, _ := getAdjacentCount(grid)
	return valid
}

func isValidElement(el string) bool {
	return el == "@"
}

func isValidPos(i, j, n, m int) bool {
	if i < 0 || j < 0 || i >= n || j >= m {
		return false
	}
	return true
}

func getAdjacentCount(grid [][]string) (int, [][]string) {
	n := len(grid)
	m := len(grid[0])
	validPositions := 0
	newGrid := util.CopyGrid(grid)
	for i, row := range grid {
		for j, cell := range row {
			if !isValidElement(cell) {
				continue
			}
			adjCount := 0
			if isValidPos(i-1, j-1, n, m) {
				if isValidElement(grid[i-1][j-1]) {
					adjCount++
				}
			}
			if isValidPos(i-1, j, n, m) {
				if isValidElement(grid[i-1][j]) {
					adjCount++
				}
			}
			if isValidPos(i-1, j+1, n, m) {
				if isValidElement(grid[i-1][j+1]) {
					adjCount++
				}
			}
			if isValidPos(i, j-1, n, m) {
				if isValidElement(grid[i][j-1]) {
					adjCount++
				}
			}
			if isValidPos(i, j+1, n, m) {
				if isValidElement(grid[i][j+1]) {
					adjCount++
				}
			}
			if isValidPos(i+1, j-1, n, m) {
				if isValidElement(grid[i+1][j-1]) {
					adjCount++
				}
			}
			if isValidPos(i+1, j, n, m) {
				if isValidElement(grid[i+1][j]) {
					adjCount++
				}
			}
			if isValidPos(i+1, j+1, n, m) {
				if isValidElement(grid[i+1][j+1]) {
					adjCount++
				}
			}

			if adjCount < 4 {
				validPositions++
				newGrid[i][j] = "."
			}
		}
	}
	return validPositions, newGrid
}

func recursiveRemoval(grid [][]string) int {
	maxRemoved, currentRemoved := 0, 0
	for {
		currentRemoved, grid = getAdjacentCount(grid)
		if currentRemoved == 0 {
			return maxRemoved
		}
		maxRemoved += currentRemoved
	}
}

func part2(input string) any {
	grid := make([][]string, 0)
	for _, line := range strings.Split(input, "\n") {
		lineChars := strings.Split(line, "")
		grid = append(grid, lineChars)
	}

	totalRemovals := recursiveRemoval(grid)

	return totalRemovals
}
