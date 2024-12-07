package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func day03_1(input string) int {
	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	ri := regexp.MustCompile(`\d{1,3}`)

	allMul := r.FindAllString(input, -1)

	result := 0

	for i := 0; i < len(allMul); i++ {
		numbers := ri.FindAllString(allMul[i], -1)
		x, err := strconv.Atoi(numbers[0])
		check(err)
		y, err := strconv.Atoi(numbers[1])
		check(err)

		res := x * y
		result += res
	}

	return result
}

// this is pretty fucking horrible but it works
func day03_2(input string) int {
	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	ri := regexp.MustCompile(`\d{1,3}`)
    rdo := regexp.MustCompile(`do\(\)`)
    rdont := regexp.MustCompile(`don\'t\(\)`)

	result := 0

	allMulIndex := r.FindAllStringIndex(input, -1)
	allMul := r.FindAllString(input, -1)

    mulIndex := make([]int, len(allMulIndex))

    for i := range allMulIndex {
        mulIndex[i] = allMulIndex[i][0]
    }

    // i only need the index from these since its just a flag
    allDoIndex := rdo.FindAllStringIndex(input, -1)
    allDontIndex := rdont.FindAllStringIndex(input, -1)

    doIndex := make([]int, len(allDoIndex))
    dontIndex := make([]int, len(allDontIndex))

    for i := range allDoIndex {
        doIndex[i] = allDoIndex[i][0]
    }
    for i := range allDontIndex {
        dontIndex[i] = allDontIndex[i][0]
    }

    indexedMult := make([][]int, len(allMul))
    for i := range indexedMult {
        indexedMult[i] = make([]int, 2)
    }

	for i := 0; i < len(allMul); i++ {
		numbers := ri.FindAllString(allMul[i], -1)
		x, err := strconv.Atoi(numbers[0])
		check(err)
		y, err := strconv.Atoi(numbers[1])
		check(err)

        indexedMult[i][0] = mulIndex[i]
        indexedMult[i][1] = x * y
	}


    do := true
    for i := 0; i < len(input); i++ {
        for j := 0; j < len(doIndex); j++ {
            if i == doIndex[j] {
                do = true
            }
        }
        for k := 0; k < len(dontIndex); k++ {
            if i == dontIndex[k] {
                do =false
            }
        }
        for l := 0; l < len(indexedMult); l++ {
            if i == indexedMult[l][0] && do == true {
                result += indexedMult[l][1]
            }
        }
    }

	return result
}
