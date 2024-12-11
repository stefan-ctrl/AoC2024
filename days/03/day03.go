package main

import (
	"AoC2024/util"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const REGEX_01 = "mul\\((\\d+),(\\d+)\\)"
const REGEX_02 = "mul\\((\\d+),(\\d+)\\)|don't\\(\\)|do\\(\\)"

func main() {
	input, err := util.ReadFile("./input/day03.txt")
	if err != nil {
		log.Fatal(err)
	}

	task01(input)
	task02(input)
}

func task01(input string) {
	matches, err := util.RegexAllMatch(input, REGEX_01)
	if err != nil {
		panic(err)
	}

	sum := 0
	for i := range matches {
		a, err := strconv.Atoi(matches[i][1])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(matches[i][2])
		if err != nil {
			panic(err)
		}
		product := a * b
		sum += product
	}
	fmt.Println(sum)
}

func task02(input string) {
	matches, err := util.RegexAllMatch(input, REGEX_02)
	if err != nil {
		panic(err)
	}

	do := true
	sum := 0
	for i := range matches {

		if strings.Contains(matches[i][0], "do()") {
			do = true
			continue
		} else if strings.Contains(matches[i][0], "don't()") {
			do = false
			continue
		}

		if !do {
			continue
		}

		a, err := strconv.Atoi(matches[i][1])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(matches[i][2])
		if err != nil {
			panic(err)
		}
		product := a * b
		sum += product
	}
	fmt.Println(sum)
}
