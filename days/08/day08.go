package main

import (
	"AoC2024/util"
	"fmt"
	"log"
)

const EmptyField = "."
const Antenna = "#"

type column int
type row int

func main() {
	lines, err := util.ReadFilePerLine("./input/day08.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := task01(lines)
	fmt.Println(result)
	result = task02(lines)
	fmt.Println(result)
}

func evaluatePositions(m *util.Matrix[string]) *map[string][]util.Coordinate {
	positions := make(map[string][]util.Coordinate)
	f := func(elem string, row, col int) {
		if elem != EmptyField {
			if positions[elem] == nil {
				positions[elem] = make([]util.Coordinate, 0)
			}
			positions[elem] = append(positions[elem], util.Coordinate{Row: row, Col: col})
		}
	}
	m.ForEach(&f)
	return &positions
}

func task01(lines []string) int {
	matrix := util.StringLinesToMatrix(lines)
	position := evaluatePositions(&matrix)
	for _, v := range *position {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				distanceColumn, distanceRow := calculateDistance(v[i], v[j])
				placeAntenna(&matrix, v[i], v[j], distanceColumn, distanceRow)
			}
		}
	}

	return countAntenna(&matrix)
}

func placeAntenna(m *util.Matrix[string], coordinate util.Coordinate, coordinate2 util.Coordinate, distanceColumn column, distanceRow row) (util.Coordinate, util.Coordinate) {
	antenna1Coordinate := util.Coordinate{Row: coordinate.Row - int(distanceRow), Col: coordinate.Col - int(distanceColumn)}
	antenna2Coordinate := util.Coordinate{Row: coordinate2.Row + int(distanceRow), Col: coordinate2.Col + int(distanceColumn)}

	if !m.IsCoordinateOutOfRange(&antenna1Coordinate) && !m.IsCoordinateOutOfRange(&coordinate) {
		m.SetElement(&antenna1Coordinate, Antenna)
	} else {
		antenna1Coordinate = util.Coordinate{Row: -1, Col: -1}
	}

	if !m.IsCoordinateOutOfRange(&antenna2Coordinate) && !m.IsCoordinateOutOfRange(&coordinate2) {
		m.SetElement(&antenna2Coordinate, Antenna)
	} else {
		antenna2Coordinate = util.Coordinate{Row: -1, Col: -1}
	}
	return antenna1Coordinate, antenna2Coordinate
}

func placeAntennaRepeatedly(m *util.Matrix[string], coordinate util.Coordinate, coordinate2 util.Coordinate) (util.Coordinate, util.Coordinate) {
	distanceColumn, distanceRow := calculateDistance(coordinate, coordinate2)
	for !m.IsCoordinateOutOfRange(&coordinate) || !m.IsCoordinateOutOfRange(&coordinate2) {
		coordinate, coordinate2 = placeAntenna(m, coordinate, coordinate2, distanceColumn, distanceRow)
	}
	return coordinate, coordinate2
}

func countAntenna(m *util.Matrix[string]) int {
	count := 0
	f := func(elem string, _, _ int) {
		if elem == Antenna {
			count++
		}
	}
	m.ForEach(&f)
	return count
}

func countRemainingSenders(m *util.Matrix[string], p *map[string][]util.Coordinate) int {
	count := 0
	f := func(elem string, _, _ int) {
		if elem != Antenna && elem != EmptyField && len((*p)[elem]) > 1 {
			count++
		}
	}
	m.ForEach(&f)
	return count
}

// calculateDistance to - from
func calculateDistance(from, to util.Coordinate) (column, row) {
	return column(to.Col - from.Col), row(to.Row - from.Row)
}

func task02(lines []string) int {
	matrix := util.StringLinesToMatrix(lines)
	position := evaluatePositions(&matrix)
	for _, v := range *position {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				placeAntennaRepeatedly(&matrix, v[i], v[j])
			}
		}
	}

	return countAntenna(&matrix) + countRemainingSenders(&matrix, position)
}
