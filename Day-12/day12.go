package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

// Position in a 2d cartesian grid with direction
type Position struct {
	x         int // East positive
	y         int // North positive
	direction int // E=0, S=1, W=2, N=3
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

func parseInstruction(ins string, position Position, waypoint Position, useWaypoint bool) (Position, Position) {
	newPosition := Position{x: position.x, y: position.y, direction: position.direction}
	newWaypoint := Position{x: waypoint.x, y: waypoint.y}

	re := regexp.MustCompile(`^(N|S|E|W|L|R|F)([0-9]+)$`)
	result := re.FindSubmatch([]byte(ins))
	val, _ := strconv.Atoi(string(result[2]))

	switch string(result[1]) {
	case "N":
		if useWaypoint {
			newWaypoint.y = newWaypoint.y + val
		} else {
			newPosition.y = newPosition.y + val
		}
		return newPosition, newWaypoint
	case "S":
		if useWaypoint {
			newWaypoint.y = newWaypoint.y - val
		} else {
			newPosition.y = newPosition.y - val
		}
		return newPosition, newWaypoint
	case "E":
		if useWaypoint {
			newWaypoint.x = newWaypoint.x + val
		} else {
			newPosition.x = newPosition.x + val
		}
		return newPosition, newWaypoint
	case "W":
		if useWaypoint {
			newWaypoint.x = newWaypoint.x - val
		} else {
			newPosition.x = newPosition.x - val
		}
		return newPosition, newWaypoint
	case "L":
		if !useWaypoint {
			t := newPosition.direction - (val / 90)
			if t < 0 {
				t = t + 4
			}
			// In Go, mod of negative numbers doesn't work
			newPosition.direction = (t) % 4
			return newPosition, newWaypoint
		}
		switch val {
		case 90:
			newWaypoint.x = waypoint.y * -1
			newWaypoint.y = waypoint.x
			return newPosition, newWaypoint
		case 180:
			newWaypoint.x = waypoint.x * -1
			newWaypoint.y = waypoint.y * -1
			return newPosition, newWaypoint
		case 270:
			newWaypoint.x = waypoint.y
			newWaypoint.y = waypoint.x * -1
			return newPosition, newWaypoint
		}
		return newPosition, newWaypoint
	case "R":
		if !useWaypoint {
			newPosition.direction = (newPosition.direction + (val / 90)) % 4
			return newPosition, newWaypoint
		}
		switch val {
		case 90:
			newWaypoint.x = waypoint.y
			newWaypoint.y = waypoint.x * -1
			return newPosition, newWaypoint
		case 180:
			newWaypoint.x = waypoint.x * -1
			newWaypoint.y = waypoint.y * -1
			return newPosition, newWaypoint
		case 270:
			newWaypoint.x = waypoint.y * -1
			newWaypoint.y = waypoint.x
			return newPosition, newWaypoint
		}
		return newPosition, newWaypoint
	case "F":
		break
	}

	// F case
	if useWaypoint {
		newPosition.x = position.x + val*waypoint.x
		newPosition.y = position.y + val*waypoint.y
	} else {
		switch newPosition.direction {
		case 0:
			newPosition.x = newPosition.x + val
			return newPosition, newWaypoint
		case 1:
			newPosition.y = newPosition.y - val
			return newPosition, newWaypoint
		case 2:
			newPosition.x = newPosition.x - val
			return newPosition, newWaypoint
		case 3:
			newPosition.y = newPosition.y + val
			return newPosition, newWaypoint
		}
	}
	return newPosition, newWaypoint
}

func main() {
	file := readFile(os.Args[1])

	position1 := Position{}
	position2 := Position{}
	waypoint := Position{x: 10, y: 1} // note: waypoint doesn't have a direction

	for _, line := range file {
		position1, _ = parseInstruction(line, position1, Position{}, false)
		position2, waypoint = parseInstruction(line, position2, waypoint, true)
	}
	fmt.Println("Part 1:", math.Abs(float64(position1.x))+math.Abs(float64(position1.y)))
	fmt.Println("Part 2:", math.Abs(float64(position2.x))+math.Abs(float64(position2.y)))
}
