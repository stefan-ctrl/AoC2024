package main

import (
	"AoC2024/util"
	"fmt"
	"log"
)

const X = 'X'
const M = 'M'
const A = 'A'
const S = 'S'

func main() {
	lines, err := util.ReadFilePerLine("./input/day04.txt")
	if err != nil {
		log.Fatal(err)
	}
	matrix := toMatrix(lines)
	result := task01(matrix)
	fmt.Println(result)
	result = task02(matrix)
	fmt.Println(result)
}

func task01(matrix [][]rune) int {
	xmasCounter := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == X {
				xmasCounter += foundXMAS(matrix, i, j)
			}
		}
	}
	return xmasCounter
}

func task02(matrix [][]rune) int {
	xmasCounter := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == A {
				xmasCounter += foundCrossMAS(matrix, i, j)
			}
		}
	}
	return xmasCounter
}

func toMatrix(lines []string) [][]rune {
	matrix := make([][]rune, len(lines))
	for i := range matrix {
		matrix[i] = make([]rune, len(lines[i]))
	}

	for i := range lines {
		for j := range lines[i] {
			matrix[i][j] = rune(lines[i][j])
		}
	}
	return matrix
}

func foundXMAS(matrix [][]rune, row, col int) int {
	times := 0
	times += foundHorizontal(matrix, row, col)
	times += foundVertical(matrix, row, col)
	times += foundDiagonal(matrix, row, col)

	return times
}

func foundCrossMAS(matrix [][]rune, row, col int) int {
	times := 0
	times += foundDiagonalCrossMAS(matrix, row, col)

	return times
}

func foundHorizontal(matrix [][]rune, row, col int) int {
	timesFound := 0

	//forward
	found := checkNeighbor(matrix, row, col, 1, 0)
	if found {
		timesFound++
	}
	//backward
	found = checkNeighbor(matrix, row, col, -1, 0)
	if found {
		timesFound++
	}

	return timesFound
}

func foundVertical(matrix [][]rune, row, col int) int {
	timesFound := 0

	//forward
	found := checkNeighbor(matrix, row, col, 0, 1)
	if found {
		timesFound++
	}
	//backward
	found = checkNeighbor(matrix, row, col, 0, -1)
	if found {
		timesFound++
	}

	return timesFound
}

func foundDiagonal(matrix [][]rune, row, col int) int {
	timesFound := 0

	//down-right
	found := checkNeighbor(matrix, row, col, 1, 1)
	if found {
		timesFound++
	}
	//down-left
	found = checkNeighbor(matrix, row, col, 1, -1)
	if found {
		timesFound++
	}
	//up-right
	found = checkNeighbor(matrix, row, col, -1, 1)
	if found {
		timesFound++
	}
	//up-left
	found = checkNeighbor(matrix, row, col, -1, -1)
	if found {
		timesFound++
	}

	return timesFound
}

func foundDiagonalCrossMAS(matrix [][]rune, row, col int) int {
	combinations := []struct {
		cross1M, cross2M, cross1S, cross2S struct {
			row int
			col int
		}
	}{
		{
			cross1M: struct {
				row int
				col int
			}{row: 1, col: 1},
			cross1S: struct {
				row int
				col int
			}{row: -1, col: -1},
			cross2M: struct {
				row int
				col int
			}{row: 1, col: -1},
			cross2S: struct {
				row int
				col int
			}{row: -1, col: 1},
		},
		{
			cross1M: struct {
				row int
				col int
			}{row: 1, col: 1},
			cross1S: struct {
				row int
				col int
			}{row: -1, col: -1},
			cross2M: struct {
				row int
				col int
			}{row: -1, col: 1},
			cross2S: struct {
				row int
				col int
			}{row: 1, col: -1},
		},
		{
			cross1M: struct {
				row int
				col int
			}{row: -1, col: -1},
			cross1S: struct {
				row int
				col int
			}{row: 1, col: 1},
			cross2M: struct {
				row int
				col int
			}{row: 1, col: -1},
			cross2S: struct {
				row int
				col int
			}{row: -1, col: 1},
		},
		{
			cross1M: struct {
				row int
				col int
			}{row: -1, col: -1},
			cross1S: struct {
				row int
				col int
			}{row: 1, col: 1},
			cross2M: struct {
				row int
				col int
			}{row: -1, col: 1},
			cross2S: struct {
				row int
				col int
			}{row: 1, col: -1},
		},
	}

	for _, combination := range combinations {
		if checkDiagonalNeighborForRune(matrix, row, col, combination.cross1M.col, combination.cross1M.row, M) &&
			checkDiagonalNeighborForRune(matrix, row, col, combination.cross1S.col, combination.cross1S.row, S) &&
			checkDiagonalNeighborForRune(matrix, row, col, combination.cross2M.col, combination.cross2M.row, M) &&
			checkDiagonalNeighborForRune(matrix, row, col, combination.cross2S.col, combination.cross2S.row, S) {
			return 1
		}
	}

	return 0
}

func checkDiagonalNeighborForRune(matrix [][]rune, row int, col int, coefficientColOffset int, coefficientRowOffset int, r rune) bool {
	numberOfRows := len(matrix)
	numberOfCols := len(matrix[0])

	colToCheck := col + coefficientColOffset
	rowToCheck := row + coefficientRowOffset
	if !outOfRange(rowToCheck, colToCheck, numberOfCols, numberOfRows) {
		return matrix[rowToCheck][colToCheck] == r
	}
	return false
}

func checkNeighbor(matrix [][]rune, row int, col int, coefficientColOffset int, coefficientRowOffset int) bool {
	buffer := string(X)
	MAS := []rune{M, A, S}
	numberOfRows := len(matrix)
	numberOfCols := len(matrix[0])

	for i := range MAS {
		colToCheck := col + (coefficientColOffset * (i + 1))
		rowToCheck := row + (coefficientRowOffset * (i + 1))
		if !outOfRange(rowToCheck, colToCheck, numberOfCols, numberOfRows) {
			if matrix[rowToCheck][colToCheck] == MAS[i] {
				buffer = buffer + string(MAS[i])
			}
		}
	}
	return buffer == "XMAS"
}

func outOfRange(row, col, colLen, rowLen int) bool {
	return col < 0 || row < 0 || col >= colLen || row >= rowLen
}
