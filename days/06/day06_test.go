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

func setup() {
	var lines, err = util.ReadFilePerLine("../../input/day06_example.txt")
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
			want1: []board.Coordinates{board.Coordinates{X: 4, Y: 1}, board.Coordinates{X: 5, Y: 1}, board.Coordinates{X: 6, Y: 1}, board.Coordinates{X: 7, Y: 1}, board.Coordinates{X: 8, Y: 1}, board.Coordinates{X: 4, Y: 2}, board.Coordinates{X: 8, Y: 2}, board.Coordinates{X: 4, Y: 3}, board.Coordinates{X: 8, Y: 3}, board.Coordinates{X: 2, Y: 4}, board.Coordinates{X: 3, Y: 4}, board.Coordinates{X: 4, Y: 4}, board.Coordinates{X: 5, Y: 4}, board.Coordinates{X: 6, Y: 4}, board.Coordinates{X: 8, Y: 4}, board.Coordinates{X: 2, Y: 5}, board.Coordinates{X: 4, Y: 5}, board.Coordinates{X: 6, Y: 5}, board.Coordinates{X: 8, Y: 5}, board.Coordinates{X: 2, Y: 6}, board.Coordinates{X: 3, Y: 6}, board.Coordinates{X: 4, Y: 6}, board.Coordinates{X: 5, Y: 6}, board.Coordinates{X: 6, Y: 6}, board.Coordinates{X: 7, Y: 6}, board.Coordinates{X: 8, Y: 6}, board.Coordinates{X: 1, Y: 7}, board.Coordinates{X: 2, Y: 7}, board.Coordinates{X: 3, Y: 7}, board.Coordinates{X: 4, Y: 7}, board.Coordinates{X: 5, Y: 7}, board.Coordinates{X: 6, Y: 7}, board.Coordinates{X: 7, Y: 7}, board.Coordinates{X: 1, Y: 8}, board.Coordinates{X: 2, Y: 8}, board.Coordinates{X: 3, Y: 8}, board.Coordinates{X: 4, Y: 8}, board.Coordinates{X: 5, Y: 8}, board.Coordinates{X: 6, Y: 8}, board.Coordinates{X: 7, Y: 8}, board.Coordinates{X: 7, Y: 9}},
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
