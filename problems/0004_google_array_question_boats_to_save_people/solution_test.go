package p0004_google_array_question_boats_to_save_people

import "testing"

func TestGoogleArrayQuestionBoatsToSavePeople(t *testing.T) {
	tests := []struct {
		name   string
		people []int
		limit  int
		want   int
	}{
		{
			name:   "example 1",
			people: []int{1, 2},
			limit:  3,
			want:   1,
		},
		{
			name:   "example 2",
			people: []int{3, 2, 2, 1},
			limit:  3,
			want:   3,
		},
		{
			name:   "example 3",
			people: []int{3, 5, 3, 4},
			limit:  5,
			want:   4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GoogleArrayQuestionBoatsToSavePeople(tc.people, tc.limit)
			if got != tc.want {
				t.Fatalf("GoogleArrayQuestionBoatsToSavePeople(%v, %d) = %d, want %d", tc.people, tc.limit, got, tc.want)
			}
		})
	}
}
