package main

import (
	"advent/utils"
	"fmt"
	"os"
)

func isTree(hill []string, row int, column int) bool {
	return rune(hill[row][column]) == '#'
}

func main() {
	file := utils.ReadFile(os.Args[1])

	countA := 0
	colA := 0
	countB := 0
	colB := 0
	countC := 0
	colC := 0
	countD := 0
	colD := 0
	countE := 0
	colE := 0

	for row, line := range file {
		if row == 0 {
			continue
		}

		// Slope A: Right 1, down 1
		colA = row % len(line)
		checkA := isTree(file, row, colA)
		if checkA {
			countA++
		}

		// Slope B: Right 3, down 1
		colB = (3 * row) % len(line)
		checkB := isTree(file, row, colB)
		if checkB {
			countB++
		}

		// Slope C: Right 5, down 1
		colC = (5 * row) % len(line)
		checkC := isTree(file, row, colC)
		if checkC {
			countC++
		}

		// Slope D: Right 7, down 1
		colD = (7 * row) % len(line)
		checkD := isTree(file, row, colD)
		if checkD {
			countD++
		}

		// Slope E: Right 1, down 2
		if row%2 == 1 {
			continue
		}
		colE = (colE + 1) % len(line)
		checkE := isTree(file, row, colE)
		if checkE {
			countE++
		}
	}
	fmt.Println("Part 1:", countB)
	fmt.Println("Part 2:", countA*countB*countC*countD*countE)
}
