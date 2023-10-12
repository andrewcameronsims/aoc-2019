package common

import (
	"os"
	"strings"
)

func ReadInputFile(path string) ([]string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(string(bytes))
	lines := strings.Fields(trimmed)

	return lines, nil
}
