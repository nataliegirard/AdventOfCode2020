package main

import (
	"advent/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countActive(cubes map[string]string, index string) int {
	active := 0

	coord := strings.Split(index, ",")
	x, _ := strconv.Atoi(coord[0])
	y, _ := strconv.Atoi(coord[1])
	z, _ := strconv.Atoi(coord[2])

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			for k := z - 1; k <= z+1; k++ {
				if i == x && j == y && k == z {
					continue
				}
				lookup := fmt.Sprintf("%d,%d,%d", i, j, k)
				if val, ok := cubes[lookup]; ok {
					if val == "#" {
						active++
					}
				}
			}
		}
	}

	return active
}

func countActive4D(cubes map[string]string, index string) int {
	active := 0

	coord := strings.Split(index, ",")
	x, _ := strconv.Atoi(coord[0])
	y, _ := strconv.Atoi(coord[1])
	z, _ := strconv.Atoi(coord[2])
	w, _ := strconv.Atoi(coord[3])

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			for k := z - 1; k <= z+1; k++ {
				for l := w - 1; l <= w+1; l++ {
					if i == x && j == y && k == z && l == w {
						continue
					}
					lookup := fmt.Sprintf("%d,%d,%d,%d", i, j, k, l)
					if val, ok := cubes[lookup]; ok {
						if val == "#" {
							active++
						}
					}
				}
			}
		}
	}

	return active
}

func iterate(cubes map[string]string) map[string]string {
	newCubes := make(map[string]string)

	var minX, maxX, minY, maxY, minZ, maxZ int

	for k := range cubes {
		coord := strings.Split(k, ",")
		x, _ := strconv.Atoi(coord[0])
		y, _ := strconv.Atoi(coord[1])
		z, _ := strconv.Atoi(coord[2])
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
		if z < minZ {
			minZ = z
		}
		if z > maxZ {
			maxZ = z
		}
	}

	for i := minX - 1; i <= maxX+1; i++ {
		for j := minY - 1; j <= maxY+1; j++ {
			for k := minZ - 1; k <= maxZ+1; k++ {
				lookup := fmt.Sprintf("%d,%d,%d", i, j, k)
				num := countActive(cubes, lookup)

				val := "."
				if _, ok := cubes[lookup]; ok {
					val = "#"
				}

				if val == "." && num == 3 {
					newCubes[lookup] = "#"
				} else if val == "#" && (num == 2 || num == 3) {
					newCubes[lookup] = "#"
				}
			}
		}
	}

	return newCubes
}

func iterate4D(cubes map[string]string) map[string]string {
	newCubes := make(map[string]string)

	var minX, maxX, minY, maxY, minZ, maxZ, minW, maxW int

	for k := range cubes {
		coord := strings.Split(k, ",")
		x, _ := strconv.Atoi(coord[0])
		y, _ := strconv.Atoi(coord[1])
		z, _ := strconv.Atoi(coord[2])
		w, _ := strconv.Atoi(coord[3])
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
		if z < minZ {
			minZ = z
		}
		if z > maxZ {
			maxZ = z
		}
		if w < minW {
			minW = w
		}
		if w > maxW {
			maxW = w
		}
	}

	for i := minX - 1; i <= maxX+1; i++ {
		for j := minY - 1; j <= maxY+1; j++ {
			for k := minZ - 1; k <= maxZ+1; k++ {
				for l := minW - 1; l <= maxW+1; l++ {
					lookup := fmt.Sprintf("%d,%d,%d,%d", i, j, k, l)
					num := countActive4D(cubes, lookup)

					val := "."
					if _, ok := cubes[lookup]; ok {
						val = "#"
					}

					if val == "." && num == 3 {
						newCubes[lookup] = "#"
					} else if val == "#" && (num == 2 || num == 3) {
						newCubes[lookup] = "#"
					}
				}
			}
		}
	}

	return newCubes
}

func main() {
	file := utils.ReadFile(os.Args[1])

	cubes := make(map[string]string)
	for i, line := range file {
		items := strings.Split(line, "")
		for j, item := range items {
			if item == "#" {
				index := fmt.Sprintf("%d,%d,0", i, j)
				cubes[index] = item
			}
		}
	}

	for i := 0; i < 6; i++ {
		cubes = iterate(cubes)
	}

	fmt.Println("Part 1:", len(cubes))

	cubes4D := make(map[string]string)
	for i, line := range file {
		items := strings.Split(line, "")
		for j, item := range items {
			if item == "#" {
				index := fmt.Sprintf("%d,%d,0,0", i, j)
				cubes4D[index] = item
			}
		}
	}

	for i := 0; i < 6; i++ {
		cubes4D = iterate4D(cubes4D)
	}

	fmt.Println("Part 2:", len(cubes4D))
}
