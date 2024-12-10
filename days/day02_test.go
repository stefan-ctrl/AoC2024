package days

import (
	"reflect"
	"testing"
)

func Test_castIntReport(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "67 65 62 60 57 56 55 52",
			args: args{
				line: "67 65 62 60 57 56 55 52",
			},
			want: []int{67, 65, 62, 60, 57, 56, 55, 52},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := castIntReport(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("castIntReport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isReportSafe(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				line: "7 6 4 2 1",
			},
			want: true,
		},
		{
			args: args{
				line: "1 2 7 8 9",
			},
			want: false,
		},
		{
			args: args{
				line: "9 7 6 2 1",
			},
			want: false,
		},
		{
			args: args{
				line: "1 3 2 4 5",
			},
			want: false,
		},
		{
			args: args{
				line: "8 6 3 4 4 1",
			},
			want: false,
		},
		{
			args: args{
				line: "1 3 6 7 9",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.line, func(t *testing.T) {
			if got := isReportSafe(tt.args.line); got != tt.want {
				t.Errorf("isReportSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}
