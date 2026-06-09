package p0008_google_array_question_first_bad_version

import "testing"

func TestGoogleArrayQuestionFirstBadVersion(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "n equals first bad",
			n:    4,
			want: 4,
		},
		{
			name: "n greater than first bad",
			n:    5,
			want: 4,
		},
		{
			name: "larger n",
			n:    100,
			want: 4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GoogleArrayQuestionFirstBadVersion(tc.n)
			if got != tc.want {
				t.Fatalf("GoogleArrayQuestionFirstBadVersion(%d) = %d, want %d", tc.n, got, tc.want)
			}
		})
	}
}
