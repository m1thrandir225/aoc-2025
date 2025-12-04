package util

import "fmt"

func PrintGrid(grid [][]string) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
}

func CopyGrid(input [][]string) [][]string {
	newGrid := make([][]string, 0)

	for _, row := range input {
		newRow := make([]string, len(row))
		copy(newRow, row)
		newGrid = append(newGrid, newRow)
	}
	return newGrid
}
