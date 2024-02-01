package reverse

import (
	"reflect"
	"testing"
)

func Test_reverseExternal(t *testing.T) {
	type args struct {
		in []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "success",
			args: args{
				in: []int32{4, 6, 3, 9, 0, 5},
			},
			want: []int32{5, 0, 9, 3, 6, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseExternal(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseExternal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseInternal(t *testing.T) {
	type args struct {
		in []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "success",
			args: args{
				in: []int32{4, 6, 3, 9, 0, 5},
			},
			want: []int32{5, 0, 9, 3, 6, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseInternal(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseInternal() = %v, want %v", got, tt.want)
			}
		})
	}
}
