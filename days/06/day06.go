package main

import (
	"AoC2024/days/06/board"
	"AoC2024/util"
	"fmt"
	"log"
)

func main() {
	lines, err := util.ReadFilePerLine("./input/day06.txt")
	if err != nil {
		log.Fatal(err)
	}
	b := board.NewBoard(lines)

	result, coordinates := task01(b)
	fmt.Println(result)
	result = task02(lines, coordinates)
	fmt.Println(result)
}

func task01(b board.Board) (int, []board.Coordinates) {
	for b.IsGuardOnBoard() {
		b.MoveGuard(board.Obstacle)
	}
	result, coords := b.CountGuardVisitedFields()
	return result, coords
}

func task02(lines []string, visitedCoordinates []board.Coordinates) int {
	loopPossibilityCounter := 0

	for i := range visitedCoordinates {
		b := varyBoard(lines, visitedCoordinates[i])

		successfulLoopMarker := wouldGuardMoveInLoop(&b)

		if successfulLoopMarker {
			loopPossibilityCounter++
		}

	}

	return loopPossibilityCounter
}

func wouldGuardMoveInLoop(b *board.Board) bool {
	for b.IsGuardOnBoard() {
		b.MoveGuard(board.Obstacle, board.LoopObstacle)
		if b.IsGuardInLoop() {
			return true
		}
	}
	return false
}

func varyBoard(lines []string, coordinate board.Coordinates) board.Board {
	b := board.NewBoard(lines)
	b.PlaceMarker(coordinate.X, coordinate.Y, board.LoopObstacle)
	return b
}
