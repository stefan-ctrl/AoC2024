package main

import (
	"AoC2024/util"
	"fmt"
	"log"
	"strconv"
)

const EmptyField = "."
const Antenna = "#"


type File struct {
	id   int
	blocks []int
}

type FreeBlocks []int

func main() {
	lines, err := util.ReadFilePerLine("./input/day08.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := task01(lines[0])
	fmt.Println(result)
	result = task02(lines[0])
	fmt.Println(result)
}

func task01(line string) int {
	dataBlocks := toDataBlocks(line)
	fileBlocks := make([]File, 0)
	idCounter := 0
	for i := 0; i < len(dataBlocks); i++ {
		if i%2 == 0 {
			fileBlocks = append(fileBlocks, File{blocks: slices., id: idCounter})
			idCounter++
		}
	}

	for i := len(fileBlocks) - 1; i >= 0; i-- {
		arrange(&dataBlocks, fileBlocks[i])
	}
	return 0
}

func task02(line string) int {
	return 0
}

func toDataBlocks(str string) []int {
	blocks := make([]int, len(str))
	for i := range str {
		blocks[i], _ = strconv.Atoi(string(str[i]))
	}
	return blocks
}
