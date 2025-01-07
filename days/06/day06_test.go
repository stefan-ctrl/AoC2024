package main

import (
	"AoC2024/days/06/board"
	"AoC2024/util"
	"log"
	"os"
	"reflect"
	"testing"
)

var exampleStartingBoard board.Board
var lines []string
var err error

func setup() {
	lines, err = util.ReadFilePerLine("../../input/day06_example.txt")
	if err != nil {
		log.Fatal(err)
	}
	exampleStartingBoard = board.NewBoard(lines)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func Test_task01(t *testing.T) {
	type args struct {
		b board.Board
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []board.Coordinates
	}{
		{
			name: "AOC_Example",
			args: args{
				b: exampleStartingBoard,
			},
			want:  41,
			want1: []board.Coordinates{{X: 4, Y: 1}, {X: 5, Y: 1}, {X: 6, Y: 1}, {X: 7, Y: 1}, {X: 8, Y: 1}, {X: 4, Y: 2}, {X: 8, Y: 2}, {X: 4, Y: 3}, {X: 8, Y: 3}, {X: 2, Y: 4}, {X: 3, Y: 4}, {X: 4, Y: 4}, {X: 5, Y: 4}, {X: 6, Y: 4}, {X: 8, Y: 4}, {X: 2, Y: 5}, {X: 4, Y: 5}, {X: 6, Y: 5}, {X: 8, Y: 5}, {X: 2, Y: 6}, {X: 3, Y: 6}, {X: 4, Y: 6}, {X: 5, Y: 6}, {X: 6, Y: 6}, {X: 7, Y: 6}, {X: 8, Y: 6}, {X: 1, Y: 7}, {X: 2, Y: 7}, {X: 3, Y: 7}, {X: 4, Y: 7}, {X: 5, Y: 7}, {X: 6, Y: 7}, {X: 7, Y: 7}, {X: 1, Y: 8}, {X: 2, Y: 8}, {X: 3, Y: 8}, {X: 4, Y: 8}, {X: 5, Y: 8}, {X: 6, Y: 8}, {X: 7, Y: 8}, {X: 7, Y: 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := task01(tt.args.b)
			if got != tt.want {
				t.Errorf("task01() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("task01() got1 = %##v,\n want %##v", got1, tt.want1)
			}
			tt.args.b.Print()
		})
	}
}

func Test_task02(t *testing.T) {
	type args struct {
		lines              []string
		visitedCoordinates []board.Coordinates
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "AOC Example",
			args: args{
				lines:              lines,
				visitedCoordinates: []board.Coordinates{{X: 4, Y: 1}, {X: 5, Y: 1}, {X: 6, Y: 1}, {X: 7, Y: 1}, {X: 8, Y: 1}, {X: 4, Y: 2}, {X: 8, Y: 2}, {X: 4, Y: 3}, {X: 8, Y: 3}, {X: 2, Y: 4}, {X: 3, Y: 4}, {X: 4, Y: 4}, {X: 5, Y: 4}, {X: 6, Y: 4}, {X: 8, Y: 4}, {X: 2, Y: 5}, {X: 4, Y: 5}, {X: 6, Y: 5}, {X: 8, Y: 5}, {X: 2, Y: 6}, {X: 3, Y: 6}, {X: 4, Y: 6}, {X: 5, Y: 6}, {X: 6, Y: 6}, {X: 7, Y: 6}, {X: 8, Y: 6}, {X: 1, Y: 7}, {X: 2, Y: 7}, {X: 3, Y: 7}, {X: 4, Y: 7}, {X: 5, Y: 7}, {X: 6, Y: 7}, {X: 7, Y: 7}, {X: 1, Y: 8}, {X: 2, Y: 8}, {X: 3, Y: 8}, {X: 4, Y: 8}, {X: 5, Y: 8}, {X: 6, Y: 8}, {X: 7, Y: 8}, {X: 7, Y: 9}},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := task02(tt.args.lines, tt.args.visitedCoordinates); got != tt.want {
				t.Errorf("task02() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wouldGuardMoveInLoop(t *testing.T) {

	tests := []struct {
		name  string
		lines []string
		want  bool
	}{
		{
			name: "corned",
			lines: []string{
				"....#.....",
				"..>..O....",
				"...#......",
				"....#.....",
				"..........",
			},
			want: true,
		},
		{
			name: "not quiet corned",
			lines: []string{
				"....#.....",
				"..>..O....",
				"..........",
				"....#.....",
				"..........",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := board.NewBoard(tt.lines)
			if got := wouldGuardMoveInLoop(&b); got != tt.want {
				t.Errorf("wouldGuardMoveInLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}
