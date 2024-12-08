package main

import "testing"

func TestSumOfMul(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	expectedSumOfMul := 161
	actualSumOfMul := SumOfMul(input)
	if actualSumOfMul != expectedSumOfMul {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expectedSumOfMul, actualSumOfMul)
	}
}
