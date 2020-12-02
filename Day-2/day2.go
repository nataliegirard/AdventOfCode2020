package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFile(filename string) []string {
	lines := make([]string, 0)
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

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
	file := readFile(os.Args[1])

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
