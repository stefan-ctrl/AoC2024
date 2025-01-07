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

	lines, err := util.ReadFilePerLine("./input/day06.txt")
	if err != nil {
		log.Fatal(err)
	}
	b := board.NewBoard(lines)
	visitedCoordinates = b.RemoveLocationsInSight(visitedCoordinates)
	for i := range visitedCoordinates {
		successfulLoopMarker := false
		b := board.NewBoard(lines)

		coordinate := visitedCoordinates[i]
		if b.IsFree(coordinate.X, coordinate.Y) {
			b.PlaceMarker(coordinate.X, coordinate.Y, board.LoopObstacle)
		} else {
			continue
		}

		eachStep := 1000
		step := 0
		for b.IsGuardOnBoard() {
			b.MoveGuard(board.Obstacle, board.LoopObstacle)
			if b.IsGuardInLoop() {
				successfulLoopMarker = true
				break
			}
			if step%eachStep == 0 {
				b.Print()
				println("------")
				println("")
			}
			step++
		}

		if successfulLoopMarker {
			fmt.Printf("Possible Location %d, %d\n", coordinate.X, coordinate.Y)
			loopPossibilityCounter++
		} else {
			fmt.Printf("Iteration not successful %d\n", i)
		}

	}

	return loopPossibilityCounter
}
