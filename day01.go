package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
    mask := x >> 31
    ax := x ^ mask
    ax = ax - mask
    return ax
}

func day01_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var left []int
	var right []int

	for _, line := range lines {
		lineSplit := strings.Split(strings.TrimSpace(line), "   ")
		leftInt, err := strconv.Atoi(lineSplit[0])
		check(err)
		rightInt, err := strconv.Atoi(lineSplit[1])
		check(err)

		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	sort.Ints(left)
	sort.Ints(right)

    total := 0

	for i := 0; i < len(left); i++ {
        d := abs(left[i] - right[i])
        total += d
	}

    fmt.Printf("Result for day 1 part 1 is: %d\n", total)
}
