package util

import (
	"bufio"
	"os"
)

func ReadFileLines(inputPath string) ([]string, error) {
	var lines []string

	file, err := os.Open(inputPath)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return lines, err
	}

	return lines, nil
}
