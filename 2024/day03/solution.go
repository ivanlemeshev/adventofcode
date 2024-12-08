package main

import (
	"regexp"
	"strconv"
)

var reg = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func SumOfMul(input string) int {
	sum := 0

	matches := reg.FindAllString(input, -1)
	for _, match := range matches {
		submatches := reg.FindStringSubmatch(match)

		a, err := strconv.Atoi(submatches[1])
		if err != nil {
			panic("Error converting a to int")
		}

		b, err := strconv.Atoi(submatches[2])
		if err != nil {
			panic("Error converting b to int")
		}

		sum += a * b
	}

	return sum
}
