package main

import (
	"advent/utils"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func withinReach(adapters []int, index int) int {
	count := 0
	for i := 1; i <= 3; i++ {
		if index+i >= len(adapters) {
			break
		}
		if adapters[index+i]-adapters[index] <= 3 {
			count++
		}
	}
	return count
}

func main() {
	file := utils.ReadFile(os.Args[1])

	adapters := make([]int, 0)
	for _, line := range file {
		num, _ := strconv.Atoi(line)
		adapters = append(adapters, num)
	}
	sort.Ints(adapters)

	ones := 0
	threes := 0
	first := 0
	for i := 0; i < len(adapters); i++ {
		diff := adapters[i] - first
		if diff == 1 {
			ones++
		} else if diff == 3 {
			threes++
		}
		first = adapters[i]
	}
	threes++

	fmt.Println("Part 1:", ones*threes)

	// Part 2 using dynamic programming
	adapters = append(adapters, 0, adapters[len(adapters)-1]+3)
	sort.Ints(adapters)
	dp := make([]int, len(adapters))
	dp[len(dp)-1] = 1
	for i := len(adapters) - 2; i >= 0; i-- {
		sum := 0
		for j := i + 1; j < len(adapters); j++ {
			if adapters[j]-adapters[i] <= 3 {
				sum += dp[j]
			} else {
				break
			}
		}
		dp[i] = sum
	}
	fmt.Println("Part 2:", dp[0])
}
