package main

import (
	"aoc2021/utils"
	"fmt"
)

const INPUT_PATH string = "../inputs/day02"

func processWithoutAim(instructions []*utils.InstVal) (dep, hor int) {
	for _, inst := range instructions {
		switch inst.Instruction {
		case "forward":
			hor += inst.Val
		case "down":
			dep += inst.Val
		case "up":
			dep -= inst.Val
		}
	}

	return hor, dep
}

func processWithAim(instructions []*utils.InstVal) (dep, hor int) {
	var aim int
	for _, inst := range instructions {
		switch inst.Instruction {
		case "forward":
			hor += inst.Val
			dep += (aim * inst.Val)
		case "down":
			aim += inst.Val
		case "up":
			aim -= inst.Val
		}
	}

	return hor, dep
}

func part1(instructions []*utils.InstVal) (hor, dep, multi int) {
	hor, dep = processWithoutAim(instructions)
	return hor, dep, hor * dep
}

func part2(instructions []*utils.InstVal) (hor, dep, multi int) {
	hor, dep = processWithAim(instructions)
	return hor, dep, hor * dep
}

func main() {
	file := utils.ReadFile(INPUT_PATH)
	instructions := utils.SplitInstructions(file)

	var hor, dep, multi int

	hor, dep, multi = part1(instructions)
	fmt.Printf("[part1] Solution: %d, %d, %d - hor, dep, multi \n", hor, dep, multi)

	hor, dep, multi = part2(instructions)
	fmt.Printf("[part2] Solution: %d, %d, %d - hor, dep, multi \n", hor, dep, multi)
}
