package main

import (
	"advent/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func processRules(rules map[int]string, index int, part2 bool) string {
	parsedRule := ""
	parts := strings.Split(rules[index], " ")

	for _, part := range parts {
		if part == "|" {
			parsedRule = parsedRule + "|"
			continue
		}

		num, _ := strconv.Atoi(part)
		var exp string
		re := regexp.MustCompile("a|b")

		// if part2 {
		// 	var exp42 string
		// 	var exp31 string

		// 	if num == 8 {
		// 		if !re.Match([]byte(rules[42])) {
		// 			exp42 = processRules(rules, 42, part2)
		// 			rules[42] = exp42
		// 		} else {
		// 			exp42 = rules[42]
		// 		}

		// 		rules[8] = "(?:" + exp42 + ")+"
		// 	} else if num == 11 {
		// 		if !re.Match([]byte(rules[42])) {
		// 			exp42 = processRules(rules, 42, part2)
		// 			rules[42] = exp42
		// 		} else {
		// 			exp42 = rules[42]
		// 		}

		// 		if !re.Match([]byte(rules[31])) {
		// 			exp31 = processRules(rules, 31, part2)
		// 			rules[31] = exp31
		// 		} else {
		// 			exp31 = rules[31]
		// 		}

		// 		rules[11] = "(" + exp42 + "(?:(?1))" + exp31 + ")"
		// 	}
		// }

		if part2 {
			if num == 8 {
				var exp8 string
				if !re.Match([]byte(rules[42])) {
					exp = processRules(rules, 42, part2)
					rules[42] = exp
				} else {
					exp = rules[42]
				}
				exp8 = "(" + exp + ")+"
				rules[8] = exp8
			} else if num == 11 {
				var exp11 string
				if !re.Match([]byte(rules[42])) {
					exp = processRules(rules, 42, part2)
					rules[42] = exp
				} else {
					exp = rules[42]
				}
				exp11 = "(" + exp + ")+"

				if !re.Match([]byte(rules[31])) {
					exp = processRules(rules, 31, part2)
					rules[31] = exp
				} else {
					exp = rules[31]
				}
				exp11 = exp11 + "(" + exp + ")+"
				rules[11] = exp11
			}
		}

		if !re.Match([]byte(rules[num])) {
			exp = processRules(rules, num, part2)
			rules[num] = exp
		} else {
			exp = rules[num]
		}

		if strings.Contains(exp, "|") {
			parsedRule = parsedRule + "(?:" + exp + ")"
		} else {
			parsedRule = parsedRule + exp
		}
	}
	return parsedRule
}

func main() {
	file := utils.ReadFile(os.Args[1])

	rules := make(map[int]string)
	phase := 0
	var ruleZero string
	var ruleZero2 string
	count := 0
	count2 := 0
	var re *regexp.Regexp
	var re2 *regexp.Regexp
	for _, line := range file {
		if line == "" {
			phase++

			rules2 := make(map[int]string)
			for k, v := range rules {
				rules2[k] = v
			}

			ruleZero = processRules(rules, 0, false)
			re = regexp.MustCompile("^" + ruleZero + "$")

			ruleZero2 = processRules(rules2, 0, true)
			re2 = regexp.MustCompile("^" + ruleZero2 + "$")
		}

		if phase == 0 {
			parts := strings.Split(line, ": ")
			number, _ := strconv.Atoi(parts[0])
			rules[number] = strings.Replace(parts[1], "\"", "", 2)
		} else {
			match := re.Match([]byte(line))
			match2 := re2.Match([]byte(line))

			if match {
				count++
			}
			if match2 {
				result := re2.FindSubmatch([]byte(line))
				n := len(result[1])
				substrings := []string{}
				for i := 0; i < len(line); i = i + n {
					substrings = append(substrings, line[i:i+n])
				}

				re42 := regexp.MustCompile(rules[42])
				re31 := regexp.MustCompile(rules[31])
				count42 := 0
				count31 := 0
				for _, sub := range substrings {
					if re42.Match([]byte(sub)) {
						count42++
					} else if re31.Match([]byte(sub)) {
						count31++
					}
				}

				if count42 > count31 {
					count2++
				}
			}
		}
	}

	fmt.Println("Part 1:", count)
	fmt.Println("Part 2:", count2)
}
