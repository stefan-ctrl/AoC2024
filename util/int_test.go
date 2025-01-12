package util

import "testing"

func Test_IntToBinaryFixedLength(t *testing.T) {
	type args struct {
		number int
		base   int
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0 -> 0",
			args: args{
				number: 0,
				base:   2,
				length: 0,
			},
			want: "0",
		},
		{
			name: "2 -> 10",
			args: args{
				number: 2,
				base:   2,
				length: 1,
			},
			want: "10",
		},
		{
			name: "10 -> 0001010",
			args: args{
				number: 10,
				base:   2,
				length: 7,
			},
			want: "0001010",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToNewBaseFixedLength(tt.args.number, tt.args.base, tt.args.length); got != tt.want {
				t.Errorf("intToNewBaseFixedLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
