package main

import (
	"advent/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func game(paper []int, length int) int {
	spoken := make(map[int]int)
	say := paper[0]

	for round := 1; round < length; round++ {
		saidOnRound := -1
		if val, ok := spoken[say]; ok {
			saidOnRound = val
		}
		spoken[say] = round
		// fmt.Println(round, say)

		if round < len(paper) {
			say = paper[round]
			continue
		}

		if saidOnRound == -1 {
			say = 0
		} else {
			say = round - saidOnRound
		}
	}

	return say
}

func main() {
	file := utils.ReadFile(os.Args[1])

	line := strings.Split(file[0], ",")
	input := []int{}
	for _, val := range line {
		num, _ := strconv.Atoi(val)
		input = append(input, num)
	}

	fmt.Println("Part 1:", game(input, 2020))
	fmt.Println("Part 2:", game(input, 30000000))
}
