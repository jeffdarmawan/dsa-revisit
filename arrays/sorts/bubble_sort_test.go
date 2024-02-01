package sorts

import (
	"reflect"
	"testing"
)

func Test_bubbleSort(t *testing.T) {
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
			want: []int32{0, 3, 4, 5, 6, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bubbleSort(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
