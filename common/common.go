package common

import (
	"os"
	"strconv"
	"strings"
)

func ReadLinesFromInput(path string) ([]string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(string(bytes))
	lines := strings.Fields(trimmed)

	return lines, nil
}

func CommaSeparatedIntFromInput(path string) ([]int, error) {
	var ints []int

	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(string(bytes))
	numbers := strings.Split(trimmed, ",")
	for _, num := range numbers {
		n, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}

		ints = append(ints, n)
	}

	return ints, nil
}
