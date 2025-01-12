package main

import (
	"AoC2024/util"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

const PLUS = "+"
const MULTIPLY = "*"
const CONCATENATION = "|"

var legalOperatorsTask01 = []string{PLUS, MULTIPLY}
var legalOperatorsTask02 = []string{PLUS, MULTIPLY, CONCATENATION}

type CalibratorEquations struct {
	checkValue int
	operators  []int
}

func newCalibratorEquations(str string) *CalibratorEquations {
	str = strings.Replace(str, ":", "", -1)
	split := strings.Split(str, " ")

	cv, _ := strconv.Atoi(split[0])
	operators := make([]int, 0)
	for i := 1; i < len(split); i++ {
		o, _ := strconv.Atoi(split[i])
		operators = append(operators, o)
	}

	return &CalibratorEquations{
		checkValue: cv,
		operators:  operators,
	}
}

func (c *CalibratorEquations) IsCheckValueReachable(mathOperatorList []string) bool {
	base := len(mathOperatorList)
	pow := len(c.operators) - 1
	validVariations := int(math.Pow(float64(base), float64(pow)))
	for i := 0; i < validVariations; i++ {
		mutationSeed := util.IntToNewBaseFixedLength(i, base, pow)
		mutatedMathOperators := toMathOperators(mutationSeed)
		if c.checkValue == c.solve(mutatedMathOperators) {
			return true
		}
	}
	return false
}

func (c *CalibratorEquations) solve(mathOperators string) int {
	result := c.operators[0]
	for i := range mathOperators {
		switch string(mathOperators[i]) {
		case PLUS:
			result += c.operators[i+1]
		case MULTIPLY:
			result *= c.operators[i+1]
		case CONCATENATION:
			leftSide := strconv.Itoa(result)
			rightSide := strconv.Itoa(c.operators[i+1])
			tmp, _ := strconv.Atoi(leftSide + rightSide)
			result = tmp
		default:
			panic("not implemented")
		}
	}
	return result
}

func toMathOperators(binaryMutationSeed string) string {
	mutatedOperators := ""
	for i := range binaryMutationSeed {
		switch binaryMutationSeed[i] {
		case '0':
			mutatedOperators += PLUS
		case '1':
			mutatedOperators += MULTIPLY
		case '2':
			mutatedOperators += CONCATENATION
		default:
			panic("not implemented")
		}
	}
	return mutatedOperators
}

func main() {
	lines, err := util.ReadFilePerLine("./input/day07.txt")
	if err != nil {
		log.Fatal(err)
	}

	ces := make([]*CalibratorEquations, 0)
	for i := 0; i < len(lines); i++ {
		ces = append(ces, newCalibratorEquations(lines[i]))
	}

	result := task01(ces)
	fmt.Println(result)
	result = task02(ces)
	fmt.Println(result)
}

func task01(ces []*CalibratorEquations) int {
	result := 0
	for i := range ces {
		if ces[i].IsCheckValueReachable(legalOperatorsTask01) {
			result += ces[i].checkValue
		}
	}
	return result
}

func task02(ces []*CalibratorEquations) int {
	result := 0
	for i := range ces {
		if ces[i].IsCheckValueReachable(legalOperatorsTask02) {
			result += ces[i].checkValue
		}
	}
	return result
}
