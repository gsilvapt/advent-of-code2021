package main

import (
	"aoc2021/utils"
	"fmt"
)

const INPUT_PATH string = "../inputs/day01"

func part1(input []int) int {
	var increments int = 0
	for i := 0; i < len(input)-1; i++ {
		if input[i+1]-input[i] > 0 {
			increments += 1
		}
	}
	return increments

}

func getThreeIncrs(input []int) []int {
	var threeMeasure []int
	for i := 0; i < len(input)-2; i++ {
		sum := input[i] + input[i+1] + input[i+2]
		threeMeasure = append(threeMeasure, sum)
	}
	return threeMeasure
}

func part2(input []int) int {
	return part1(getThreeIncrs(input))
}

func main() {
	var result int
	file := utils.ReadFile(INPUT_PATH)
	input := utils.SplitNumsByLine(file)
	result = part1(input)
	fmt.Printf("[part1]: got %d measurements larger than the previous \n", result)

	result = part2(input)
	fmt.Printf("[part2]: got %d measurements larger than the previous \n", result)
}
