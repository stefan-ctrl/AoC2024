package main

import (
	"AoC2024/util"
	"fmt"
	"log"
	"strconv"
)

type FileId int

var writtenBlocks = make(map[int]bool)
var movedBlocks = make(map[int]bool)

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

func task02(line string) int {
	fileBlocks := toFileBlocksComplete(line)
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

func toFileBlocksComplete(str string) []FileId {
	fileBlocks := make([]FileId, 0)
	reverse_index := len(str) - 1
	var blockSizeToWrite int
	var appendGap bool
	reverseBlockSizeToWrite, _ := strconv.Atoi(string(str[reverse_index]))
	for i := range str {
		blockSizeToWrite, _ = strconv.Atoi(string(str[i]))
		var extendTo int
		if i == reverse_index {
			extendTo = len(fileBlocks) + reverseBlockSizeToWrite
		} else {
			extendTo = len(fileBlocks) + blockSizeToWrite
		}
		for len(fileBlocks) < extendTo && !appendGap {
			if isFileAndNotMoved(i) {
				fileBlocks = append(fileBlocks, FileId(i/2))
				writtenBlocks[i] = true
			} else if isFileMoved(i) {
				fileBlocks = append(fileBlocks, -1)
			} else {
				for !isFileAndNotMoved(reverse_index) {
					reverse_index--
					if reverse_index >= 0 {
						reverseBlockSizeToWrite, _ = strconv.Atoi(string(str[reverse_index]))
					} else {
						appendGap = true
					}
				}
				if extendTo-len(fileBlocks) >= reverseBlockSizeToWrite {
					fileBlocks = append(fileBlocks, FileId(reverse_index/2))
					reverseBlockSizeToWrite--
					if reverseBlockSizeToWrite == 0 {
						movedBlocks[reverse_index] = true
						writtenBlocks[reverse_index] = true
						reverse_index--
					}
				} else {
					reverse_index--
				}
				if appendGap {
					fileBlocks = append(fileBlocks, -1)
				}
			}
		}
		if len(writtenBlocks)*2 >= len(str) {
			break
		} else {
			reverse_index = len(str) - 1
			reverseBlockSizeToWrite, _ = strconv.Atoi(string(str[reverse_index]))
			appendGap = false
		}
	}
	return fileBlocks
}

func checksum(blocks []FileId) int {
	sum := 0
	for i := range blocks {
		v := blocks[i]
		if blocks[i] == -1 {
			v = 0
		}
		sum += i * int(v)
	}
	return sum
}

func isFile(i int) bool {
	return i%2 == 0
}

func isFileAndNotMoved(i int) bool {
	return i%2 == 0 && !isFileMoved(i)
}

func isFileMoved(i int) bool {
	_, found := movedBlocks[i]
	return found
}
