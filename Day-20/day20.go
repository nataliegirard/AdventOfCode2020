package main

import (
	"advent/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type tile struct {
	id      int
	image   [][]rune
	include bool
}

type match struct {
	first  tile
	second tile
	edge   string
}

type coord struct {
	x int
	y int
}

func printImage(image [][]rune) {
	for _, row := range image {
		fmt.Println(string(row))
	}
}

func printImageMap(imageMap map[coord]tile) {
	var minX, maxX, minY, maxY int
	for c := range imageMap {
		if c.x < minX {
			minX = c.x
		}
		if c.x > maxX {
			maxX = c.x
		}
		if c.y < minY {
			minY = c.y
		}
		if c.y > maxY {
			maxY = c.y
		}
	}

	for i := minX; i <= maxX; i++ {
		row := []int{}
		for j := minY; j <= maxY; j++ {
			c := coord{x: i, y: j}
			row = append(row, imageMap[c].id)
		}
		fmt.Println(row)
	}
}

func getEdge(t tile, side string) string {
	edge := []rune{}

	for i := 0; i < len(t.image); i++ {
		for j := 0; j < len(t.image[i]); j++ {
			if i == 0 && side == "top" {
				edge = append(edge, t.image[i][j])
			}

			if i == len(t.image)-1 && side == "bottom" {
				edge = append(edge, t.image[i][j])
			}

			if j == 0 && side == "left" {
				edge = append(edge, t.image[i][j])
			}

			if j == len(t.image[i])-1 && side == "right" {
				edge = append(edge, t.image[i][j])
			}
		}
	}

	return string(edge)
}

func rotate90(oldTile tile) tile {
	newTile := oldTile
	newTile.image = [][]rune{}
	for i := 0; i < len(oldTile.image); i++ {
		row := []rune{}
		for j := 0; j < len(oldTile.image[i]); j++ {
			row = append(row, ' ')
		}
		newTile.image = append(newTile.image, row)
	}

	for i := 0; i < len(oldTile.image); i++ {
		for j := 0; j < len(oldTile.image[i]); j++ {
			newTile.image[j][len(oldTile.image[i])-1-i] = oldTile.image[i][j]
		}
	}

	return newTile
}

func rotate180(oldTile tile) tile {
	newTile := oldTile
	newTile.image = [][]rune{}
	for i := 0; i < len(oldTile.image); i++ {
		row := []rune{}
		for j := 0; j < len(oldTile.image[i]); j++ {
			row = append(row, ' ')
		}
		newTile.image = append(newTile.image, row)
	}

	for i := 0; i < len(oldTile.image); i++ {
		for j := 0; j < len(oldTile.image[i]); j++ {
			newTile.image[len(oldTile.image[i])-1-i][len(oldTile.image)-1-j] = oldTile.image[i][j]
		}
	}

	return newTile
}

func rotate270(oldTile tile) tile {
	newTile := oldTile
	newTile.image = [][]rune{}
	for i := 0; i < len(oldTile.image); i++ {
		row := []rune{}
		for j := 0; j < len(oldTile.image[i]); j++ {
			row = append(row, ' ')
		}
		newTile.image = append(newTile.image, row)
	}

	for i := 0; i < len(oldTile.image); i++ {
		for j := 0; j < len(oldTile.image[i]); j++ {
			newTile.image[len(oldTile.image)-1-j][i] = oldTile.image[i][j]
		}
	}

	return newTile
}

func flipVertical(oldTile tile) tile {
	newTile := oldTile

	newTile.image = [][]rune{}
	for i := len(oldTile.image) - 1; i >= 0; i-- {
		newTile.image = append(newTile.image, oldTile.image[i])
	}

	return newTile
}

func flipHorizontal(oldTile tile) tile {
	newTile := oldTile
	newTile.image = [][]rune{}
	for i := 0; i < len(oldTile.image); i++ {
		row := []rune{}
		for j := len(oldTile.image[i]) - 1; j >= 0; j-- {
			row = append(row, oldTile.image[i][j])
		}
		newTile.image = append(newTile.image, row)
	}

	return newTile
}

