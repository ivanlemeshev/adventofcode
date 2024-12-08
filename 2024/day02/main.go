package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("input")
	if err != nil {
		log.Fatalf("Error opening input file: %v", err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	var reports [][]int

	for fileScanner.Scan() {
		reportLevels := strings.Split(fileScanner.Text(), " ")
		levels := make([]int, 0, len(reportLevels))
		for _, level := range reportLevels {
			levelInt, err := strconv.Atoi(level)
			if err != nil {
				log.Fatalf("Error converting level to int: %v", err)
			}

			levels = append(levels, levelInt)
		}

		reports = append(reports, levels)
	}

	safeReportsCount := CountSafeReports(reports)
	fmt.Printf("Safe reports count: %d\n", safeReportsCount)
}
