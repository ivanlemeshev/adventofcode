package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	inputFile, err := os.Open("input")
	if err != nil {
		log.Fatalf("Error opening input file: %v", err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	sumOfMul := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		sumOfMul += SumOfMul(line)
	}

	fmt.Printf("Sum of multiplication: %d\n", sumOfMul)
}
