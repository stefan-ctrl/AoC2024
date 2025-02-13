package main

import (
	"reflect"
	"testing"
)

func Test_toFileBlocks(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want []FileId
	}{
		{
			name: "AoC Example",
			args: args{
				line: "2333133121414131402",
			},
			want: []FileId{
				0, 0,
				9, 9,
				8,
				1, 1, 1,
				8, 8, 8,
				2,
				7, 7, 7,
				3, 3, 3,
				6,
				4, 4,
				6,
				5, 5, 5, 5,
				6,
				6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toFileBlocks(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toFileBlocks() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}

func Test_task01(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "AoC Example",
			args: args{
				line: "2333133121414131402",
			},
			want: 1928,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := task01(tt.args.line); got != tt.want {
				t.Errorf("task01() = %v, want %v", got, tt.want)
			}
		})
	}
}
