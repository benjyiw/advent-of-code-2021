package y2021

import (
	"errors"
	"strconv"

	"github.com/benjyiw/aoc/internal/utils"
)

func convertStringSpliceToInt(lines []string) []int {
	var newLines []int
	for _, line := range lines {
		lineValue, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		newLines = append(newLines, lineValue)
	}
	return newLines
}

func calcPart2Ints(lines []int) []int {
	var depths []int
	window := 3

	for i := 0; i <= len(lines)-window; i++ {
		var windowDepth int
		for _, depth := range lines[i : i+window] {
			windowDepth += depth
		}
		depths = append(depths, windowDepth)
	}

	return depths
}

func calcIncreases(depths []int) int {
	var previousDepth int
	var totalIncreases int
	for i, depth := range depths {
		if i != 0 {
			if depth > previousDepth {
				totalIncreases++
			}
		}
		previousDepth = depth
	}
	return totalIncreases
}

func Run01(inputFile string, part int) (output int, err error) {
	if inputFile == "" {
		inputFile = "./internal/2021/inputs/01/example.txt"
	}
	lines, err := utils.ReadFile(inputFile)
	if err != nil {
		return
	}

	depths := convertStringSpliceToInt(lines)
	switch part {
	case 1:
		output = calcIncreases(depths)
	case 2:
		part2Depths := calcPart2Ints(depths)
		output = calcIncreases(part2Depths)
	default:
		err = errors.New("invalid part specified")
	}
	return
}
