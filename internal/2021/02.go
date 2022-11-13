package y2021

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/benjyiw/aoc/internal/files"
)

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

func Run02(inputFile string, part int) error {
	if inputFile == "" {
		inputFile = "./internal/2021/inputs/02/example.txt"
	}
	lines, err := files.ReadFile(inputFile)
	if err != nil {
		return err
	}

	var depth, horizontal int
	switch part {
	case 1:
		depth, horizontal = calcPart1Depth(lines)
	case 2:
		depth, horizontal = calcPart2Depth(lines)
	default:
		return errors.New("invalid part specified")
	}

	fmt.Println(depth * horizontal)
	return nil
}