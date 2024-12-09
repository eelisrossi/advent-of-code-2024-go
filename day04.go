package main

import (
	"strings"
)

// not completely my own answer, had to look at some other solutions on yt first
// correct answer 2344
func searchXmas(grid []string, dir [2]int, row int, col int) bool {
	xmas := []byte{88, 77, 65, 83}
	rowOffset := dir[0]
	colOffset := dir[1]

	for i, c := range xmas {
		newRow := row + i*rowOffset
		newCol := col + i*colOffset

		if !(newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[row])) {
			return false
		}
		if grid[newRow][newCol] != c {
			return false
		}
	}
	return true
}

// X = 88
// M = 77
// A = 65
// S = 83
func day04_1(input string) int {
	grid := strings.Split(strings.TrimSpace(input), "\n")
	output := 0

	directions := [][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			foundXmas := false
			for _, dir := range directions {
				foundXmas = searchXmas(grid, dir, row, col)
				if foundXmas {
					output++
				}
			}

		}
	}

	return output
}

// now its a little bit less horrible
func findCross(grid []string, row int, col int, directions [][2]int) bool {
	first := 0
	second := 0

	for _, dir := range directions {
		rowOffset := dir[0]
		colOffset := dir[1]

		diff := rowOffset + colOffset

		newRow := row + 1*rowOffset
		newCol := col + 1*colOffset

		if !(newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[row])) {
			return false
		}

		if diff == 0 {
			if grid[newRow][newCol] == 77 && second-77 != 0 {
				second += 77
			}
			if grid[newRow][newCol] == 83 && second-83 != 0 {
				second += 83
			}
		} else {
			if grid[newRow][newCol] == 77 && first-77 != 0 {
				first += 77
			}
			if grid[newRow][newCol] == 83 && first-83 != 0 {
				first += 83
			}
		}

	}
	if first == 160 && second == 160 {
		return true
	} else {
		return false
	}

}

// {-1, -1} {-1, 0} {-1, 1}
// {0, -1}          {0, 1}
// {1, -1}  {1, 0}  {1, 1}
// correct answer: 1815
func day04_2(input string) int {
	grid := strings.Split(strings.TrimSpace(input), "\n")
	output := 0

	directions := [][2]int{
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			foundMas := false
			if grid[row][col] == 65 {
				foundMas = findCross(grid, row, col, directions)
			}
			if foundMas {
				output++
			}

		}
	}

	return output
}
