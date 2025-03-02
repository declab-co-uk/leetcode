package __two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "basic case",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{1, 0},
		},
		{
			name:   "negative numbers",
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{4, 2},
		},
		{
			name:   "zero target",
			nums:   []int{0, 2, 3, 0},
			target: 0,
			want:   []int{3, 0},
		},
		{
			name:   "duplicate numbers",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{1, 0},
		},
		{
			name:   "larger numbers",
			nums:   []int{1000005, 1000001, 1000002, 1000003},
			target: 2000003,
			want:   []int{2, 1},
		},
		{
			name:   "sorted array",
			nums:   []int{1, 2, 3, 4, 6},
			target: 7,
			want:   []int{3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
