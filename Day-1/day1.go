package main

import (
	"advent/utils"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func sumTwo(line []int) int {
	bottom := 0
	top := len(line) - 1
	sum := line[bottom] + line[top]

	for sum != 2020 {
		if sum > 2020 {
			top--
		} else {
			bottom++
		}
		sum = line[bottom] + line[top]
	}

	return line[bottom] * line[top]
}

func sumThree(line []int) int {
	for i := 0; i < len(line); i++ {
		a := line[i]
		bc := 2020 - a

		for j := i + 1; j < len(line); j++ {
			b := line[j]
			if b >= bc {
				break
			}

			c := 2020 - (a + b)
			for k := 0; k < len(line); k++ {
				if line[k] == a || line[k] == b {
					continue
				}
				if line[k] == c {
					return a * b * c
				}
			}
		}
	}
	return 0
}

func main() {
	lines := utils.ReadFile(os.Args[1])
	line := []int{}
	for _, l := range lines {
		n, _ := strconv.Atoi(l)
		line = append(line, n)
	}
	sort.Ints(line)

	twoProduct := sumTwo(line)
	fmt.Println("Part 1:", twoProduct)

	threeProduct := sumThree(line)
	fmt.Println("Part 2:", threeProduct)
}
