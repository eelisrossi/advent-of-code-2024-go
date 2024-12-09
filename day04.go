package main

import (
	"os"
	"strings"
)

// not completely my own answer, had to look at some other solutions on yt first
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

// this is really fucking horrible but again it just works so we go with it
func findCross(grid []string, row int, col int, directions [][2]int) bool {
    first := true

    firstM := false
    firstS := false
    secondM := false
    secondS := false

	for _, dir := range directions {
		rowOffset := dir[0]
		colOffset := dir[1]

        if rowOffset + colOffset == 0 {
            first = false
        } else {
            first = true
        }

		newRow := row + 1*rowOffset
		newCol := col + 1*colOffset

		if !(newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[row])) {
			return false
		}

        if first && !firstM {
            if grid[newRow][newCol] == 77 {
                firstM = true
            }
        }
        if first && !firstS {
            if grid[newRow][newCol] == 83 {
                firstS = true
            }
        }

        if !first && !secondM {
            if grid[newRow][newCol] == 77 {
                secondM = true
            }
        }
        if !first && !secondS {
            if grid[newRow][newCol] == 83 {
                secondS  = true
            }
        }

	}
    if firstM && secondM && firstS && secondS {
        return true
    } else {
        return false
    }

}

func getTestInput() string {
    content, _ := os.ReadFile("./inputs/day4-testinput.txt")
    return string(content)
}

// {-1, -1} {-1, 0} {-1, 1}
// {0, -1}          {0, 1}
// {1, -1}  {1, 0}  {1, 1}
func day04_2(input string) int {
	grid := strings.Split(strings.TrimSpace(input), "\n")
    // grid := strings.Split(strings.TrimSpace(getTestInput()), "\n")
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
