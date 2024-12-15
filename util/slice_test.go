package util

import (
	"reflect"
	"testing"
)

func TestSliceGetMiddleValue(t *testing.T) {
	type args struct {
		slice *[]any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "int - uneven",
			args: args{
				slice: &[]any{1, 2, 3},
			},
			want: 2,
		},
		{
			name: "int - even",
			args: args{
				slice: &[]any{1, 2, 3, 4},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceGetMiddleValue(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceGetMiddleValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