func findMatches(tiles []tile, t tile, index int) []match {
	matches := []match{}
	for i, m := range tiles {
		if i == index {
			continue
		}

		for k := 0; k < 8; k++ {

			var tempTile tile
			switch k {
			case 0:
				tempTile = m
			case 1:
				tempTile = rotate90(m)
			case 2:
				tempTile = rotate180(m)
			case 3:
				tempTile = rotate270(m)
			case 4:
				tempTile = flipHorizontal(m)
			case 5:
				tempTile = flipVertical(m)
			case 6:
				tempTile = rotate90(flipHorizontal(m))
			case 7:
				tempTile = rotate270(flipHorizontal(m))
			}

			if getEdge(t, "top") == getEdge(tempTile, "bottom") {
				tiles[i] = tempTile
				matches = append(matches, match{first: t, second: tempTile, edge: "top"})
			}
			if getEdge(t, "bottom") == getEdge(tempTile, "top") {
				tiles[i] = tempTile
				matches = append(matches, match{first: t, second: tempTile, edge: "bottom"})
			}
			if getEdge(t, "left") == getEdge(tempTile, "right") {
				tiles[i] = tempTile
				matches = append(matches, match{first: t, second: tempTile, edge: "left"})
			}
			if getEdge(t, "right") == getEdge(tempTile, "left") {
				tiles[i] = tempTile
				matches = append(matches, match{first: t, second: tempTile, edge: "right"})
			}
		}
	}
	return matches
}

func matchEdges(tiles []tile) int {
	answer := 1
	for index, t := range tiles {
		matches := findMatches(tiles, t, index)

		if len(matches) == 2 {
			answer = answer * t.id
		}
	}
	return answer
}

func assembleImageMap(imageMap map[coord]tile, tiles []tile, first tile) map[coord]tile {
	var tileCoord coord
	for i, m := range imageMap {
		if m.id == first.id {
			tileCoord = i
			break
		}
	}
	matches := findMatches(tiles, first, first.id)
	var add match
	for _, m := range matches {
		exists := false
		for _, i := range imageMap {
			if i.id == m.second.id {
				exists = true
				break
			}
		}
		if exists {
			continue
		}

		add = m
		var newCoord coord
		if add.edge == "top" {
			newCoord = coord{x: tileCoord.x - 1, y: tileCoord.y}
		}
		if add.edge == "bottom" {
			newCoord = coord{x: tileCoord.x + 1, y: tileCoord.y}
		}
		if add.edge == "left" {
			newCoord = coord{x: tileCoord.x, y: tileCoord.y - 1}
		}
		if add.edge == "right" {
			newCoord = coord{x: tileCoord.x, y: tileCoord.y + 1}
		}

		// make sure the new tile is in the correct orientation
		for i, t := range tiles {
			if t.id == add.second.id {
				tiles[i] = add.second
				break
			}
		}

		// fmt.Println("Adding", tileCoord, first.id, newCoord, add.second.id, add.edge)
		// printImage(first.image)
		// fmt.Println("")
		// printImage(add.second.image)
		// fmt.Println("")

		imageMap[newCoord] = add.second
		imageMap = assembleImageMap(imageMap, tiles, add.second)
	}

	return imageMap
}

func removeEdges(image [][]rune) [][]rune {
	newImage := [][]rune{}

	for i := 1; i < len(image)-1; i++ {
		row := []rune{}
		for j := 1; j < len(image[i])-1; j++ {
			row = append(row, image[i][j])
		}
		newImage = append(newImage, row)
	}

	return newImage
}

func appendImage(imageRow [][]rune, image [][]rune) [][]rune {
	newImage := [][]rune{}

	for i := 0; i < len(image); i++ {
		var row []rune
		if len(imageRow) == 0 {
			row = []rune{}
		} else {
			row = imageRow[i]
		}

		row = append(row, image[i]...)
		newImage = append(newImage, row)
	}

	return newImage
}

func assembleImage(imageMap map[coord]tile) [][]rune {
	fullImage := [][]rune{}
	var minX, maxX, minY, maxY int
	for c := range imageMap {
		if c.x < minX {
			minX = c.x
		}
		if c.x > maxX {
			maxX = c.x
		}
		if c.y < minY {
			minY = c.y
		}
		if c.y > maxY {
			maxY = c.y
		}
	}

	for i := minX; i <= maxX; i++ {
		imageRow := [][]rune{}
		for j := minY; j <= maxY; j++ {
			c := coord{x: i, y: j}
			imageRow = appendImage(imageRow, removeEdges(imageMap[c].image))
		}
		fullImage = append(fullImage, imageRow...)
	}

	return fullImage
}

