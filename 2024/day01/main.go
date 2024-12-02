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

	var leftList []int
	var rightList []int

	for fileScanner.Scan() {
		parts := strings.Split(fileScanner.Text(), "   ")
		if len(parts) != 2 {
			log.Fatalf("Invalid input line: %s", fileScanner.Text())
		}

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Error converting left to int: %v", err)
		}

		leftList = append(leftList, left)

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Error converting right to int: %v", err)
		}

		rightList = append(rightList, right)
	}

	totalDistance := FindTotalDistance(leftList, rightList)
	fmt.Printf("Total distance: %d\n", totalDistance)
}
