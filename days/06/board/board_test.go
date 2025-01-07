package board

import (
	"AoC2024/util"
	"log"
	"os"
	"reflect"
	"testing"
)

var exampleStartingBoard Board

func setup() {
	var lines, err = util.ReadFilePerLine("../../../input/day06_example.txt")
	if err != nil {
		log.Fatal(err)
	}
	exampleStartingBoard = NewBoard(lines)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestBoard_IsGuardOnBoard(t *testing.T) {
	tests := []struct {
		name  string
		board Board
		want  bool
	}{
		{
			name:  "Example, Starting Position",
			board: exampleStartingBoard,
			want:  true,
		},
		{
			name: "Starting Position without Guard",
			board: Board{
				guardX: -1,
				guardY: -1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.board.IsGuardOnBoard(); got != tt.want {
				t.Errorf("IsGuardOnBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBoard(t *testing.T) {
	var lines, err = util.ReadFilePerLine("../../../input/day06_example.txt")
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		name string
		want Board
	}{
		{
			name: "AOC Example",
			want: Board{
				matrix: [][]string{
					{".", ".", ".", ".", "#", ".", ".", ".", ".", "."},
					{".", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
					{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
					{".", ".", "#", ".", ".", ".", ".", ".", ".", "."},
					{".", ".", ".", ".", ".", ".", ".", "#", ".", "."},
					{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
					{".", "#", ".", ".", "^", ".", ".", ".", ".", "."},
					{".", ".", ".", ".", ".", ".", ".", ".", "#", "."},
					{"#", ".", ".", ".", ".", ".", ".", ".", ".", "."},
					{".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
				},
				guardX:              4,
				guardY:              6,
				guardDirectionIndex: 0,
				guardDirection:      GuardNorth,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBoard(lines); !reflect.DeepEqual(got, tt.want) && !reflect.DeepEqual(got, exampleStartingBoard) {
				t.Errorf("NewBoard() = %##v, want %##v", got, tt.want)
				t.Errorf("ExampleBoard = %##v, want %##v", got, exampleStartingBoard)
			}
		})
	}
}

func TestBoard_CountGuardVisitedFields(t *testing.T) {
	type fields struct {
		matrix [][]string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		want1  []Coordinates
	}{
		{
			name:  "empty",
			want:  0,
			want1: []Coordinates{},
		},
		{
			name: "2 fields",
			fields: fields{
				matrix: [][]string{
					{".", "X", ".", ".", "#"},
					{".", "X", ".", ".", "X"},
				},
			},
			want: 3,
			want1: []Coordinates{
				{X: 1, Y: 0},
				{X: 1, Y: 1},
				{X: 4, Y: 1},
			},
		},

		{
			name: "aoc_example",
			fields: fields{
				matrix: [][]string{
					{".", ".", ".", ".", "#", ".", ".", ".", ".", "."},
					{".", ".", ".", ".", "X", "X", "X", "X", "X", "#"},
					{".", ".", ".", ".", "X", ".", ".", ".", "X", "."},
					{".", ".", "#", ".", "X", ".", ".", ".", "X", "."},
					{".", ".", "X", "X", "X", "X", "X", "#", "X", "."},
					{".", ".", "X", ".", "X", ".", "X", ".", "X", "."},
					{".", "#", "X", "X", "X", "X", "X", "X", "X", "."},
					{".", "X", "X", "X", "X", "X", "X", "X", "#", "."},
					{"#", "X", "X", "X", "X", "X", "X", "X", ".", "."},
					{".", ".", ".", ".", ".", ".", "#", "X", ".", "."},
				},
			},
			want:  41,
			want1: []Coordinates{{X: 4, Y: 1}, {X: 5, Y: 1}, {X: 6, Y: 1}, {X: 7, Y: 1}, {X: 8, Y: 1}, {X: 4, Y: 2}, {X: 8, Y: 2}, {X: 4, Y: 3}, {X: 8, Y: 3}, {X: 2, Y: 4}, {X: 3, Y: 4}, {X: 4, Y: 4}, {X: 5, Y: 4}, {X: 6, Y: 4}, {X: 8, Y: 4}, {X: 2, Y: 5}, {X: 4, Y: 5}, {X: 6, Y: 5}, {X: 8, Y: 5}, {X: 2, Y: 6}, {X: 3, Y: 6}, {X: 4, Y: 6}, {X: 5, Y: 6}, {X: 6, Y: 6}, {X: 7, Y: 6}, {X: 8, Y: 6}, {X: 1, Y: 7}, {X: 2, Y: 7}, {X: 3, Y: 7}, {X: 4, Y: 7}, {X: 5, Y: 7}, {X: 6, Y: 7}, {X: 7, Y: 7}, {X: 1, Y: 8}, {X: 2, Y: 8}, {X: 3, Y: 8}, {X: 4, Y: 8}, {X: 5, Y: 8}, {X: 6, Y: 8}, {X: 7, Y: 8}, {X: 7, Y: 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				matrix: tt.fields.matrix,
			}
			got, got1 := b.CountGuardVisitedFields()
			if got != tt.want {
				t.Errorf("CountGuardVisitedFields() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CountGuardVisitedFields() got1 = %##v, want %##v", got1, tt.want1)
			}
		})
	}
}

func TestBoard_MoveGuard(t *testing.T) {
	type fields struct {
		matrix [][]string
	}
	type args struct {
		obstacleCheckList []string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   [][]string
	}{
		{
			name: "forward north",
			fields: fields{
				matrix: [][]string{
					{".", "#", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", "^", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
				},
			},
			args: args{
				obstacleCheckList: []string{Obstacle},
			},
			want: [][]string{
				{".", "#", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", "^", "."},
				{".", "X", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
			},
		},
		{
			name: "forward north, corner",
			fields: fields{
				matrix: [][]string{
					{".", "#", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", "#", "."},
					{".", "^", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
				},
			},
			args: args{
				obstacleCheckList: []string{Obstacle},
			},
			want: [][]string{
				{".", "#", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", "#", "."},
				{".", "X", ">"},
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
			},
		},
		{
			name: "forward south",
			fields: fields{
				matrix: [][]string{
					{".", "#", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", "v", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
				},
			},
			args: args{
				obstacleCheckList: []string{Obstacle},
			},
			want: [][]string{
				{".", "#", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", "X", "."},
				{".", "v", "."},
				{".", ".", "."},
				{".", ".", "."},
			},
		},
		{
			name: "forward south, double corner",
			fields: fields{
				matrix: [][]string{
					{".", "#", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
					{"#", "v", "."},
					{".", "O", "."},
					{".", ".", "."},
					{".", ".", "."},
				},
			},
			args: args{
				obstacleCheckList: []string{Obstacle, LoopObstacle},
			},
			want: [][]string{
				{".", "#", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", "^", "."},
				{"#", "X", "."},
				{".", "O", "."},
				{".", ".", "."},
				{".", ".", "."},
			},
		},
		{
			name: "forward east",
			fields: fields{
				matrix: [][]string{
					{".", "#", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ">", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
				},
			},
			args: args{
				obstacleCheckList: []string{Obstacle},
			},
			want: [][]string{
				{".", "#", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", "X", ">"},
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
			},
		},
		{
			name: "forward west",
			fields: fields{
				matrix: [][]string{
					{".", "#", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", "<", "."},
					{".", ".", "."},
					{".", ".", "."},
					{".", ".", "."},
				},
			},
			args: args{
				obstacleCheckList: []string{Obstacle},
			},
			want: [][]string{
				{".", "#", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
				{"<", "X", "."},
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				matrix: tt.fields.matrix,
			}
			b.indexStartingPosition()
			b.MoveGuard(tt.args.obstacleCheckList...)
			if !reflect.DeepEqual(b.matrix, tt.want) {
				t.Errorf("got %v, want %v", b.matrix, tt.want)
			}
		})
	}
}

func TestBoard_isGuardFacingObstacle(t *testing.T) {
	type fields struct {
		matrix [][]string
	}
	type args struct {
		obstacle string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "south",
			fields: fields{
				matrix: [][]string{
					{".", ".", "."},
					{".", "v", "."},
					{".", "#", "."},
				},
			},
			args: args{
				obstacle: Obstacle,
			},
			want: true,
		},
		{
			name: "south free",
			fields: fields{
				matrix: [][]string{
					{".", "#", "."},
					{"#", "v", "#"},
					{".", ".", "."},
				},
			},
			args: args{
				obstacle: Obstacle,
			},
			want: false,
		},
		{
			name: "south edge",
			fields: fields{
				matrix: [][]string{
					{".", ".", "."},
					{".", "v", "."},
				},
			},
			args: args{
				obstacle: Obstacle,
			},
			want: false,
		},
		{
			name: "north",
			fields: fields{
				matrix: [][]string{
					{".", "#", "."},
					{".", "^", "#"},
					{".", ".", "."},
				},
			},
			args: args{
				obstacle: Obstacle,
			},
			want: true,
		},
		{
			name: "north free ",
			fields: fields{
				matrix: [][]string{
					{".", ".", "."},
					{"#", "^", "#"},
					{".", "#", "."},
				},
			},
			args: args{
				obstacle: Obstacle,
			},
			want: false,
		},
		{
			name: "north edge",
			fields: fields{
				matrix: [][]string{
					{".", "^", "#"},
					{".", ".", "."},
				},
			},
			args: args{
				obstacle: Obstacle,
			},
			want: false,
		},
		{
			name: "east",
			fields: fields{
				matrix: [][]string{
					{".", ".", "."},
					{".", ">", "#"},
					{".", ".", "."},
				},
			},
			args: args{
				obstacle: Obstacle,
			},
			want: true,
		},
		{
			name: "east free",
			fields: fields{
				matrix: [][]string{
					{".", "#", "."},
					{"#", ">", "."},
					{".", "#", "."},
				},
			},
			args: args{
				obstacle: Obstacle,
			},
			want: false,
		},
		{
			name: "east edge",
			fields: fields{
				matrix: [][]string{
					{".", "."},
					{".", ">"},
					{".", "."},
				},
			},
			args: args{
				obstacle: Obstacle,
			},
			want: false,
		},
		{
			name: "west",
			fields: fields{
				matrix: [][]string{
					{".", ".", "."},
					{"#", "<", "."},
					{".", ".", "."},
				},
			},
			args: args{
				obstacle: Obstacle,
			},
			want: true,
		},
		{
			name: "west free",
			fields: fields{
				matrix: [][]string{
					{".", "#", "."},
					{".", "<", "#"},
					{".", "#", "."},
				},
			},
			args: args{
				obstacle: Obstacle,
			},
			want: false,
		},
		{
			name: "west edge",
			fields: fields{
				matrix: [][]string{
					{".", "."},
					{"<", "."},
					{".", "."},
				},
			},
			args: args{
				obstacle: Obstacle,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				matrix: tt.fields.matrix,
			}
			b.indexStartingPosition()
			if got := b.isGuardFacingObstacle(tt.args.obstacle); got != tt.want {
				t.Errorf("isGuardFacingObstacle() = %v, want %v", got, tt.want)
			}
		})
	}
}
