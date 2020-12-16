package main

import (
	"advent/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ranges struct {
	name         string
	lowerFirst   int
	lowerSecond  int
	higherFirst  int
	higherSecond int
	fieldIndex   int
}

func followsRule(rule ranges, value int) bool {
	if value >= rule.lowerFirst && value <= rule.lowerSecond {
		return true
	}
	if value >= rule.higherFirst && value <= rule.higherSecond {
		return true
	}

	return false
}

func inRange(rules []ranges, value int) bool {
	for _, rule := range rules {
		if followsRule(rule, value) {
			return true
		}
	}

	return false
}

func findInvalid(rules []ranges, nearby [][]int) ([]int, [][]int) {
	invalid := []int{}
	validTickets := [][]int{}
	for i := 0; i < len(nearby); i++ {
		isInvalid := false
		for j := 0; j < len(nearby[i]); j++ {
			if !inRange(rules, nearby[i][j]) {
				invalid = append(invalid, nearby[i][j])
				isInvalid = true
				break
			}
		}
		if !isInvalid {
			validTickets = append(validTickets, nearby[i])
		}
	}

	return invalid, validTickets
}

func findField(rule []int, usedIndexes []int) int {
	for _, r := range rule {
		used := false
		for _, u := range usedIndexes {
			if r == u {
				used = true
				break
			}
		}
		if !used {
			return r
		}
	}
	return -1
}

func determineFields(rules []ranges, nearby [][]int) []ranges {
	validFields := [][]int{}
	for _, rule := range rules {
		validIndexes := []int{}
		for i := 0; i < len(nearby[0]); i++ {
			valid := true
			for j := 0; j < len(nearby); j++ {
				if !followsRule(rule, nearby[j][i]) {
					valid = false
					break
				}
			}

			if valid {
				validIndexes = append(validIndexes, i)
			}
		}
		validFields = append(validFields, validIndexes)
	}

	usedIndexes := []int{}
	for i := 1; i < len(nearby); i++ {
		for index, rule := range validFields {
			if len(rule) == i {
				field := findField(rule, usedIndexes)
				rules[index].fieldIndex = field
				usedIndexes = append(usedIndexes, field)
			}
		}
	}

	return rules
}

func main() {
	file := utils.ReadFile(os.Args[1])

	nearby := [][]int{}
	myticket := []int{}
	rules := []ranges{}
	phase := 0
	for _, line := range file {
		if line == "" {
			phase++
			continue
		}

		if phase == 0 {
			// rules
			re := regexp.MustCompile(`^([a-z\s]+): ([0-9]+-[0-9]+) or ([0-9]+-[0-9]+)$`)
			result := re.FindSubmatch([]byte(line))
			low := strings.Split(string(result[2]), "-")
			lowFirst, _ := strconv.Atoi(low[0])
			lowSecond, _ := strconv.Atoi(low[1])
			high := strings.Split(string(result[3]), "-")
			highFirst, _ := strconv.Atoi(high[0])
			highSecond, _ := strconv.Atoi(high[1])
			newRule := ranges{
				name:         string(result[1]),
				lowerFirst:   lowFirst,
				lowerSecond:  lowSecond,
				higherFirst:  highFirst,
				higherSecond: highSecond,
			}
			rules = append(rules, newRule)
		} else if phase == 1 {
			// your ticket
			if strings.Contains(line, "your ticket") {
				continue
			}
			ticket := strings.Split(line, ",")
			for _, val := range ticket {
				num, _ := strconv.Atoi(val)
				myticket = append(myticket, num)
			}
		} else {
			// nearby tickets
			if strings.Contains(line, "nearby tickets") {
				continue
			}
			newTicket := []int{}
			ticket := strings.Split(line, ",")
			for _, val := range ticket {
				num, _ := strconv.Atoi(val)
				newTicket = append(newTicket, num)
			}
			nearby = append(nearby, newTicket)
		}
	}

	invalidNumbers, validTickets := findInvalid(rules, nearby)
	sum := 0
	for _, val := range invalidNumbers {
		sum = sum + val
	}
	fmt.Println("Part 1:", sum)

	validTickets = append(validTickets, myticket)
	rules = determineFields(rules, validTickets)

	total := 1
	for _, rule := range rules {
		if strings.Contains(rule.name, "departure") {
			total = total * myticket[rule.fieldIndex]
		}
	}
	fmt.Println("Part 2:", total)
}
