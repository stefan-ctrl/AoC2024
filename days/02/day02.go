package main

import (
	"AoC2024/util"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines, err := util.ReadFilePerLine("./input/day02.txt")
	if err != nil {
		log.Fatal(err)
	}

	task01(lines)
	task02(lines)
}

func task02(lines []string) {
	safeReports := 0
	for i := range lines {
		intLine := castIntReport(lines[i])
		intLineReversed := reverseIntSlice(intLine)
		if isReportSafe(intLine, intLineReversed) {
			safeReports++
		} else {
			for j := 0; j < len(intLineReversed); j++ {
				dampedLine := introduceDampener(intLine, j)
				dampedLineReversed := reverseIntSlice(dampedLine)
				if isReportSafe(dampedLine, dampedLineReversed) {
					safeReports++
					break
				}
			}
		}
	}
	fmt.Println("Part 2:", safeReports)
}

func introduceDampener(intLine []int, index int) []int {
	dampedLine := make([]int, len(intLine))
	copy(dampedLine, intLine)
	dampedLine = append(dampedLine[:index], dampedLine[index+1:]...)
	return dampedLine
}

func task01(lines []string) {
	safeReports := 0
	for i := range lines {
		intLine := castIntReport(lines[i])
		intLineReversed := reverseIntSlice(intLine)
		if isReportSafe(intLine, intLineReversed) {
			safeReports++
		}
	}
	fmt.Println("Part 1:", safeReports)
}

func isReportSafe(intLine, reversedIntLine []int) bool {

	if !sort.IntsAreSorted(intLine) && !sort.IntsAreSorted(reversedIntLine) {
		return false
	}

	for i := 0; i < len(intLine)-1; i++ {
		diff := math.Abs(float64(intLine[i] - intLine[i+1]))
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func reverseIntSlice(intLine []int) []int {
	var intLineReversed []int

	for i := len(intLine) - 1; i >= 0; i-- {
		intLineReversed = append(intLineReversed, intLine[i])
	}
	return intLineReversed
}

func castIntReport(line string) []int {
	var lineAsInts []int
	splitLine := strings.Split(line, " ")

	for i := range splitLine {
		asInt, err := strconv.Atoi(splitLine[i])
		if err != nil {
			log.Fatal(err)
		}
		lineAsInts = append(lineAsInts, asInt)
	}
	return lineAsInts
}
