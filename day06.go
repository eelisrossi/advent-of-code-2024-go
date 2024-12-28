package main

import (
	"fmt"
	"strings"
)

func day06_1(input string) int {
	grid := strings.Split(strings.TrimSpace(input), "\n")
	output := 0

	guard := byte('^')
	obst := byte('#')
	xg := 0
	yg := 0

	for y := range grid {
		for x := range grid {
			if grid[y][x] == guard {
				yg = y
				xg = x
			}
		}
	}

	dir := [2]int{0, -1}

	route := 1

	path := make([][]int, 0)

	fmt.Printf("Starting position: %d, %d\n", yg, xg)
	for (yg+dir[1]) < len(grid) || (yg+dir[1]) >= 0 || (xg+dir[0]) < len(grid[yg]) || (xg+dir[0]) >= 0 {
		walked := false
		if (yg+dir[1]) >= len(grid) || (yg+dir[1]) < 0 || (xg+dir[0]) >= len(grid[yg]) || (xg+dir[0]) < 0 {
			break
		}
		if grid[yg+dir[1]][xg+dir[0]] == obst {
			fmt.Println("Found obstruction")
			dir = turnRight(dir)
		}
		if grid[yg+dir[1]][xg+dir[0]] != obst {
			d := []int{xg, yg}
			for i := range path {
				if path[i][0] == d[0] && path[i][1] == d[1] {
					fmt.Printf("%d == %d\n", path[i], d)
					walked = true
				}
			}

			if !walked {
				path = append(path, d)
				route++
			}

			yg += dir[1]
			xg += dir[0]

			fmt.Printf("x: %d, y: %d\t%d\n", xg, yg, route)
		}

	}

	fmt.Println("out of bounds")
	output = route

	return output
}

func turnRight(dir [2]int) [2]int {
	newDir := [2]int{0, 0}

	right := [2]int{1, 0}
	down := [2]int{0, 1}
	left := [2]int{-1, 0}
	up := [2]int{0, -1}

	if dir == up {
		newDir = right
	} else if dir == right {
		newDir = down
	} else if dir == down {
		newDir = left
	} else if dir == left {
		newDir = up
	}

	return newDir
}

func day06_2(input string) int {
	// lines := strings.Split(strings.TrimSpace(input), "\n")
	output := 0

	return output
}
