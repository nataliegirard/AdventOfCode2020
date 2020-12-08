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

type replacement struct {
	location    int
	instruction string
}

func readFile(filename string) []string {
	lines := make([]string, 0)
	file, err := os.Open(filename)
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

func parseInstruction(line string) (string, int) {
	re := regexp.MustCompile(`(acc|jmp|nop) ([+|-][0-9]+)`)
	result := re.FindSubmatch([]byte(line))
	val, _ := strconv.Atoi(string(result[2]))
	return string(result[1]), val
}

func getNewValue(line string) int {
	opp, val := parseInstruction(line)
	switch opp {
	case "acc":
		return val
	case "jmp":
		return 0
	case "nop":
		return 0
	}
	return 0
}

func getNewLine(line string) int {
	opp, val := parseInstruction(line)
	switch opp {
	case "acc":
		return 1
	case "jmp":
		return val
	case "nop":
		return 1
	}
	return 0
}

func newInstruction(prev replacement) string {
	inst, _ := parseInstruction(prev.instruction)
	var newLine string
	if inst == "jmp" {
		newLine = strings.Replace(prev.instruction, "jmp", "nop", 1)
	} else if inst == "nop" {
		newLine = strings.Replace(prev.instruction, "nop", "jmp", 1)
	} else {
		return ""
	}
	return newLine
}

func main() {
	file := readFile(os.Args[1])

	acc := 0
	line := 0
	visited := make(map[int]bool)
	for line >= 0 && line < len(file) {
		if _, ok := visited[line]; ok {
			break
		}

		acc = acc + getNewValue(file[line])
		visited[line] = true
		line = line + getNewLine(file[line])
	}

	fmt.Println("Part 1:", acc)

	replaceList := make([]replacement, 0)
	for i, line := range file {
		prev := replacement{
			instruction: line,
			location:    i,
		}
		prev.instruction = newInstruction(prev)
		if prev.instruction != "" {
			replaceList = append(replaceList, prev)
		}

	}

	acc = 0
	line = 0
	prev := replaceList[0]
	replaceList = replaceList[1:]
	file[prev.location] = prev.instruction

	for line >= 0 && line < len(file) {
		if _, ok := visited[line]; ok {
			acc = 0
			line = 0
			visited = make(map[int]bool)

			file[prev.location] = newInstruction(prev) // restore old
			prev = replaceList[0]
			replaceList = replaceList[1:]
			file[prev.location] = prev.instruction // make a change
		}

		acc = acc + getNewValue(file[line])
		visited[line] = true
		line = line + getNewLine(file[line])
	}
	fmt.Println("Part 2:", acc)
}
