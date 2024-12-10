package util

import (
	"bufio"
	"os"
)

func ReadFilePerLine(file string) ([]string, error) {
	readFile, err := os.Open(file)
	lines := make([]string, 0)
	if err != nil {
		return lines, err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines, readFile.Close()
}
