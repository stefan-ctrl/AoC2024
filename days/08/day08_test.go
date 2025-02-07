package main

import (
	"AoC2024/util"
	"reflect"
	"testing"
)

var lines, err = util.ReadFilePerLine("../../input/day08_example.txt")

func Test_evaluatePositions(t *testing.T) {

	tests := []struct {
		name string
		want *map[string][]util.Coordinate
	}{
		{
			name: "AOC Example",
			want: &map[string][]util.Coordinate{
				"0": {
					{Row: 4, Col: 4},
					{Row: 2, Col: 5},
					{Row: 3, Col: 7},
					{Row: 1, Col: 8},
				},
				"A": {
					{Row: 5, Col: 6},
					{Row: 8, Col: 8},
					{Row: 9, Col: 9},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := util.StringLinesToMatrix(lines)
			if got := evaluatePositions(&m); !reflect.DeepEqual(*got, *tt.want) {
				t.Errorf("evaluatePositions() = %v, wantColDif %v", *got, *tt.want)
			}
		})
	}
}

func Test_calculateDistance(t *testing.T) {
	type args struct {
		from util.Coordinate
		to   util.Coordinate
	}
	tests := []struct {
		name       string
		args       args
		wantColDif column
		wantRowDif row
	}{
		{
			name: "AOC_Example: Next each other",
			args: args{
				from: util.Coordinate{Row: 8, Col: 8},
				to:   util.Coordinate{Row: 9, Col: 9},
			},
			wantColDif: 1,
			wantRowDif: 1,
		},
		{
			name: "AOC_Example: Next each other, flipped",
			args: args{
				from: util.Coordinate{Row: 9, Col: 9},
				to:   util.Coordinate{Row: 8, Col: 8},
			},
			wantColDif: -1,
			wantRowDif: -1,
		},
		{
			name: "AOC_Example: Random Distance",
			args: args{
				from: util.Coordinate{Row: 4, Col: 4},
				to:   util.Coordinate{Row: 2, Col: 5},
			},
			wantColDif: 1,
			wantRowDif: -2,
		},
		{
			name: "AOC_Example: Sam Row",
			args: args{
				from: util.Coordinate{Row: 4, Col: 4},
				to:   util.Coordinate{Row: 4, Col: 6},
			},
			wantColDif: 2,
			wantRowDif: 0,
		},
		{
			name: "AOC_Example: Sam Col",
			args: args{
				from: util.Coordinate{Row: 0, Col: 6},
				to:   util.Coordinate{Row: 5, Col: 6},
			},
			wantColDif: 0,
			wantRowDif: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			difCol, difRow := calculateDistance(tt.args.from, tt.args.to)
			if difCol != tt.wantColDif {
				t.Errorf("calculateDistance() difCol = %v, wantColDif %v", difCol, tt.wantColDif)
			}
			if difRow != tt.wantRowDif {
				t.Errorf("calculateDistance() difRow = %v, wantRowDif %v", difRow, tt.wantRowDif)
			}
		})
	}
}

func Test_placeAntenna(t *testing.T) {
	type args struct {
		coordinate  util.Coordinate
		coordinate2 util.Coordinate
	}
	tests := []struct {
		name  string
		args  args
		want  util.Coordinate
		want1 util.Coordinate
	}{
		{
			name: "AOC_Example: Next each other",
			args: args{
				coordinate:  util.Coordinate{Row: 8, Col: 8},
				coordinate2: util.Coordinate{Row: 9, Col: 9},
			},
			want:  util.Coordinate{Row: 7, Col: 7},
			want1: util.Coordinate{Row: 10, Col: 10},
		},
		{
			name: "AOC_Example: Next each other, flipped",
			args: args{
				coordinate:  util.Coordinate{Row: 9, Col: 9},
				coordinate2: util.Coordinate{Row: 8, Col: 8},
			},
			want:  util.Coordinate{Row: 10, Col: 10},
			want1: util.Coordinate{Row: 7, Col: 7},
		},
		{
			name: "AOC_Example: Respect Out-Of-Bounce",
			args: args{
				coordinate:  util.Coordinate{Row: 5, Col: 6},
				coordinate2: util.Coordinate{Row: 9, Col: 9},
			},
			want:  util.Coordinate{Row: 1, Col: 3},
			want1: util.Coordinate{Row: -1, Col: -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := util.StringLinesToMatrix(lines)
			got, got1 := placeAntenna(&m, tt.args.coordinate, tt.args.coordinate2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("placeAntenna() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("placeAntenna() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_task01(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "AOC Example",
			args: args{
				lines: lines,
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := task01(tt.args.lines); got != tt.want {
				t.Errorf("task01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_task02(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "AOC Example",
			args: args{
				lines: lines,
			},
			want: 34,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := task02(tt.args.lines); got != tt.want {
				t.Errorf("task01() = %v, want %v", got, tt.want)
			}
		})
	}
}
