package p0007_arrays_question_find_first_and_last_position_of_element_in_sorted_array

import (
	"reflect"
	"testing"
)

func TestArraysQuestionFindFirstAndLastPositionOfElementInSortedArray(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "example 1",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		{
			name:   "example 2",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		{
			name:   "example 3",
			nums:   []int{},
			target: 0,
			want:   []int{-1, -1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ArraysQuestionFindFirstAndLastPositionOfElementInSortedArray(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("ArraysQuestionFindFirstAndLastPositionOfElementInSortedArray(%v, %d) = %v, want %v", tc.nums, tc.target, got, tc.want)
			}
		})
	}
}
