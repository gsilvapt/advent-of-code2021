package utils

import (
	"os"
	"strconv"
	"strings"
)

// SplitNumsByLine takes a slice of bytes and converts it to a slice of ints.
// Useful when the input provided is a list of numbers, for example.
func SplitNumsByLine(data []byte) []int {
	var nums []int
	for _, line := range strings.Split(string(data), "\n") {
		if len(line) == 0 {
			continue
		}

		num, err := strconv.Atoi(line)
		PanicIfError(err)

		nums = append(nums, num)
	}

	return nums
}

// SplitStringsByLine takes an array of bytes and converts it to a string.
// Useful when inputs contain strings and/or have a mixed set of characters.
func SplitStringsByLine(data []byte) []string {
	var lines []string
	for _, line := range strings.Split(string(data), "\n") {
		if len(line) == 0 {
			continue
		}

		lines = append(lines, string(line))
	}

	return lines
}

// ReadFile takes a path and tries to read it as a file. Panics on error.
func ReadFile(path string) []byte {
	data, err := os.ReadFile(path)
	PanicIfError(err)

	return data
}

// Custom struct to hold the instructions and their corresponding values
type InstVal struct {
	Instruction string
	Val         int
}

// Parses the binary data into
func SplitInstructions(data []byte) []*InstVal {
	var instructions []*InstVal
	for _, line := range strings.Split(string(data), "\n") {
		if len(line) == 0 {
			continue
		}

		iSet := strings.Split(line, " ")
		inst := iSet[0]
		val, _ := strconv.Atoi(iSet[1])

		instructions = append(instructions, &InstVal{Instruction: inst, Val: val})
	}

	return instructions
}
