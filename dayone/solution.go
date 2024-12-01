package dayone

import (
	"adventofcode/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func PartOne() {
	start := time.Now()
	puzzleInput, err := util.ReadFileLines("./inputs/dayone.txt")
	if err != nil {
		fmt.Println("could not read input lines: ", err)
		return
	}

	left, right := parseLists(puzzleInput)

	sort.Ints(left)
	sort.Ints(right)

	sum := 0

	for i, left := range left {
		right := right[i]
		diff := right - left

		if diff < 0 {
			sum += -diff
		} else {
			sum += diff
		}
	}

	end := time.Now()
	fmt.Printf("Part one: %d (%dμs)\n", sum, util.DurationInMicroSeconds(start, end))
}

func PartTwo() {
	start := time.Now()
	puzzleInput, err := util.ReadFileLines("./inputs/dayone.txt")
	if err != nil {
		fmt.Println("could not read input lines: ", err)
		return
	}

	left, right := parseLists(puzzleInput)

	// Turn the right list into a map of values to # of times it appears
	rightVals := make(map[int]int)

	for _, v := range right {
		rightVals[v] += 1
	}

	// Calculate sum of similarity scores
	similarityScore := 0
	for _, v := range left {
		similarityScore += v * rightVals[v]
	}

	end := time.Now()
	fmt.Printf("Part two: %d (%dμs)\n", similarityScore, util.DurationInMicroSeconds(start, end))
}

func parseLists(lines []string) ([]int, []int) {
	var left []int
	var right []int

	for _, line := range lines {
		split := strings.Split(line, "   ")

		parsedLeft, _ := strconv.Atoi(split[0])
		left = append(left, parsedLeft)

		parsedRight, _ := strconv.Atoi(split[1])
		right = append(right, parsedRight)
	}

	return left, right
}
