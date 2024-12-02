package main

import "testing"

func TestFindTotalDistance(t *testing.T) {
	leftList := []int{3, 4, 2, 1, 3, 3}
	rightList := []int{4, 3, 5, 3, 9, 3}
	expectedDistance := 11
	actualDistance := FindTotalDistance(leftList, rightList)
	if actualDistance != expectedDistance {
		t.Errorf("Expected distance of %d but got %d", expectedDistance, actualDistance)
	}
}
