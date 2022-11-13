package y2021

import (
	"strconv"
	"strings"

	"github.com/benjyiw/aoc/internal/utils"
)

func day03part1(lines []string) (output int, err error) {
	var positionCounts []int
	for i, line := range lines {
		if i == 0 {
			positionCounts = make([]int, len(line))
		}
		for j, char := range line {
			if char == '1' {
				positionCounts[j]++
			}
		}
	}

	half := len(lines) / 2
	gammaRateBuilder := make([]string, len(positionCounts))
	epsilonRateBuilder := make([]string, len(positionCounts))
	for i, p := range positionCounts {
		if p > half {
			gammaRateBuilder[i] = "1"
			epsilonRateBuilder[i] = "0"
		} else {
			gammaRateBuilder[i] = "0"
			epsilonRateBuilder[i] = "1"
		}
	}

	gammaRate, _ := strconv.ParseInt(strings.Join(gammaRateBuilder[:], ""), 2, 64)
	epsilonRate, _ := strconv.ParseInt(strings.Join(epsilonRateBuilder[:], ""), 2, 64)
	output = int(gammaRate * epsilonRate)
	return
}

func day03part2(lines []string) (output int, err error) {
	var oxygenGeneratorRatingBuilder []string
	var co2ScrubberRatingBuilder []string
	var lastCO2Line string

	for i := 0; i < len(lines[0]); i++ {
		oxygenUnfilteredLineCount := 0
		oxygenCurrentPositionCount := 0
		co2UnfilteredLineCount := 0
		co2CurrentPositionCount := 0

		for _, line := range lines {
			r := []rune(line)

			matchesOxygenFilter := true
			for f := range oxygenGeneratorRatingBuilder {
				if string(r[f]) != oxygenGeneratorRatingBuilder[f] {
					matchesOxygenFilter = false
					break
				}
			}
			if matchesOxygenFilter {
				oxygenUnfilteredLineCount++
				if r[i] == '1' {
					oxygenCurrentPositionCount++
				}
			}

			matchesCO2Filter := true
			for f := range co2ScrubberRatingBuilder {
				if string(r[f]) != co2ScrubberRatingBuilder[f] {
					matchesCO2Filter = false
					break
				}
			}
			if matchesCO2Filter {
				co2UnfilteredLineCount++
				if r[i] == '1' {
					co2CurrentPositionCount++
				}
				lastCO2Line = line
			}
		}

		oxygenHalf := float64(oxygenUnfilteredLineCount) / 2.0
		if float64(oxygenCurrentPositionCount) >= oxygenHalf {
			oxygenGeneratorRatingBuilder = append(oxygenGeneratorRatingBuilder, "1")
		} else {
			oxygenGeneratorRatingBuilder = append(oxygenGeneratorRatingBuilder, "0")
		}

		if co2UnfilteredLineCount != 1 {
			co2Half := float64(co2UnfilteredLineCount) / 2.0
			if float64(co2CurrentPositionCount) >= co2Half {
				co2ScrubberRatingBuilder = append(co2ScrubberRatingBuilder, "0")
			} else {
				co2ScrubberRatingBuilder = append(co2ScrubberRatingBuilder, "1")
			}
		}
	}

	oxygenGeneratorRating, _ := strconv.ParseInt(strings.Join(oxygenGeneratorRatingBuilder[:], ""), 2, 64)
	co2ScrubberRating, _ := strconv.ParseInt(lastCO2Line, 2, 64)
	output = int(oxygenGeneratorRating * co2ScrubberRating)
	return
}

func Run03(input string, part int) (output int, err error) {
	if input == "" {
		input = getExampleInput("03")
	}
	lines, err := utils.ReadFile(input)
	if err != nil {
		return
	}

	if part == 1 {
		output, err = day03part1(lines)
	} else if part == 2 {
		output, err = day03part2(lines)
	}

	return
}
