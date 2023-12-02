package main

import "testing"

func TestGetSumOfLine(t *testing.T) {
	testTable := map[string]int{
		"two1nine":         11,
		"eightwothree":     00,
		"abcone2threexyz":  22,
		"xtwone3four":      33,
		"4nineeightseven2": 42,
		"zoneight234":      24,
		"7pqrstsixteen":    77,
	}

	for k, v := range testTable {
		output := getSumOfLine(k)

		if output != v {
			t.Errorf("For line %s, expected %d but got %d", k, v, output)
		}
	}

}