func replaceMonster(line []rune, index int, num int) []rune {
	newline := line

	switch num {
	case 1:
		newline[18+index] = 'O'
	case 2:
		newline[0+index] = 'O'
		newline[5+index] = 'O'
		newline[6+index] = 'O'
		newline[11+index] = 'O'
		newline[12+index] = 'O'
		newline[17+index] = 'O'
		newline[18+index] = 'O'
		newline[19+index] = 'O'
	case 3:
		newline[1+index] = 'O'
		newline[4+index] = 'O'
		newline[7+index] = 'O'
		newline[10+index] = 'O'
		newline[13+index] = 'O'
		newline[16+index] = 'O'
	}

	return newline
}

func findSeaMonsters(image [][]rune) int {
	monster1 := regexp.MustCompile(`.{18}#.`)
	monster2 := regexp.MustCompile(`#.{4}#{2}.{4}#{2}.{4}#{3}`)
	monster3 := regexp.MustCompile(`.#.{2}#.{2}#.{2}#.{2}#.{2}#.{3}`)
	finalImage := image

	for r := 0; r < 8; r++ {
		imageTile := tile{image: finalImage}
		var tempImage [][]rune
		switch r {
		case 0:
			tempImage = imageTile.image
		case 1:
			tempImage = rotate90(imageTile).image
		case 2:
			tempImage = rotate180(imageTile).image
		case 3:
			tempImage = rotate270(imageTile).image
		case 4:
			tempImage = flipHorizontal(imageTile).image
		case 5:
			tempImage = flipVertical(imageTile).image
		case 6:
			tempImage = flipHorizontal(rotate90(imageTile)).image
		case 7:
			tempImage = flipVertical(rotate90(imageTile)).image
		}
		for i := 0; i < len(tempImage)-2; i++ {
			for j := 0; j < len(tempImage[i])-20; j++ {
				line1 := string(tempImage[i][j : j+20])
				line2 := string(tempImage[i+1][j : j+20])
				line3 := string(tempImage[i+2][j : j+20])

				if monster1.Match([]byte(line1)) && monster2.Match([]byte(line2)) && monster3.Match([]byte(line3)) {
					// fmt.Println("Found monster", i, j, r, len(tempImage), len(tempImage[i]))
					newImage := [][]rune{}
					for l := 0; l < i; l++ {
						newImage = append(newImage, tempImage[l])
					}
					newImage = append(newImage, replaceMonster(tempImage[i], j, 1))
					newImage = append(newImage, replaceMonster(tempImage[i+1], j, 2))
					newImage = append(newImage, replaceMonster(tempImage[i+2], j, 3))
					for l := i + 3; l < len(tempImage); l++ {
						newImage = append(newImage, tempImage[l])
					}

					// unorient image
					newTile := tile{image: newImage}
					switch r {
					case 0:
						finalImage = newImage
					case 1:
						finalImage = rotate270(newTile).image
					case 2:
						finalImage = rotate180(newTile).image
					case 3:
						finalImage = rotate90(newTile).image
					case 4:
						finalImage = flipHorizontal(newTile).image
					case 5:
						finalImage = flipVertical(newTile).image
					case 6:
						finalImage = rotate270(flipHorizontal(newTile)).image
					case 7:
						finalImage = rotate270(flipVertical(newTile)).image
					}
				}
			}

		}
	}
	// printImage(finalImage)
	count := 0
	for i := 0; i < len(finalImage); i++ {
		for j := 0; j < len(finalImage[i]); j++ {
			if finalImage[i][j] == '#' {
				count++
			}
		}
	}

	return count
}

func main() {
	file := utils.ReadFile(os.Args[1])

	tiles := []tile{}
	var temp tile
	var image [][]rune
	for _, line := range file {
		if line == "" {
			temp.image = image
			tiles = append(tiles, temp)
			image = [][]rune{}
			continue
		}
		if strings.Contains(line, "Tile") {
			id, _ := strconv.Atoi(line[5:9])
			temp.id = id
		} else {
			image = append(image, []rune(line))
		}
	}
	temp.image = image
	tiles = append(tiles, temp)

	part1 := matchEdges(tiles)
	fmt.Println("Part 1:", part1)

	first := tiles[0]
	imageMap := make(map[coord]tile)
	imageMap[coord{}] = first
	imageMap = assembleImageMap(imageMap, tiles, first)
	// fmt.Println(len(imageMap))
	// printImageMap(imageMap)

	fullImage := assembleImage(imageMap)
	fullTile := tile{image: fullImage}
	fullTile.image = flipVertical(fullTile).image
	// printImage(fullTile.image)

	part2 := findSeaMonsters(fullTile.image)

	fmt.Println("Part 2:", part2)
}
