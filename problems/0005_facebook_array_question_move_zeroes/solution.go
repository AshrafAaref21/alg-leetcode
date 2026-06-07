package p0005_facebook_array_question_move_zeroes

// https://leetcode.com/problems/move-zeroes/description/
// time complexity: O(n), where n is the length of the input array
// space complexity: O(1)

func FacebookArrayQuestionMoveZeroes(nums []int) {
	zeroIndex := 0
	n := len(nums)

	for nonZeroIndex := 0; nonZeroIndex < n; nonZeroIndex++ {
		if nums[nonZeroIndex] != 0 {
			nums[zeroIndex] = nums[nonZeroIndex]
			zeroIndex++
		}
	}

	for i := zeroIndex; i < n; i++ {
		nums[i] = 0
	}
}
