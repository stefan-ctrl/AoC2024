package days

import (
	"AoC2024/util"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Day02() {
	lines, err := util.ReadFilePerLine("./input/day02.txt")
	if err != nil {
		log.Fatal(err)
	}

	safeReports := 0
	for i := range lines {
		if isReportSafe(lines[i]) {
			safeReports++
		}
	}
	fmt.Println("Part 1:", safeReports)
}

func isReportSafe(line string) bool {
	intLine := castIntReport(line)
	var intLineReversed []int

	for i := len(intLine) - 1; i >= 0; i-- {
		intLineReversed = append(intLineReversed, intLine[i])
	}

	if !sort.IntsAreSorted(intLine) && !sort.IntsAreSorted(intLineReversed) {
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
