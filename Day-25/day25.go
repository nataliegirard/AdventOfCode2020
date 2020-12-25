package main

import (
	"advent/utils"
	"fmt"
	"os"
	"strconv"
)

func determineLoopSize(targetValue int) int {
	value := 1
	loopSize := 0
	subjectNumber := 7

	for value != targetValue {
		value = value * subjectNumber
		value = value % 20201227
		loopSize++
	}

	return loopSize
}

func findEncryptionNumber(subjectNumber int, loopSize int) int {
	value := 1

	for i := 0; i < loopSize; i++ {
		value = value * subjectNumber
		value = value % 20201227
	}
	return value
}

func main() {
	file := utils.ReadFile(os.Args[1])

	cardPK, _ := strconv.Atoi(file[0])
	doorPK, _ := strconv.Atoi(file[1])

	cardLoopSize := determineLoopSize(cardPK)

	encryptionKey := findEncryptionNumber(doorPK, cardLoopSize)
	fmt.Println("Part 1:", encryptionKey)
}
