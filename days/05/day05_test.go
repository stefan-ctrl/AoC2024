package main

import (
	"AoC2024/util"
	"reflect"
	"testing"
)

var lines, err = util.ReadFilePerLine("../../input/day05_example.txt")

func Test_lineParser(t *testing.T) {
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name  string
		want  []PageOrderRule
		want1 []Order
	}{
		{
			name: "AOC_Example",
			want: []PageOrderRule{
				{47, 53},
				{97, 13},
				{97, 61},
				{97, 47},
				{75, 29},
				{61, 13},
				{75, 53},
				{29, 13},
				{97, 29},
				{53, 29},
				{61, 53},
				{97, 53},
				{61, 29},
				{47, 13},
				{75, 47},
				{97, 75},
				{47, 61},
				{75, 61},
				{47, 29},
				{75, 13},
				{53, 13},
			},
			want1: []Order{
				{75, 47, 61, 53, 29},
				{97, 61, 53, 29, 13},
				{75, 29, 13},
				{75, 97, 47, 61, 53},
				{61, 13, 29},
				{97, 13, 75, 29, 47},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := lineParser(lines)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lineParser() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("lineParser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUpdateOrder_IsOrderSuccessful(t *testing.T) {
	tests := []struct {
		name string
		u    Order
		want bool
	}{
		{
			name: "Line 1",
			u:    Order{75, 47, 61, 53, 29},
			want: true,
		},
		{
			name: "Line 2",
			u:    Order{97, 61, 53, 29, 13},
			want: true,
		},
		{
			name: "Line 3",
			u:    Order{75, 29, 13},
			want: true,
		},
		{
			name: "Line 4",
			u:    Order{75, 97, 47, 61, 53},
			want: false,
		},
		{
			name: "Line 5",
			u:    Order{61, 13, 29},
			want: false,
		},
		{
			name: "Line 6",
			u:    Order{97, 13, 75, 29, 47},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pageOrderRules, _ := lineParser(lines)
			notAllowed := inverseRuleset(pageOrderRules)
			if got := tt.u.IsOrderSuccessful(notAllowed); got != tt.want {
				t.Errorf("IsOrderSuccessful() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createRuleset(t *testing.T) {
	tests := []struct {
		name       string
		notAllowed map[int][]int
	}{
		{
			name: "AOC_Example",
			notAllowed: map[int][]int{
				13: {97, 61, 29, 47, 75, 53},
				29: {75, 97, 53, 61, 47},
				47: {97, 75},
				53: {47, 75, 61, 97},
				61: {97, 47, 75},
				75: {97},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pageOrderRules, _ := lineParser(lines)
			got1 := inverseRuleset(pageOrderRules)
			if !reflect.DeepEqual(got1, tt.notAllowed) {
				t.Errorf("inverseRuleset() got1 = %v, want %v", got1, tt.notAllowed)
			}
		})
	}
}

func TestUpdateOrder_CorrectOrder(t *testing.T) {
	tests := []struct {
		name string
		u    Order
		want Order
	}{
		{
			name: "61,13,29",
			u: Order{
				61, 13, 29,
			},
			want: Order{
				61, 29, 13,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pageOrderRules, _ := lineParser(lines)
			notAllowed := inverseRuleset(pageOrderRules)
			if got := tt.u.CorrectOrder(notAllowed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CorrectOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
