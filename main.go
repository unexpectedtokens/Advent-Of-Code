package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type testInput struct {
	Lines []string
}

func loadTestInput() *testInput {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	output := testInput{}

	output.Lines = strings.Split(string(data), "\n")

	return &output

}

func getSumOfLine(line string) int {
	splitStr := strings.Split(line, "")

	firstFound := false
	var first int
	var second int
	var index int
	for !firstFound {
		fmt.Println(line, index, first)
		curChar := splitStr[index]
		potNum, err := strconv.Atoi(curChar)
		if err == nil {
			first = potNum
			firstFound = true
			secondFound := false
			for _, charSecond := range splitStr[index:] {

				potSecNum, err := strconv.Atoi(charSecond)
				if err == nil {
					secondFound = true
					second = potSecNum
				}
			}

			if !secondFound {
				second = first
			}
		} else {
			index += 1
			if index == len(splitStr) {
				first = 0
				second = 0
			}
		}
	}

	outNum, err := strconv.Atoi(fmt.Sprintf("%d%d", first, second))
	if err != nil {
		panic(err)
	}

	// Leave for debug
	// fmt.Printf("For line (%s) got for first %d and for second %d]\n", line, first, second)
	// time.Sleep(time.Second * 2)

	return outNum

}

func main() {
	inputData := loadTestInput()

	var total int

	for _, line := range inputData.Lines {

		if len(line) != 0 {

			totalOfLine := getSumOfLine(line)
			total += totalOfLine
		}

	}

	fmt.Println(total)
}
