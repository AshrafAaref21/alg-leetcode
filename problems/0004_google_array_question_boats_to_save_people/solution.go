package p0004_google_array_question_boats_to_save_people

import "sort"

// https://leetcode.com/problems/boats-to-save-people/
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
