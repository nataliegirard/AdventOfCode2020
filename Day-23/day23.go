package main

import (
	"advent/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	value int
	next  *node
}

func contains(list []int, item int) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

func (n *node) insert(newNode *node) {
	newNode.next = n.next
	n.next = newNode
}

func (n *node) remove() *node {
	removed := n.next
	n.next = removed.next
	return removed
}

func (n *node) getList() []int {
	current := n
	values := []int{}
	for true {
		values = append(values, current.value)
		current = current.next
		if current.value == n.value {
			break
		}
	}
	return values
}

func playGame(cups []int, turns int) map[int]*node {
	min := cups[0]
	max := cups[0]
	for _, c := range cups {
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}

	cupMap := map[int]*node{}
	for _, i := range cups {
		cupMap[i] = &node{value: i}
	}

	for index, i := range cups {
		cupMap[i].next = cupMap[cups[(index+1)%len(cups)]]
	}

	current := cupMap[cups[0]]
	for round := 0; round < turns; round++ {
		// fmt.Println("\nRound", round, "List:", current.getList())
		// fmt.Println("Current:", current.value)

		pickedUp := []*node{}
		picked := []int{}
		for i := 0; i < 3; i++ {
			pickedUp = append(pickedUp, current.remove())
			picked = append(picked, pickedUp[i].value)
		}
		// fmt.Println("Picked up:", picked)

		destinationValue := current.value - 1
		for {
			if destinationValue < min {
				destinationValue = max
			}
			if !contains(picked, destinationValue) {
				break
			}
			destinationValue--
		}
		// fmt.Println("Destination looking for:", destinationValue)

		destination := cupMap[destinationValue]
		// fmt.Println("Destination:", destination.value)

		for i := len(pickedUp) - 1; i >= 0; i-- {
			destination.insert(pickedUp[i])
		}
		current = current.next
	}
	// fmt.Println("final:", current.getList())
	return cupMap
}

func part1(cups []int) string {
	cupMap := playGame(cups, 100)
	answer := ""
	cup := cupMap[1].next
	for cup != cupMap[1] {
		answer = answer + fmt.Sprintf("%d", cup.value)
		cup = cup.next
	}

	return answer
}

func part2(cups []int) int {
	max := cups[0]
	for _, c := range cups {
		if c > max {
			max = c
		}
	}
	for i := 10; i <= 1000000; i++ {
		cups = append(cups, i)
	}

	cupMap := playGame(cups, 10000000)

	answer := cupMap[1].next.value * cupMap[1].next.next.value

	return answer
}

func main() {
	file := utils.ReadFile(os.Args[1])

	cupsInput := strings.Split(file[0], "")
	cups := []int{}
	for _, c := range cupsInput {
		cup, _ := strconv.Atoi(c)
		cups = append(cups, cup)
	}

	part1 := part1(cups)
	fmt.Println("Part 1:", part1)

	part2 := part2(cups)
	fmt.Println("Part 2:", part2)
}
