package main

import (
	"strconv"
	"strings"
)

type Board struct {
	Numbers [5][5]int
	Draws   [5][5]bool
}

func NewBoard(nums []string) *Board {
	b := &Board{Numbers: [5][5]int{}, Draws: [5][5]bool{}}
	for r, vals := range nums {
		var nums []int = parseNums(vals)
		for c, n := range nums {
			b.Numbers[r][c] = n
		}
	}

	return b
}

func (b *Board) UpdateDraws(n int) *Board {
	for row, cols := range b.Numbers {
		for col, val := range cols {
			if val == n {
				b.Draws[row][col] = true
				break
			}
		}
	}
	return b
}

func (b *Board) isWinner() bool {
NEXT_ROW:
	for _, cols := range b.Draws {
		for _, drawn := range cols {
			if !drawn {
				continue NEXT_ROW
			}
		}
		return true
	}

NEXT_COL:
	for c := 0; c < len(b.Draws); c++ {
		for _, row := range b.Draws {
			if !row[c] {
				continue NEXT_COL
			}
		}
		return true
	}
	return false
}

func (b *Board) Sum(marked bool) int {
	var sum int = 0
	for r, cols := range b.Draws {
		for c, drawn := range cols {
			if drawn == marked {
				sum += b.Numbers[r][c]
			}
		}
	}
	return sum
}

func parseNums(part string) []int {
	var nums []int = []int{}
	for _, p := range strings.Split(part, " ") {
		if p == "" {
			continue
		}

		val, _ := strconv.Atoi(p)
		nums = append(nums, val)
	}
	return nums
}
