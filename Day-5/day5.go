package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

func parseLine(line string) int {
	low := 0
	high := 127
	total := 128
	left := 0
	right := 7
	num := 8
	for _, char := range line {
		if low != high {
			if char == 'F' {
				total = total / 2
				high = high - total
			} else if char == 'B' {
				total = total / 2
				low = low + total
			}
		} else {
			if char == 'L' {
				num = num / 2
				right = right - num
			} else if char == 'R' {
				num = num / 2
				left = left + num
			}
		}
	}
	return low*8 + left
}

func main() {
	file := readFile(os.Args[1])

	max := 0
	seats := make([]int, 0)
	for _, line := range file {
		seat := parseLine(line)
		seats = append(seats, seat)
		if seat > max {
			max = seat
		}
	}
	fmt.Println("Part 1:", max)

	sort.Ints(seats)
	prev := seats[0]

	for index, val := range seats {
		if index > 0 && val != prev+1 {
			break
		}
		prev = val
	}
	fmt.Println("Part 2:", prev+1)
}
