package main

import (
	"advent/utils"
	"fmt"
	"os"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func countGroup(choice map[string]int, num int) int {
	c := 0
	for _, i := range choice {
		if i == num {
			c++
		}
	}
	return c
}

func main() {
	file := utils.ReadFile(os.Args[1])

	answers := make([]string, 0)
	count := 0
	for _, line := range file {
		if line == "" {
			count = count + len(answers)
			answers = make([]string, 0)
			continue
		}
		for _, char := range line {
			if !contains(answers, string(char)) {
				answers = append(answers, string(char))
			}
		}
	}
	count = count + len(answers)
	fmt.Println("Part 1:", count)

	choice := make(map[string]int)
	total := 0
	num := 0
	for _, line := range file {
		if line == "" {
			total = total + countGroup(choice, num)
			choice = make(map[string]int)
			num = 0
			continue
		}

		num++
		for _, t := range line {
			char := string(t)
			if v, ok := choice[char]; ok {
				choice[char] = v + 1
			} else {
				choice[char] = 1
			}
		}
	}
	total = total + countGroup(choice, num)
	fmt.Println("Part 1:", total)
}
