package p0003_container_with_most_water

func ContainerWithMostWater(height []int) int {
	maxArea := 0
	start, end := 0, len(height)-1

	for start < end {
		left, right := height[start], height[end]
		area := (end - start) * min(left, right)
		if area > maxArea {
			maxArea = area
		}

		if left < right {
			start++
		} else {
			end--
		}
	}

	return maxArea
}
