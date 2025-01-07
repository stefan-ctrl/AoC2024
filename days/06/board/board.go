package board

import (
	"fmt"
	"log"
	"slices"
	"strings"
)

const (
	Obstacle     = "#"
	GuardNorth   = "^"
	GuardEast    = ">"
	GuardSouth   = "v"
	GuardWest    = "<"
	FREE         = "."
	GuardVisited = "X"
	LoopObstacle = "O"
)

var RightTurnOrder = []string{GuardNorth, GuardEast, GuardSouth, GuardWest}

func NewBoard(lines []string) Board {
	matrix := make([][]string, len(lines))
	for i, line := range lines {
		matrix[i] = strings.Split(line, "")
	}
	b := Board{matrix: matrix}
	b.indexStartingPosition()
	b.trace = make(map[string]bool)
	return b
}

type Board struct {
	matrix                 [][]string
	guardX                 int
	guardY                 int
	guardDirection         string
	guardDirectionIndex    int
	startingGuardX         int
	startingGuardY         int
	startingGuardDirection string
	trace                  map[string]bool
}

type Coordinates struct {
	X, Y int
}
type CoordinatesWithDirection struct {
	Coordinates
	direction string
}

func (b *Board) IsGuardOnBoard() bool {
	return b.guardY < len(b.matrix) && b.guardY >= 0 && b.guardX < len(b.matrix[0]) && b.guardX >= 0
}

func (b *Board) CountGuardVisitedFields() (int, []Coordinates) {
	counter := 0
	uniqueVisitedCoordinates := make([]Coordinates, 0)
	for y := range b.matrix {
		for x := range b.matrix[y] {
			if b.matrix[y][x] == GuardVisited {
				uniqueVisitedCoordinates = append(uniqueVisitedCoordinates, Coordinates{x, y})
				counter++
			}
		}
	}
	return counter, uniqueVisitedCoordinates
}

func (b *Board) MoveGuard(obstacleCheckList ...string) {
	oldX := b.guardX
	oldY := b.guardY

	for b.isGuardFacingObstacle(obstacleCheckList...) {
		b.registerPosition()
		b.turnGuard()
	}

	b.registerPosition()
	b.stepGuardForward()

	if oldX != b.guardX || oldY != b.guardY {
		b.PlaceMarker(oldX, oldY, GuardVisited)
	}
}

func (b *Board) registerPosition() {
	str := b.generateCurrentPositionIdentifier()
	if b.trace != nil {
		b.trace[str] = true
	}
}

func (b *Board) PlaceMarker(x, y int, char string) {
	b.matrix[y][x] = char
}

func (b *Board) isGuardFacingObstacle(obstacle ...string) bool {
	yCheck := 0
	xCheck := 0
	switch b.guardDirection {
	case GuardNorth:
		if b.guardY == 0 {
			return false
		}
		yCheck = b.guardY - 1
		xCheck = b.guardX
	case GuardEast:
		if b.guardX == len(b.matrix[0])-1 {
			return false
		}
		yCheck = b.guardY
		xCheck = b.guardX + 1
	case GuardSouth:
		if b.guardY == len(b.matrix)-1 {
			return false
		}
		yCheck = b.guardY + 1
		xCheck = b.guardX
	case GuardWest:
		if b.guardX == 0 {
			return false
		}
		yCheck = b.guardY
		xCheck = b.guardX - 1
	default:
		log.Fatalf("Invalid guard pointing %s at X: %d and Y: %d", b.guardDirection, b.guardX, b.guardY)
	}
	return slices.Contains(obstacle, b.matrix[yCheck][xCheck])
}

func (b *Board) turnGuard() {
	b.guardDirectionIndex = (b.guardDirectionIndex + 1) % len(RightTurnOrder)
	b.guardDirection = RightTurnOrder[b.guardDirectionIndex]
}

func (b *Board) Print() {
	for y := range b.matrix {
		for x := range b.matrix[y] {
			fmt.Print(b.matrix[y][x])
		}
		fmt.Println()
	}
}

func (b *Board) indexStartingPosition() {
	found := false
	for y := range b.matrix {
		for x := range b.matrix[y] {
			if b.matrix[y][x] != FREE && b.matrix[y][x] != Obstacle {
				b.guardX = x
				b.guardY = y
				b.guardDirection = b.matrix[y][x]
				found = true
			}
			if found {
				break
			}
		}
	}
	b.guardDirectionIndex = slices.Index(RightTurnOrder, b.guardDirection)
	b.startingGuardDirection = b.guardDirection
	b.startingGuardX = b.guardX
	b.startingGuardY = b.guardY
}

func (b *Board) stepGuardForward() {
	b.matrix[b.guardY][b.guardX] = FREE

	switch b.guardDirection {
	case GuardNorth:
		b.guardY -= 1
	case GuardEast:
		b.guardX += 1
	case GuardSouth:
		b.guardY += 1
	case GuardWest:
		b.guardX -= 1
	default:
		log.Fatalf("Invalid guard pointing %s at X: %d and Y: %d", b.guardDirection, b.guardX, b.guardY)
	}

	if b.guardY >= 0 && b.guardY <= len(b.matrix)-1 &&
		b.guardX >= 0 && b.guardX <= len(b.matrix[0])-1 {
		b.matrix[b.guardY][b.guardX] = b.guardDirection
	}
}

func (b *Board) generateCurrentPositionIdentifier() string {
	c := CoordinatesWithDirection{
		Coordinates: Coordinates{
			X: b.guardX,
			Y: b.guardY,
		},
		direction: b.guardDirection,
	}
	str := fmt.Sprintf("%v", c)
	return str
}

func (b *Board) IsFree(x int, y int) bool {
	return b.matrix[y][x] == FREE
}

func (b *Board) IsGuardInLoop() bool {
	str := b.generateCurrentPositionIdentifier()
	_, ok := b.trace[str]
	return ok
}

func (b *Board) RemoveLocationsInSight(coordinates []Coordinates) []Coordinates {
	possibilities := make([]Coordinates, 0)

	for i := range coordinates {
		switch b.startingGuardDirection {
		case GuardNorth:
			if !(coordinates[i].Y < b.guardY && coordinates[i].X == b.guardX) {
				possibilities = append(possibilities, coordinates[i])
			}
		case GuardSouth:
			if !(coordinates[i].Y > b.guardY && coordinates[i].X == b.guardX) {
				possibilities = append(possibilities, coordinates[i])
			}
		case GuardWest:
			if !(coordinates[i].Y == b.guardY && coordinates[i].X < b.guardX) {
				possibilities = append(possibilities, coordinates[i])
			}
		case GuardEast:
			if !(coordinates[i].Y == b.guardY && coordinates[i].X > b.guardX) {
				possibilities = append(possibilities, coordinates[i])
			}
		}
	}
	return possibilities
}
