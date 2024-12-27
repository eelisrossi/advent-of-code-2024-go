package main

import (
	"strconv"
	"strings"
)

func getRulesAndUpdates(lines []string) ([][]int, [][]int) {
	rules := make([][]int, 0)
	updates := make([][]int, 0)
	gotLines := false

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

	return rules, updates
}

func checkUpdate(update []int, rules [][]int) bool {
	for index, val := range update {
		for _, rule := range rules {
			if val == rule[0] {
				if !checkRule(update, rule, index) {
					return false
				}
			}
		}
	}

	return true
}

func checkRule(update []int, rule []int, index int) bool {
	for i := 0; i < index; i++ {
		if update[i] == rule[1] {
			return false
		}
	}

	return true
}

func day05_1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	output := 0

	rules, updates := getRulesAndUpdates(lines)

	for _, update := range updates {
		isSafe := true
		isSafe = checkUpdate(update, rules)
		if isSafe {
			middleIndex := len(update) / 2
			middlePage := update[middleIndex]
			output += middlePage
		}
	}

	return output
}

func day05_2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	output := 0

	rules, updates := getRulesAndUpdates(lines)

	for _, update := range updates {
		isSafe := true
		isSafe = checkUpdate(update, rules)
		if !isSafe {
			correct := fixOrdering(update, rules)
			middleIndex := len(correct) / 2
			middlePage := correct[middleIndex]
			output += middlePage
		}
	}

	return output
}

func fixOrdering(update []int, rules [][]int) []int {
	correctUpdate := make([]int, len(update))
    newRules := getRightRules(update, rules)

    for i, value := range update {
        if correctUpdate[i] == 0 {
            correctUpdate[i] = value
        }
        for _, rule := range newRules {
            before := rule[0]
            after := rule[1]
            if value == before {
                for index := 0; index < i; index++ {
                    if correctUpdate[index] == after {
                        a := correctUpdate[index]
                        b := correctUpdate[i]
                        correctUpdate[index] = b
                        correctUpdate[i] = a
                    }
                }
            }
        }
    }

    safe := checkUpdate(correctUpdate, rules)
    if !safe {
        correctUpdate = fixOrdering(correctUpdate, rules)
    }

	return correctUpdate
}

func getRightRules(update []int, rules [][]int) [][]int {
	tempRules := make([][]int, 0)
	newRules := make([][]int, 0)

	for _, rule := range rules {
		for _, u := range update {
			if rule[0] == u || rule[1] == u {
				tempRules = append(tempRules, rule)
			}
		}
	}

    for _, r := range tempRules {
        exists := false
        if len(newRules) == 0 {
            newRules = append(newRules, r)
        }
        if len(newRules) > 0 {
            for _, nr := range newRules {
                if r[0] == nr[0] && r[1] == nr[1] {
                    exists = true
                }
            }
        }
        if !exists {
            newRules = append(newRules, r)
        }
    }

    return newRules
}
