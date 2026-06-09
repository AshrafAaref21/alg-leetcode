package p0008_google_array_question_first_bad_version

import "testing"

func TestGoogleArrayQuestionFirstBadVersion(t *testing.T) {
	originalChecker := badVersionChecker
	t.Cleanup(func() {
		badVersionChecker = originalChecker
	})

	tests := []struct {
		name     string
		n        int
		firstBad int
		want     int
	}{
		{
			name:     "n equals first bad",
			n:        4,
			firstBad: 4,
			want:     4,
		},
		{
			name:     "n greater than first bad",
			n:        5,
			firstBad: 4,
			want:     4,
		},
		{
			name:     "larger n",
			n:        100,
			firstBad: 4,
			want:     4,
		},
		{
			name:     "first bad is 1",
			n:        10,
			firstBad: 1,
			want:     1,
		},
		{
			name:     "first bad near end",
			n:        10,
			firstBad: 9,
			want:     9,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			badVersionChecker = func(version int) bool {
				return version >= tc.firstBad
			}

			got := GoogleArrayQuestionFirstBadVersion(tc.n)
			if got != tc.want {
				t.Fatalf("GoogleArrayQuestionFirstBadVersion(%d) = %d, want %d", tc.n, got, tc.want)
			}
		})
	}
}
