package y2021

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/benjyiw/aoc/internal/files"
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

func Run01(inputFile string, part int) error {
	if inputFile == "" {
		inputFile = "./internal/2021/inputs/01/example.txt"
	}

	lines, err := files.ReadFile(inputFile)
	if err != nil {
		return err
	}

	depths := convertStringSpliceToInt(lines)
	if part == 1 {
		fmt.Printf("total increases: %d\n", calcIncreases(depths))
	} else if part == 2 {
		part2Depths := calcPart2Ints(depths)
		fmt.Printf("total increases: %d\n", calcIncreases(part2Depths))
	} else {
		return errors.New("invalid part specified")
	}
	return nil
}
