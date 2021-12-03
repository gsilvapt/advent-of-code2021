package main

import (
	"aoc2021/utils"
	"bytes"
	"fmt"
	"strconv"
)

const INPUT_PATH string = "../inputs/day03"

func parseBinString(s string) int {
	i, e := strconv.ParseInt(s, 2, 64)
	if e != nil {
		fmt.Printf("could not parse string: %s \n", s)
	}

	return int(i)
}

func getRating(inputs []string, oxygen bool) int {
	var remainingKeys []string = inputs
	for i := 0; i < len(inputs[0]); i++ {
		cb := getCommonBit(remainingKeys, i, oxygen)
		remainingKeys = remainMatchingKeys(remainingKeys, cb, i)
	}

	if len(remainingKeys) != 1 {
		fmt.Println("failed computing oxygen rating: last key is invalid")
	}

	return parseBinString(remainingKeys[0])
}

// Returns subset of strings where position bit matches the provided bit.
func remainMatchingKeys(input []string, bit string, pos int) []string {
	if len(input) == 1 {
		return input
	}

	var res []string
	for _, k := range input {
		if string(k[pos]) != bit {
			continue
		}

		res = append(res, k)
	}

	return res
}

// Returns most common bit in provided string position. rev flag works to get the reverse result, useful to compute the
// scruber rating
func getCommonBit(inputs []string, i int, rev bool) (cb string) {
	var zc, oc int
	for _, k := range inputs { // for each key in inputs
		switch string(k[i]) {
		case "0":
			zc++
		case "1":
			oc++
		}
	}

	switch rev {
	case true:
		cb = "1"
		if oc >= zc {
			cb = "0"
		}
	default:
		cb = "0"
		if oc >= zc {
			cb = "1"
		}

	}
	return cb
}

func part1(inputs []string) (gamma, epsilon int) {
	var ks, kn int
	var gr, er bytes.Buffer
	ks = len(inputs[0])
	kn = len(inputs)
	for k := 0; k < ks; k++ { // for each entry in the keys
		var zc, oc int
		for n := 0; n < kn; n++ { // for each key
			v, _ := strconv.Atoi(string(inputs[n][k]))

			switch v {
			case 0:
				zc++
			case 1:
				oc++
			}
		}
		switch {
		case zc > oc:
			gr.WriteString("0")
			er.WriteString("1")
		case zc < oc:
			gr.WriteString("1")
			er.WriteString("0")
		}
	}

	return parseBinString(gr.String()), parseBinString(er.String())
}

func part2(inputs []string) (oxygen, co2 int) {
	return getRating(inputs, false), getRating(inputs, true)
}

func main() {
	var gamma, eps, or, cs int
	file := utils.ReadFile(INPUT_PATH)
	diagnostic := utils.SplitStringsByLine(file)

	gamma, eps = part1(diagnostic)
	fmt.Printf("[part1] gamma: %d, epsilon: %d, multi: %d \n", gamma, eps, gamma*eps)

	or, cs = part2(diagnostic)
	fmt.Printf("[part2] oxygen rating: %d, co2scrubber: %d, life support rating: %d \n", or, cs, or*cs)
}
