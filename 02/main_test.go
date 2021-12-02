package main

import (
	"aoc2021/utils"
	"testing"
)

var instructions []*utils.InstVal = []*utils.InstVal{
	{Instruction: "forward", Val: 5},
	{Instruction: "down", Val: 5},
	{Instruction: "forward", Val: 8},
	{Instruction: "up", Val: 3},
	{Instruction: "down", Val: 8},
	{Instruction: "forward", Val: 2},
}

type test struct {
	withAim     bool
	expectedHor int
	expectedDep int
	expectedRes int
}

var tests []*test = []*test{
	{false, 15, 10, 150},
	{true, 15, 60, 900},
}

func TestDepthCalculation(t *testing.T) {
	var actualHor, actualDep, actualRes int
	for _, tc := range tests {
		if tc.withAim {
			actualHor, actualDep = processWithAim(instructions)
		} else {
			actualHor, actualDep = processWithoutAim(instructions)
		}
		actualRes = actualHor * actualDep

		if tc.expectedHor != actualHor {
			t.Errorf("Expected %d horizontal and got %d", tc.expectedHor, actualHor)
		}

		if tc.expectedDep != actualDep {
			t.Errorf("Expected %d depth and got %d", tc.expectedDep, actualDep)
		}

		if tc.expectedRes != actualRes {
			t.Errorf("Expected %d as multiplication, got %d", tc.expectedRes, actualRes)
		}
	}
}

func TestDepthWithAim(t *testing.T) {

}
