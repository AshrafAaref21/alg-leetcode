package p0006_amazon_array_question_longest_substring_without_repeating_characters

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
// time complexity: O(n), where n is the length of the input string
// space complexity: O(min(m, n)), where m is the size of the character set and n is the length of the input string
func AmazonArrayQuestionLongestSubstringWithoutRepeatingCharacters(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	output := 0
	m := make(map[byte]int)
	right, left := 0, 0

	for i := 0; i < len(s); i++ {
		if val, ok := m[s[i]]; ok {
			left = int(max(left, val+1))
		}
		m[s[i]] = i
		right += 1
		output = int(max(output, right-left))
	}

	return output
}
