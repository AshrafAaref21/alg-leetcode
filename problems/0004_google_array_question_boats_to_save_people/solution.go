package p0004_google_array_question_boats_to_save_people

import "sort"

// https://leetcode.com/problems/boats-to-save-people/
// time complexity: O(n log n) due to sorting, where n is the number of people
// space complexity: O(1) if we ignore the space used by sorting, otherwise O(n) due to sorting
func GoogleArrayQuestionBoatsToSavePeople(people []int, limit int) int {
	boats := 0
	sort.Ints(people)

	heavyP := len(people) - 1
	lightP := 0

	for heavyP >= lightP {
		if people[heavyP]+people[lightP] <= limit {
			lightP++
		}
		heavyP--
		boats++
	}

	return boats
}
