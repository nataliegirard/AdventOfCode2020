package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func readFile(filename string) []int {
	lines := make([]int, 0)
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		l := scanner.Text()
		line, _ := strconv.Atoi(l)
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

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
	line := readFile(os.Args[1])
	sort.Ints(line)

	twoProduct := sumTwo(line)
	fmt.Println("Part 1:", twoProduct)

	threeProduct := sumThree(line)
	fmt.Println("Part 2:", threeProduct)
}
