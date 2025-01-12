package main

import (
	"AoC2024/util"
	"log"
	"reflect"
	"testing"
)

func Test_newCalibratorEquations(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want *CalibratorEquations
	}{
		{
			name: "empty",
			want: &CalibratorEquations{
				checkValue: 0,
				operators:  make([]int, 0),
			},
		},
		{
			name: "AOC First Line",
			args: args{
				str: "190: 10 19",
			},
			want: &CalibratorEquations{
				checkValue: 190,
				operators:  []int{10, 19},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newCalibratorEquations(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newCalibratorEquations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toMathOperators(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "00",
			want: PLUS + PLUS,
		},
		{
			name: "010110",
			want: PLUS + MULTIPLY + PLUS + MULTIPLY + MULTIPLY + PLUS,
		},
		{
			name: "020112",
			want: PLUS + CONCATENATION + PLUS + MULTIPLY + MULTIPLY + CONCATENATION,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toMathOperators(tt.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toMathOperators() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalibratorEquations_solve(t *testing.T) {
	type fields struct {
		checkValue int
		operators  []int
	}
	type args struct {
		mathOperators string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "AOC Task 1 First Line",
			fields: fields{
				checkValue: 190,
				operators:  []int{10, 19},
			},
			args: args{
				mathOperators: MULTIPLY,
			},
			want: 190,
		},
		{
			name: "AOC Task 1 First Line with plus",
			fields: fields{
				checkValue: 190,
				operators:  []int{10, 19},
			},
			args: args{
				mathOperators: PLUS,
			},
			want: 29,
		},
		{
			name: "AOC Task 2 First Line with CONCATENATION",
			fields: fields{
				checkValue: 190,
				operators:  []int{10, 19},
			},
			args: args{
				mathOperators: CONCATENATION,
			},
			want: 1019,
		},
		{
			name: "AOC Task 1  Example: 292",
			fields: fields{
				checkValue: 292,
				operators:  []int{11, 6, 16, 20},
			},
			args: args{
				mathOperators: PLUS + MULTIPLY + PLUS,
			},
			want: 292,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CalibratorEquations{
				checkValue: tt.fields.checkValue,
				operators:  tt.fields.operators,
			}
			if got := c.solve(tt.args.mathOperators); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalibratorEquations_IsCheckValueReachable(t *testing.T) {
	type fields struct {
		checkValue int
		operators  []int
	}
	tests := []struct {
		name             string
		fields           fields
		mathOperatorList []string
		want             bool
	}{
		{
			name: "AOC task 01 Reachable Example",
			fields: fields{
				checkValue: 3267,
				operators:  []int{81, 40, 27},
			},
			mathOperatorList: legalOperatorsTask01,
			want:             true,
		},
		{
			name: "AOC task 01 Not Reachable Example",
			fields: fields{
				checkValue: 7290,
				operators:  []int{6, 8, 6, 15},
			},
			mathOperatorList: legalOperatorsTask01,
			want:             false,
		},
		{
			name: "AOC task 02 Reachable Example",
			fields: fields{
				checkValue: 7290,
				operators:  []int{6, 8, 6, 15},
			},
			mathOperatorList: legalOperatorsTask02,
			want:             true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CalibratorEquations{
				checkValue: tt.fields.checkValue,
				operators:  tt.fields.operators,
			}
			if got := c.IsCheckValueReachable(tt.mathOperatorList); got != tt.want {
				t.Errorf("IsCheckValueReachable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_task01(t *testing.T) {
	lines, err := util.ReadFilePerLine("../../input/day07_example.txt")
	if err != nil {
		log.Fatal(err)
	}
	ces := make([]*CalibratorEquations, 0)
	for i := 0; i < len(lines); i++ {
		ces = append(ces, newCalibratorEquations(lines[i]))
	}

	tests := []struct {
		name string
		want int
	}{
		{
			name: "AOC Example",
			want: 3749,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := task01(ces); got != tt.want {
				t.Errorf("task01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_task02(t *testing.T) {
	lines, err := util.ReadFilePerLine("../../input/day07_example.txt")
	if err != nil {
		log.Fatal(err)
	}
	ces := make([]*CalibratorEquations, 0)
	for i := 0; i < len(lines); i++ {
		ces = append(ces, newCalibratorEquations(lines[i]))
	}

	tests := []struct {
		name string
		want int
	}{
		{
			name: "AOC Example",
			want: 11387,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := task02(ces); got != tt.want {
				t.Errorf("task02() = %v, want %v", got, tt.want)
			}
		})
	}
}
