package main

import "testing"

func TestCountSafeReports(t *testing.T) {
	reports := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
		{8, 7, 6, 4, 2, 1},
	}
	expectedSafeReportsCount := 3
	actualsafeReportsCount := CountSafeReports(reports)
	if actualsafeReportsCount != expectedSafeReportsCount {
		t.Errorf("Expected safe reports count of %d but got %d", expectedSafeReportsCount, actualsafeReportsCount)
	}
}
