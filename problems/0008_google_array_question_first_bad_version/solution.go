package p0008_google_array_question_first_bad_version

// https://leetcode.com/problems/first-bad-version/
// time complexity: O(log n), where n is the number of versions
// space complexity: O(1)
func GoogleArrayQuestionFirstBadVersion(n int) int {
	left, right := 1, n

	for left < right {
		mid := left + (right-left)/2

		if isBadVersion(mid) {
			if mid == 1 || (mid > 1 && !isBadVersion(mid-1)) {
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func isBadVersion(version int) bool {
	return version >= 4
}
