package main

import (
	"strconv"
	"strings"
)

func day02_1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	safeCount := 0

	for _, line := range lines {
		parts := strings.Fields(line)
		levels := make([]int, len(parts))

		for i, part := range parts {
			num, err := strconv.Atoi(part)
			check(err)
			levels[i] = num
		}

		isIncreasing := true
		isLineSafe := true

		for i := 1; i < len(levels); i++ {
			diff := levels[i] - levels[i-1]

			if diff == 0 {
				isLineSafe = false
				break
			}

			if diff < -3 || diff > 3 {
				isLineSafe = false
				break
			}

			if i == 1 && diff < 0 {
				isIncreasing = false
			} else if (diff < 0 && isIncreasing) || (!isIncreasing && diff > 0) {
				isLineSafe = false
				break
			}
		}

		if isLineSafe {
			safeCount++
		}
	}

	return safeCount
}

func day02_2(input string) int {

	safeCount := 0
	return safeCount
}
