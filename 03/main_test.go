package main

import "testing"

var inputs []string = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestPart1(t *testing.T) {
	testCases := []struct {
		input   []string
		gamma   int
		epsilon int
		res     int
	}{
		{input: inputs, gamma: 22, epsilon: 9, res: 198},
	}

	var gamma, epsilon int = part1(inputs)
	var res int = gamma * epsilon

	for _, tc := range testCases {
		if tc.gamma != gamma {
			t.Errorf("failed computing gammas. Expected %d and got %d", tc.gamma, gamma)
		}

		if tc.epsilon != epsilon {
			t.Errorf("failed computing epsilon. Expected %d and got %d", tc.epsilon, epsilon)
		}

		if tc.res != res {
			t.Errorf("failed computing result. Expected %d and got %d", tc.res, res)
		}
	}

}

func TestPart2(t *testing.T) {
	testCases := []struct {
		input []string
		oc    int
		co2   int
		sr    int
	}{
		{input: inputs, oc: 23, co2: 10, sr: 230},
	}

	var oc, co2 int = part2(inputs)
	var sr int = oc * co2

	for _, tc := range testCases {
		if tc.co2 != co2 {
			t.Errorf("failed computing co2 rating. Expected %d and got %d", tc.co2, co2)
		}

		if tc.oc != oc {
			t.Errorf("failed computing oxygen rating. Expected %d and got %d", tc.oc, oc)
		}

		if tc.sr != sr {
			t.Errorf("failed computing life support rating. Expected %d and got %d", tc.sr, sr)
		}
	}
}
