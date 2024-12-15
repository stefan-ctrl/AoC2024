package main

import (
	"AoC2024/util"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type PageOrderRule struct {
	Before int
	After  int
}

type Order []int

func (u Order) IsOrderSuccessful(notAllowed map[int][]int) bool {
	for i := range u {
		for j := i + 1; j < len(u); j++ {
			if util.FindInSlice(notAllowed[u[i]], u[j]) {
				return false
			}
		}
	}
	return true
}

func (u Order) CorrectOrder(notAllowed map[int][]int) Order {
	result := u
	for i := range u {
		for j := i + 1; j < len(u); j++ {
			if util.FindInSlice(notAllowed[result[i]], result[j]) {
				tmp := result[i]
				result[i] = result[j]
				result[j] = tmp
				result = result.CorrectOrder(notAllowed)
			}
		}
	}
	return result
}

func main() {
	lines, err := util.ReadFilePerLine("./input/day05.txt")
	if err != nil {
		log.Fatal(err)
	}
	p, u := lineParser(lines)
	result := task01(p, u)
	fmt.Println(result)
	result = task02(p, u)
	fmt.Println(result)
}

func task01(pageOrderRules []PageOrderRule, order []Order) int {
	middlePagesOfTrueUpdates := make([]int, 0)
	notAllowed := inverseRuleset(pageOrderRules)

	for i := range order {
		if order[i].IsOrderSuccessful(notAllowed) {
			middlePagesOfTrueUpdates = append(middlePagesOfTrueUpdates, util.SliceGetMiddleValue((*[]int)(&(order[i]))))
		}
	}

	sum := 0
	for i := range middlePagesOfTrueUpdates {
		sum += middlePagesOfTrueUpdates[i]
	}
	return sum
}

func task02(pageOrderRules []PageOrderRule, order []Order) int {
	middlePagesOfTrueUpdates := make([]int, 0)

	notAllowed := inverseRuleset(pageOrderRules)

	for i := range order {
		if !order[i].IsOrderSuccessful(notAllowed) {
			correctedOrder := order[i].CorrectOrder(notAllowed)
			middlePagesOfTrueUpdates = append(middlePagesOfTrueUpdates, util.SliceGetMiddleValue((*[]int)(&correctedOrder)))
		}
	}

	sum := 0
	for i := range middlePagesOfTrueUpdates {
		sum += middlePagesOfTrueUpdates[i]
	}
	return sum
}

func lineParser(lines []string) ([]PageOrderRule, []Order) {
	pageOrderRules := make([]PageOrderRule, 0)
	updates := make([]Order, 0)

	createPageOrderRules := true
	for i := range lines {
		if lines[i] == "" {
			createPageOrderRules = false
			continue
		}

		if createPageOrderRules {
			pageOrderRules = append(pageOrderRules, pageOrderRule(lines[i]))
		} else {
			updates = append(updates, order(lines[i]))
		}
	}
	return pageOrderRules, updates

}

func order(s string) Order {
	split := strings.Split(s, ",")
	splitAsInt := make([]int, len(split))
	for i := range split {
		var err error
		splitAsInt[i], err = strconv.Atoi(split[i])
		if err != nil {
			panic(err)
		}
	}
	return splitAsInt
}

func pageOrderRule(str string) PageOrderRule {
	split := strings.Split(str, "|")
	before, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}
	after, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}

	return PageOrderRule{
		Before: before,
		After:  after,
	}
}

func inverseRuleset(pageOrderRules []PageOrderRule) map[int][]int {
	notAllowedSuccessionMap := make(map[int][]int)
	for i := range pageOrderRules {
		if notAllowedSuccessionMap[pageOrderRules[i].After] == nil {
			notAllowedSuccessionMap[pageOrderRules[i].After] = make([]int, 0)
		}
		notAllowedSuccessionMap[pageOrderRules[i].After] = append(notAllowedSuccessionMap[pageOrderRules[i].After], pageOrderRules[i].Before)
	}
	return notAllowedSuccessionMap
}
