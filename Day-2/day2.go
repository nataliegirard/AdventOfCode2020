package main

import (
	"advent/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseLine(line string) (int, int, string, string) {
	re := regexp.MustCompile(`([0-9]+)-([0-9]+) ([a-z]): (.+)`)
	result := re.FindSubmatch([]byte(line))
	min, _ := strconv.Atoi(string(result[1]))
	max, _ := strconv.Atoi(string(result[2]))
	char := string(result[3])
	pword := string(result[4])
	return min, max, char, pword
}

func checkMatchSimple(min int, max int, char string, pword string) bool {
	count := strings.Count(pword, char)
	return count >= min && count <= max
}

func checkMatchComplex(min int, max int, char string, pword string) bool {
	first := string(pword[min-1]) == char
	second := string(pword[max-1]) == char

	return (first || second) && !(first && second)
}

func main() {
	file := utils.ReadFile(os.Args[1])

	countSimple := 0
	countComplex := 0
	for _, line := range file {
		min, max, char, pword := parseLine(line)

		validSimple := checkMatchSimple(min, max, char, pword)
		if validSimple {
			countSimple++
		}

		validComplex := checkMatchComplex(min, max, char, pword)
		if validComplex {
			countComplex++
		}
	}
	fmt.Println("Part 1:", countSimple)
	fmt.Println("Part 2:", countComplex)
}
