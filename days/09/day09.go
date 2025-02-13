package main

import (
	"AoC2024/util"
	"fmt"
	"log"
	"strconv"
)

type FileId int

func main() {
	lines, err := util.ReadFilePerLine("./input/day09.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := task01(lines[0])
	fmt.Println(result)
	result = task02(lines[0])
	fmt.Println(result)
}

func task01(line string) int {
	fileBlocks := toFileBlocks(line)

	return checksum(fileBlocks)
}

func toFileBlocks(str string) []FileId {
	fileBlocks := make([]FileId, 0)
	reverse_index := len(str) - 1
	var encodedValue int
	encodedInverseValue, _ := strconv.Atoi(string(str[reverse_index]))
	for i := range str {
		encodedValue, _ = strconv.Atoi(string(str[i]))
		var extendTo int
		if i == reverse_index {
			extendTo = len(fileBlocks) + encodedInverseValue
		} else {
			extendTo = len(fileBlocks) + encodedValue
		}
		for len(fileBlocks) < extendTo {
			if isFile(i) {
				fileBlocks = append(fileBlocks, FileId(i/2))
			} else {
				for !isFile(reverse_index) {
					reverse_index--
					encodedInverseValue, _ = strconv.Atoi(string(str[reverse_index]))
				}
				fileBlocks = append(fileBlocks, FileId(reverse_index/2))
				encodedInverseValue--
				if encodedInverseValue == 0 {
					reverse_index--
				}
			}
		}
		if i >= reverse_index {
			break
		}
	}
	return fileBlocks
}

func checksum(blocks []FileId) int {
	sum := 0
	for i := range blocks {
		sum += i * int(blocks[i])
	}
	return sum
}

func isFile(i int) bool {
	return i%2 == 0
}

func task02(line string) int {
	return len(line)
}
