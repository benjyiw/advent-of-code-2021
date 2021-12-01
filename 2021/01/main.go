package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func convertStringSpliceToInt(lines []string) []int {
	var newLines[]int
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

	for i := 0; i <= len(lines) - window; i++{
		var windowDepth int
		for _, depth := range lines[i:i+window] {
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

func main() {
	inputFile := flag.String("i", "inputs/example.txt", "input filepath")
	part := flag.Int("p", 1, "specify which part to execution [1,2]")
	flag.Parse()

	file, err := os.Open(*inputFile)
	if err != nil {
		panic(err) // don't care, just panic
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	depths := convertStringSpliceToInt(lines)
	if *part == 1 {
		fmt.Println("total increases:", calcIncreases(depths))
	} else if *part == 2 {
		part2Depths := calcPart2Ints(depths)
		fmt.Println("total increases:", calcIncreases(part2Depths))
	} else {
		fmt.Println("invalid part specified")
	}
}
