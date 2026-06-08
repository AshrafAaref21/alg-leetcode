package p0007_arrays_question_find_first_and_last_position_of_element_in_sorted_array

func ArraysQuestionFindFirstAndLastPositionOfElementInSortedArray(nums []int, target int) []int {
	first := findFirst(nums, target)
	last := findLast(nums, target)
	return []int{first, last}
}

func findFirst(nums []int, target int) int {
	return -1
}

func findLast(nums []int, target int) int {
	return -1
}
