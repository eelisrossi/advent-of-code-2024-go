package main

import (
	"strconv"
	"strings"
)

func day05_1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rules := make([][]int, 0)
	updates := make([][]int, 0)
	gotLines := false
	output := 0

	for _, line := range lines {
		if !gotLines && len(line) != 0 {
			ruleSplit := strings.Split(line, "|")
			r1, _ := strconv.Atoi(ruleSplit[0])
			r2, _ := strconv.Atoi(ruleSplit[1])
			rule := make([]int, 2)
			rule[0] = r1
			rule[1] = r2
			rules = append(rules, rule)
		}
		if gotLines {
            updateSplit := strings.Split(line, ",")
            update := make([]int, len(updateSplit))
            for i, u := range updateSplit {
               update[i], _ = strconv.Atoi(u)
            }
			updates = append(updates, update)
		}
		if len(line) == 0 && !gotLines {
			gotLines = true
		}
	}

    for _, update := range updates {
        isSafe := true
        for index, val := range update {

            for _, rule := range rules {
                if val == rule[0] {
                    for i := 0; i < index; i++ {
                        if update[i] == rule[1] {
                            isSafe = false
                            break
                        }
                    }
                }
            }
            if !isSafe {
                break
            }
        }
        if isSafe {
            middleIndex := len(update) / 2
            middlePage := update[middleIndex]
            output += middlePage
        }
    }



	return output
}

func day05_2(input string) int {
	// lines := strings.Split(strings.TrimSpace(input), "\n")
	output := 0

	return output
}
