package main

import (
	"AoC2024/util"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines, err := util.ReadFilePerLine("./input/day01.txt")
	if err != nil {
		log.Fatal(err)
	}

	day01_task1(lines)
	day01_task2(lines)
}

func day01_task1(lines []string) {
	leftSide := make([]int, len(lines))
	rightSide := make([]int, len(lines))
	for i := range lines {
		left, right := receiveSides(lines[i])
		leftSide[i] = left
		rightSide[i] = right
	}

	sort.Ints(leftSide)
	sort.Ints(rightSide)

	total := 0
	for i := range leftSide {
		total += int(math.Abs(float64(leftSide[i] - rightSide[i])))
	}
	println(total)
}

func day01_task2(lines []string) {

	leftMap := make(map[int]int)
	rightMap := make(map[int]int)

	for i := range lines {
		left, right := receiveSides(lines[i])
		if _, ok := leftMap[left]; !ok {
			leftMap[left] = 1
		} else {
			leftMap[left]++
		}

		if _, ok := rightMap[right]; !ok {
			rightMap[right] = 1
		} else {
			rightMap[right]++
		}
	}

	total := 0

	for i := range leftMap {
		total += i * rightMap[i]
	}
	println(total)
}

func receiveSides(line string) (int, int) {
	sides := strings.Split(line, " ")
	left, _ := strconv.Atoi(sides[0])
	right, _ := strconv.Atoi(sides[3])

	return left, right
}
