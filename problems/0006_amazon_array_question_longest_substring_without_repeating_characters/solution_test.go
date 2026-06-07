package p0006_amazon_array_question_longest_substring_without_repeating_characters

import "testing"

func TestAmazonArrayQuestionLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example 1",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "example 2",
			s:    "bbbbb",
			want: 1,
		},
		{
			name: "example 3",
			s:    "pwwkew",
			want: 3,
		},
		{
			name: "empty",
			s:    "",
			want: 0,
		},
		{
			name: "single char",
			s:    "a",
			want: 1,
		},
		{
			name: "repeating pattern",
			s:    "abba",
			want: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := AmazonArrayQuestionLongestSubstringWithoutRepeatingCharacters(tc.s)
			if got != tc.want {
				t.Fatalf("AmazonArrayQuestionLongestSubstringWithoutRepeatingCharacters(%q) = %d, want %d", tc.s, got, tc.want)
			}
		})
	}
}
