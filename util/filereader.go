package util

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
)

// GetDailyInput missing authentication
//func GetDailyInput(year, day int) string {
//	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
//	resp, err := http.Get(url)
//	if err != nil {
//		panic(err)
//	}
//	bodyBytes, err := io.ReadAll(resp.Body)
//	if err != nil {
//		panic(err)
//	}
//	return string(bodyBytes)
//}

func ReadFilePerLine(file string) ([]string, error) {
	filePath, _ := filepath.Abs(file)
	readFile, err := os.Open(filePath)
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

func ReadFile(file string) (string, error) {
	filePath, _ := filepath.Abs(file)
	readFile, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	result, err := io.ReadAll(readFile)
	if err != nil {
		return "", err
	}

	return string(result), readFile.Close()
}
