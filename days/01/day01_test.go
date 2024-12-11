package main

import "testing"

func Test_receiveSides(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "simple",
			args: args{
				line: "12 21",
			},
			want:  12,
			want1: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := receiveSides(tt.args.line)
			if got != tt.want {
				t.Errorf("receiveSides() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("receiveSides() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
