package y2021

import "fmt"

const inputsDir = "./internal/2021/inputs/"

func getExampleInput(day string) string {
	return fmt.Sprintf("%s/%s/example.txt", inputsDir, day)
}
