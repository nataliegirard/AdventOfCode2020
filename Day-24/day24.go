package main

import (
	"advent/utils"
	"fmt"
	"os"
	"regexp"
)

type coord struct {
	x int
	y int
}

func runPaths(paths [][]string) (int, map[coord]bool) {
	blackTiles := map[coord]bool{}

	for _, path := range paths {
		currentCoord := coord{}
		for _, i := range path {
			switch i {
			case "e":
				currentCoord.x += 2
			case "ne":
				currentCoord.x++
				currentCoord.y++
			case "se":
				currentCoord.x++
				currentCoord.y--
			case "w":
				currentCoord.x -= 2
			case "sw":
				currentCoord.x--
				currentCoord.y--
			case "nw":
				currentCoord.x--
				currentCoord.y++
			}
		}

		if _, ok := blackTiles[currentCoord]; ok {
			blackTiles[currentCoord] = !blackTiles[currentCoord]
		} else {
			blackTiles[currentCoord] = true
		}
	}
	countBlack := 0
	countWhite := 0
	for _, i := range blackTiles {
		if i {
			countBlack++
		} else {
			countWhite++
		}
	}
	return countBlack, blackTiles
}

func countAdjacent(blackTiles map[coord]bool, currentCoord coord) int {
	count := 0

	if cur, ok := blackTiles[coord{x: currentCoord.x + 2, y: currentCoord.y}]; ok && cur {
		count++
	}
	if cur, ok := blackTiles[coord{x: currentCoord.x - 2, y: currentCoord.y}]; ok && cur {
		count++
	}
	if cur, ok := blackTiles[coord{x: currentCoord.x + 1, y: currentCoord.y + 1}]; ok && cur {
		count++
	}
	if cur, ok := blackTiles[coord{x: currentCoord.x + 1, y: currentCoord.y - 1}]; ok && cur {
		count++
	}
	if cur, ok := blackTiles[coord{x: currentCoord.x - 1, y: currentCoord.y + 1}]; ok && cur {
		count++
	}
	if cur, ok := blackTiles[coord{x: currentCoord.x - 1, y: currentCoord.y - 1}]; ok && cur {
		count++
	}

	return count
}

func gameOfLife(blackTiles map[coord]bool) int {
	for round := 0; round < 100; round++ {
		var minX, maxX, minY, maxY int
		for i := range blackTiles {
			if i.x < minX {
				minX = i.x
			}
			if i.x > maxX {
				maxX = i.x
			}
			if i.y < minY {
				minY = i.y
			}
			if i.y > maxY {
				maxY = i.y
			}
		}
		minX -= 2
		maxX += 2
		minY--
		maxY++

		c := 0
		for _, i := range blackTiles {
			if i {
				c++
			}
		}
		// fmt.Println("Round", round, " has:", c)

		newTiles := map[coord]bool{}
		for i := minX; i <= maxX; i++ {
			for j := minY; j <= maxY; j++ {
				currentCoord := coord{x: i, y: j}
				num := countAdjacent(blackTiles, currentCoord)

				if cur, ok := blackTiles[currentCoord]; ok && cur {
					if num == 0 {
						newTiles[currentCoord] = false
					} else if num > 2 {
						newTiles[currentCoord] = false
					} else {
						newTiles[currentCoord] = true
					}
				} else {
					if num == 2 {
						newTiles[currentCoord] = true
					}
				}
			}
		}
		blackTiles = newTiles
	}

	count := 0
	for _, i := range blackTiles {
		if i {
			count++
		}
	}
	return count
}

func main() {
	file := utils.ReadFile(os.Args[1])
	re := regexp.MustCompile(`(e|se|ne|w|sw|nw)`)

	paths := [][]string{}
	for _, line := range file {
		p := re.FindAll([]byte(line), -1)
		path := []string{}
		for _, b := range p {
			path = append(path, string(b))
		}
		paths = append(paths, path)
	}

	part1, blackTiles := runPaths(paths)
	fmt.Println("Part 1:", part1)

	part2 := gameOfLife(blackTiles)
	fmt.Println("Part 2:", part2)
}
