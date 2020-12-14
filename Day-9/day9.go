package main

import (
	"advent/utils"
	"fmt"
	"os"
	"strconv"
)

func checkSum(num int, numbers []int) bool {

	return false
}

func checkNumber(numbers []int, index int, length int) bool {
	subList := numbers[index-length : index]
	for i := 0; i < len(subList)-1; i++ {
		for j := i + 1; j < len(subList); j++ {
			if subList[i]+subList[j] == numbers[index] {
				return true
			}
		}
	}
	return false
}

func main() {
	file := utils.ReadFile(os.Args[1])
	checkLen := 25

	numbers := make([]int, 0)
	for _, line := range file {
		num, _ := strconv.Atoi(line)
		numbers = append(numbers, num)
	}

	invalid := -1
	for i := checkLen; i < len(numbers); i++ {
		if !checkNumber(numbers, i, checkLen) {
			invalid = i
			break
		}
	}
	if invalid == -1 {
		fmt.Println("Something went wrong")
	}
	fmt.Println("Part 1:", numbers[invalid])

	firstIndex := 0
	curIndex := 2
	sum := numbers[0] + numbers[1]
	for sum != numbers[invalid] {
		if sum > numbers[invalid] || curIndex >= len(numbers) {
			firstIndex++
			sum = numbers[firstIndex] + numbers[firstIndex+1]
			curIndex = firstIndex + 2
		} else {
			sum = sum + numbers[curIndex]
			curIndex++
		}
	}

	min := numbers[invalid]
	max := 0
	for _, val := range numbers[firstIndex:curIndex] {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	fmt.Println("Part 2", max+min)
}
