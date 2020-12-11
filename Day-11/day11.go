package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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

func countOccupied(seating [][]string, row int, col int) int {
	count := 0

	for i := row - 1; i <= row+1; i++ {
		if i < 0 || i >= len(seating) {
			continue
		}
		for j := col - 1; j <= col+1; j++ {
			if j < 0 || j >= len(seating[i]) {
				continue
			}
			if i == row && j == col {
				continue
			}
			if seating[i][j] == "#" {
				count++
			}
		}
	}

	return count
}

func nextStep(seating [][]string) [][]string {
	new := [][]string{}
	for i := 0; i < len(seating); i++ {
		row := []string{}
		for j := 0; j < len(seating[i]); j++ {
			taken := countOccupied(seating, i, j)
			if seating[i][j] == "L" && taken == 0 {
				row = append(row, "#")
			} else if seating[i][j] == "#" && taken >= 4 {
				row = append(row, "L")
			} else {
				row = append(row, seating[i][j])
			}
		}
		new = append(new, row)
	}
	return new
}

func isTaken(seating [][]string, row int, col int) int {
	if seating[row][col] == "#" {
		return 1
	} else if seating[row][col] == "L" {
		return 0
	}
	return -1
}

func checkOccupancy(seating [][]string, row int, col int) int {
	count := 0

	// right
	for j := col + 1; j < len(seating[row]); j++ {
		taken := isTaken(seating, row, j)
		if taken >= 0 {
			count = count + taken // taken == 1 if occupied
			break
		}
	}

	// left
	for j := col - 1; j >= 0; j-- {
		taken := isTaken(seating, row, j)
		if taken >= 0 {
			count = count + taken // taken == 1 if occupied
			break
		}
	}

	// up
	for i := row - 1; i >= 0; i-- {
		taken := isTaken(seating, i, col)
		if taken >= 0 {
			count = count + taken // taken == 1 if occupied
			break
		}
	}

	// down
	for i := row + 1; i < len(seating); i++ {
		taken := isTaken(seating, i, col)
		if taken >= 0 {
			count = count + taken // taken == 1 if occupied
			break
		}
	}

	// down right
	for i := 1; ; i++ {
		if row+i >= len(seating) {
			break
		}
		if col+i >= len(seating[row]) {
			break
		}
		taken := isTaken(seating, row+i, col+i)
		if taken >= 0 {
			count = count + taken // taken == 1 if occupied
			break
		}
	}

	// down left
	for i := 1; ; i++ {
		if row+i >= len(seating) {
			break
		}
		if col-i < 0 {
			break
		}
		taken := isTaken(seating, row+i, col-i)
		if taken >= 0 {
			count = count + taken // taken == 1 if occupied
			break
		}
	}

	// up right
	for i := 1; ; i++ {
		if row-i < 0 {
			break
		}
		if col+i >= len(seating[row]) {
			break
		}
		taken := isTaken(seating, row-i, col+i)
		if taken >= 0 {
			count = count + taken // taken == 1 if occupied
			break
		}
	}

	// up left
	for i := 1; ; i++ {
		if row-i < 0 {
			break
		}
		if col-i < 0 {
			break
		}
		taken := isTaken(seating, row-i, col-i)
		if taken >= 0 {
			count = count + taken // taken == 1 if occupied
			break
		}
	}

	return count
}

func iterate(seating [][]string) [][]string {
	new := [][]string{}
	for i := 0; i < len(seating); i++ {
		row := []string{}
		for j := 0; j < len(seating[i]); j++ {
			taken := checkOccupancy(seating, i, j)
			if seating[i][j] == "L" && taken == 0 {
				row = append(row, "#")
			} else if seating[i][j] == "#" && taken >= 5 {
				row = append(row, "L")
			} else {
				row = append(row, seating[i][j])
			}
		}
		new = append(new, row)
	}
	return new
}

func seatingChange(seatingA [][]string, seatingB [][]string) bool {
	for i := 0; i < len(seatingA); i++ {
		for j := 0; j < len(seatingA[i]); j++ {
			if seatingA[i][j] != seatingB[i][j] {
				return true
			}
		}
	}

	return false
}

func main() {
	file := readFile(os.Args[1])

	seating := [][]string{}
	for _, line := range file {
		seating = append(seating, strings.Split(line, ""))
	}
	room := seating // used for part 2

	for true {
		newSeating := nextStep(seating)
		changed := seatingChange(seating, newSeating)
		seating = newSeating
		if !changed {
			break
		}
	}

	count := 0
	for i := 0; i < len(seating); i++ {
		for j := 0; j < len(seating[i]); j++ {
			if seating[i][j] == "#" {
				count++
			}
		}
	}
	fmt.Println("Part 1:", count)

	// t := 0
	for true {
		// t++
		newRoom := iterate(room)
		changed := seatingChange(room, newRoom)
		room = newRoom
		if !changed {
			break
		}
	}

	count = 0
	for i := 0; i < len(room); i++ {
		for j := 0; j < len(room[i]); j++ {
			if room[i][j] == "#" {
				count++
			}
		}
	}
	fmt.Println("Part 2:", count)
}
