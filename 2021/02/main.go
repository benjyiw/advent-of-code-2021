package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func readFile(filePath string) []string {
	file, err := os.Open(filePath)
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
	return lines
}

func calcPart1Depth(lines []string) (int, int) {
	var depth, horizontal int

	for _, line := range lines {
		lineFields := strings.Fields(line)

		modifier, err := strconv.Atoi(lineFields[1])
		if err != nil {
			panic(err)
		}

		switch lineFields[0] {
		case "forward":
			horizontal += modifier
		case "down":
			depth += modifier
		case "up":
			depth -= modifier
		}
	}

	return depth, horizontal
}

func calcPart2Depth(lines []string) (int, int) {
	var depth, horizontal, aim int

	for _, line := range lines {
		lineFields := strings.Fields(line)

		modifier, err := strconv.Atoi(lineFields[1])
		if err != nil {
			panic(err)
		}

		switch lineFields[0] {
		case "forward":
			horizontal += modifier
			depth += aim * modifier
		case "down":
			aim += modifier
		case "up":
			aim -= modifier
		}
	}

	return depth, horizontal
}

func main() {
	inputFile := flag.String("i", "inputs/input.txt", "input filepath")
	part := flag.Int("p", 1, "specify which part to execution [1,2]")
	flag.Parse()

	lines := readFile(*inputFile)

	var depth, horizontal int
	switch *part {
	case 1:
		depth, horizontal = calcPart1Depth(lines)
	case 2:
		depth, horizontal = calcPart2Depth(lines)
	default:
		fmt.Println("invalid part specified")
		os.Exit(1)
	}

	fmt.Println(depth*horizontal)
}
