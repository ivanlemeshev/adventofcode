package main

import (
	"sort"
)

func FindTotalDistance(leftList, rightList []int) int {
	// Copy the slices to avoid modifying the original slices.
	left := make([]int, len(leftList))
	right := make([]int, len(rightList))

	copy(left, leftList)
	copy(right, rightList)

	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0

	for i := 0; i < len(leftList); i++ {
		distance := leftList[i] - rightList[i]
		if distance < 0 {
			distance = -distance
		}
		totalDistance += distance
	}

	return totalDistance
}
