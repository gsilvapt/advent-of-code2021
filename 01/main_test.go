package main

import "testing"

func TestPart1(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	actual := part1(input)
	if actual != 7 {
		t.Errorf("sum should be 7, got %d", actual)
	}
}

func TestPart2(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expectedIncs := []int{607, 618, 618, 617, 647, 716, 769, 792}

	actualIncs := getThreeIncrs(input)
	for i := range expectedIncs {
		if expectedIncs[i] != actualIncs[i] {
			t.Errorf("3 increments do not match correctly: have %v and was expecting %v", actualIncs, expectedIncs)
		}
	}

	actual := part2(input)
	if actual != 5 {
		t.Errorf("sum should be 5, got %d", actual)
	}
}
