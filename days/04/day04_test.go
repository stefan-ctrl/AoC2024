package main

import (
	"reflect"
	"testing"
)

func Test_toMatrix(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want [][]rune
	}{
		{
			name: "AoC Example",
			args: args{
				lines: []string{
					"MMMSXXMASM",
					"MSAMXMSMSA",
					"AMXSXMAAMM",
					"MSAMASMSMX",
					"XMASAMXAMM",
					"XXAMMXXAMA",
					"SMSMSASXSS",
					"SAXAMASAAA",
					"MAMMMXMMMM",
					"MXMXAXMASX",
				},
			},
			want: [][]rune{
				{M, M, M, S, X, X, M, A, S, M},
				{M, S, A, M, X, M, S, M, S, A},
				{A, M, X, S, X, M, A, A, M, M},
				{M, S, A, M, A, S, M, S, M, X},
				{X, M, A, S, A, M, X, A, M, M},
				{X, X, A, M, M, X, X, A, M, A},
				{S, M, S, M, S, A, S, X, S, S},
				{S, A, X, A, M, A, S, A, A, A},
				{M, A, M, M, M, X, M, M, M, M},
				{M, X, M, X, A, X, M, A, S, X},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toMatrix(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_outOfRange(t *testing.T) {
	type args struct {
		row     int
		col     int
		col_len int
		row_len int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "within range",
			args: args{
				row:     0,
				col:     0,
				col_len: 1,
				row_len: 1,
			},
			want: false,
		},
		{
			name: "row, col larger by 0",
			args: args{
				row:     1,
				col:     1,
				col_len: 1,
				row_len: 1,
			},
			want: true,
		},
		{
			name: "col larger by 1",
			args: args{
				row:     0,
				col:     2,
				col_len: 1,
				row_len: 1,
			},
			want: true,
		},
		{
			name: "row larger by 1",
			args: args{
				row:     2,
				col:     0,
				col_len: 1,
				row_len: 1,
			},
			want: true,
		},
		{
			name: "row, col smaller by -1",
			args: args{
				row:     -1,
				col:     -1,
				col_len: 1,
				row_len: 1,
			},
			want: true,
		},
		{
			name: "col smaller by -1",
			args: args{
				row:     0,
				col:     -1,
				col_len: 1,
				row_len: 1,
			},
			want: true,
		},
		{
			name: "row smaller by -1",
			args: args{
				row:     -1,
				col:     0,
				col_len: 1,
				row_len: 1,
			},
			want: true,
		},
		{
			name: "matrix size 0,0",
			args: args{
				row:     0,
				col:     0,
				col_len: 0,
				row_len: 0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := outOfRange(tt.args.row, tt.args.col, tt.args.col_len, tt.args.row_len); got != tt.want {
				t.Errorf("outOfRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_foundHorizontal(t *testing.T) {
	matrix := [][]rune{
		{M, M, M, S, X, X, M, A, S, M},
		{M, S, A, M, X, M, S, M, S, A},
		{A, M, X, S, X, M, A, A, M, M},
		{M, S, A, M, A, S, M, S, M, X},
		{X, M, A, S, A, M, X, A, M, M},
		{X, X, A, M, M, X, X, A, M, A},
		{S, M, S, M, S, A, S, X, S, S},
		{S, A, X, A, M, A, S, A, A, A},
		{M, A, M, M, M, X, M, M, M, M},
		{M, X, M, X, A, X, M, A, S, X},
	}
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no xmas at 0,3",
			args: args{
				row: 0,
				col: 3,
			},
			want: 0,
		},
		{
			name: "no xmas at 0,4",
			args: args{
				row: 0,
				col: 4,
			},
			want: 0,
		},
		{
			name: "xmas at 0,5",
			args: args{
				row: 0,
				col: 5,
			},
			want: 1,
		},
		{
			name: "xmas (backwards) at 1,4",
			args: args{
				row: 1,
				col: 4,
			},
			want: 1,
		},
		{
			name: "xmas at 9,5",
			args: args{
				row: 9,
				col: 5,
			},
			want: 1,
		},
		{
			name: "xmas (backwards) at 4,6",
			args: args{
				row: 4,
				col: 6,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := foundHorizontal(matrix, tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("foundHorizontal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_foundXMAS(t *testing.T) {
	matrix := [][]rune{
		{M, M, M, S, X, X, M, A, S, M},
		{M, S, A, M, X, M, S, M, S, A},
		{A, M, X, S, X, M, A, A, M, M},
		{M, S, A, M, A, S, M, S, M, X},
		{X, M, A, S, A, M, X, A, M, M},
		{X, X, A, M, M, X, X, A, M, A},
		{S, M, S, M, S, A, S, X, S, S},
		{S, A, X, A, M, A, S, A, A, A},
		{M, A, M, M, M, X, M, M, M, M},
		{M, X, M, X, A, X, M, A, S, X},
	}

	type args struct {
		row int
		col int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no xmas at 0,3",
			args: args{
				row: 0,
				col: 3,
			},
			want: 0,
		},
		{
			name: "forward at 9,5",
			args: args{
				row: 9,
				col: 5,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := foundXMAS(matrix, tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("foundXMAS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_task01(t *testing.T) {
	matrix := [][]rune{
		{M, M, M, S, X, X, M, A, S, M},
		{M, S, A, M, X, M, S, M, S, A},
		{A, M, X, S, X, M, A, A, M, M},
		{M, S, A, M, A, S, M, S, M, X},
		{X, M, A, S, A, M, X, A, M, M},
		{X, X, A, M, M, X, X, A, M, A},
		{S, M, S, M, S, A, S, X, S, S},
		{S, A, X, A, M, A, S, A, A, A},
		{M, A, M, M, M, X, M, M, M, M},
		{M, X, M, X, A, X, M, A, S, X},
	}

	tests := []struct {
		name string
		want int
	}{
		{
			name: "AoC Example",
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := task01(matrix); got != tt.want {
				t.Errorf("task01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_foundVertical(t *testing.T) {
	matrix := [][]rune{
		{M, M, M, S, X, X, M, A, S, M},
		{M, S, A, M, X, M, S, M, S, A},
		{A, M, X, S, X, M, A, A, M, M},
		{M, S, A, M, A, S, M, S, M, X},
		{X, M, A, S, A, M, X, A, M, M},
		{X, X, A, M, M, X, X, A, M, A},
		{S, M, S, M, S, A, S, X, S, S},
		{S, A, X, A, M, A, S, A, A, A},
		{M, A, M, M, M, X, M, M, M, M},
		{M, X, M, X, A, X, M, A, S, X},
	}
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no xmas at 0,0",
			args: args{
				row: 0,
				col: 0,
			},
			want: 0,
		},
		{
			name: "xmas at 4,6",
			args: args{
				row: 4,
				col: 6,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := foundVertical(matrix, tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("foundVertical() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_task02(t *testing.T) {
	matrix := [][]rune{
		{M, M, M, S, X, X, M, A, S, M},
		{M, S, A, M, X, M, S, M, S, A},
		{A, M, X, S, X, M, A, A, M, M},
		{M, S, A, M, A, S, M, S, M, X},
		{X, M, A, S, A, M, X, A, M, M},
		{X, X, A, M, M, X, X, A, M, A},
		{S, M, S, M, S, A, S, X, S, S},
		{S, A, X, A, M, A, S, A, A, A},
		{M, A, M, M, M, X, M, M, M, M},
		{M, X, M, X, A, X, M, A, S, X},
	}
	tests := []struct {
		name string
		want int
	}{
		{
			name: "AdventOfCode Example",
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := task02(matrix); got != tt.want {
				t.Errorf("task02() = %v, want %v", got, tt.want)
			}
		})
	}
}
