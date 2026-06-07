package p0005_facebook_array_question_move_zeroes

import (
	"reflect"
	"testing"
)

func TestFacebookArrayQuestionMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "example 1",
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0},
			want: []int{0, 0, 0},
		},
		{
			name: "no zeros",
			nums: []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "mixed with trailing zeros",
			nums: []int{4, 0, 5, 0, 0, 3, 0, 1},
			want: []int{4, 5, 3, 1, 0, 0, 0, 0},
		},
		{
			name: "empty",
			nums: []int{},
			want: []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			FacebookArrayQuestionMoveZeroes(tc.nums)
			if !reflect.DeepEqual(tc.nums, tc.want) {
				t.Fatalf("FacebookArrayQuestionMoveZeroes() got %v, want %v", tc.nums, tc.want)
			}
		})
	}
}
