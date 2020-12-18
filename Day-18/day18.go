package main

import (
	"advent/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func evaluate(line string, advanced bool) int {
	equation := line
	countParens := strings.Count(line, "(")
	for countParens > 0 {
		startExp := strings.LastIndex(equation, "(")
		endExp := strings.Index(equation[startExp:], ")")
		exp := equation[startExp+1 : startExp+endExp]
		var expVal int
		if advanced {
			expVal = findValueAdvanced(exp)
		} else {
			expVal = findValue(exp)
		}
		val := fmt.Sprintf("%d", expVal)
		equation = equation[:startExp] + val + equation[startExp+endExp+1:]
		countParens--
	}
	var answer int
	if advanced {
		answer = findValueAdvanced(equation)
	} else {
		answer = findValue(equation)
	}
	return answer
}

// Has no parantheses
func findValue(line string) int {
	var result int
	var operator string
	values := strings.Split(line, " ")

	for _, val := range values {
		if val == "+" {
			operator = "+"
		} else if val == "*" {
			operator = "*"
		} else {
			integer, _ := strconv.Atoi(val)
			if operator == "+" {
				result = result + integer
			} else if operator == "*" {
				result = result * integer
			} else {
				result = integer
			}
		}
	}
	return result
}

func findValueAdvanced(line string) int {
	equation := strings.Split(line, " ")

	for i := 0; i < len(equation); i++ {
		if equation[i] == "+" {
			first, _ := strconv.Atoi(equation[i-1])
			second, _ := strconv.Atoi(equation[i+1])
			newVal := fmt.Sprintf("%d", first+second)
			newEquation := append(equation[:i-1], newVal)
			equation = append(newEquation, equation[i+2:]...)
			i--
		}
	}

	return findValue(strings.Join(equation, " ")) // should only have multiplications
}

func main() {
	file := utils.ReadFile(os.Args[1])
	sum := 0
	sumAdv := 0
	for _, line := range file {
		val := evaluate(line, false)
		valAdv := evaluate(line, true)
		sum = sum + val
		sumAdv = sumAdv + valAdv
	}
	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sumAdv)
}
