package main

import (
	"sort"
	"strconv"
	"strings"
)

func getSortedLists(input string) ([]int, []int) {
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

	return left, right
}

func day01_1(input string) int {
	left, right := getSortedLists(input)

	total := 0

	for i := 0; i < len(left); i++ {
		d := abs(left[i] - right[i])
		total += d
	}

	return total
}

func day01_2(input string) int {
	left, right := getSortedLists(input)

    similarityScore := 0

	for i := 0; i < len(left); i++ {
		cur := left[i]

		similarCount := 0

		for j := 0; j < len(right); j++ {
			if right[j] == cur {
				similarCount++
			}
		}

        similarity := cur * similarCount

        similarityScore += similarity
	}

	return similarityScore
}
