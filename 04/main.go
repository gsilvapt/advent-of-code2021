package main

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
	"strings"
)

const INPUT_PATH string = "../inputs/day04"

func part2(draws []int, boards []*Board) int {
	for _, d := range draws {
		var nonWinning []*Board = []*Board{}
		for _, b := range boards {
			b.UpdateDraws(d)
			if !b.isWinner() {
				nonWinning = append(nonWinning, b)
			} else if b.isWinner() && len(boards) == 1 {
				sum := boards[0].Sum(false)
				fmt.Printf("Got last board with draw %v\n", d)
				return d * sum
			}
		}
		boards = nonWinning
	}

	fmt.Println("Iterated all boards and draws, could not the last board - Something is wrong!?")
	return 0
}

func part1(draws []int, boards []*Board) int {
	for _, d := range draws {
		for _, b := range boards {
			b.UpdateDraws(d)
			if b.isWinner() {
				sum := b.Sum(false)
				fmt.Printf("Got winner board with draw %v\n", d)
				return d * sum
			}
		}
	}
	return 0
}

func readDraw(inputs []string) []int {
	var draw []int
	sD := inputs[0]
	for _, i := range strings.Split(sD, ",") { // just the first line, split by ,
		v, _ := strconv.Atoi(i)
		draw = append(draw, v)
	}
	return draw
}

// func createBoards(inputs []string) (boards []*Board) {
// 	return boards
// }

func createBoards(cards [][]string) []*Board {
	var boards []*Board = []*Board{}
	for _, c := range cards { // each bingo card will contain gt
		boards = append(boards, NewBoard(c))
	}
	return boards
}

func main() {
	data := utils.ReadChunks(INPUT_PATH)
	var draws []int = readDraw(data[0])
	var boards []*Board = createBoards(data[1:])

	fmt.Printf("[part1]: %v", part1(draws, boards))
	fmt.Printf("[part2]: %v", part2(draws, boards))

}
