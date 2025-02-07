package util

import "C"
import (
	"fmt"
	"strings"
)

// Matrix top left corner as center of origin
type Matrix[T any] [][]T

type Coordinate struct {
	Row, Col int
}

func StringLinesToMatrix(lines []string) Matrix[string] {
	matrix := make([][]string, len(lines))
	for i, line := range lines {
		matrix[i] = strings.Split(line, "")
	}
	return matrix
}

func (m *Matrix[T]) GetElement(c *Coordinate) T {
	return (*m)[c.Row][c.Col]
}

func (m *Matrix[T]) SetElement(c *Coordinate, value T) {
	(*m)[c.Row][c.Col] = value
}

func (m *Matrix[T]) IsCoordinateOutOfRange(c *Coordinate) bool {
	return c.Row < 0 || c.Col < 0 || c.Col >= m.LenCol() || c.Row >= m.LenRow()
}

func (m *Matrix[T]) LenRow() int {
	return len(*m)
}

func (m *Matrix[T]) LenCol() int {
	return len((*m)[0])
}

func (m *Matrix[T]) ForEach(action *func(T, int, int)) {
	for col := range *m {
		for row := range (*m)[col] {
			(*action)((*m)[row][col], row, col)
		}
	}
}

func (m *Matrix[T]) Print() {
	for col := range *m {
		for row := range (*m)[col] {
			fmt.Print((*m)[col][row])
		}
		fmt.Println()
	}
}
