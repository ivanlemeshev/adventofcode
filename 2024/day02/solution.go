package main

func CountSafeReports(reports [][]int) int {
	totalReports := len(reports)
	unsafeReports := 0

	for _, report := range reports {
		increasing := false
		decreasing := false

		for i := 0; i < len(report)-1; i++ {
			diff := report[i] - report[i+1]

			if diff > 0 {
				increasing = true
			}

			if diff < 0 {
				decreasing = true
			}

			if increasing && decreasing {
				unsafeReports++
				break // usafe
			}

			if diff == 0 {
				unsafeReports++
				break // usafe
			}

			if diff > 0 && diff > 3 {
				unsafeReports++
				break // unsafe
			}

			if diff < 0 && diff < -3 {
				unsafeReports++
				break // unsafe
			}
		}
	}

	return totalReports - unsafeReports
}
