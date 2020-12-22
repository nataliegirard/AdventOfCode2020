package main

import (
	"advent/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://yourbasic.org/golang/compare-slices/
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func contains(list [][]int, item []int) bool {
	for _, i := range list {
		if equal(i, item) {
			return true
		}
	}

	return false
}

func determineScore(winner []int) int {
	sum := 0
	for i := 0; i < len(winner); i++ {
		sum = sum + winner[i]*(len(winner)-i)
	}
	return sum
}

func playWar(player1 []int, player2 []int) int {
	var winner []int
	for true {
		card1 := player1[0]
		card2 := player2[0]

		if card1 > card2 {
			player1 = append(player1[1:], card1, card2)
			player2 = player2[1:]
		} else {
			player1 = player1[1:]
			player2 = append(player2[1:], card2, card1)
		}
		// fmt.Println(card1, card2, player1, player2)

		if len(player1) == 0 {
			winner = player2
			break
		}
		if len(player2) == 0 {
			winner = player1
			break
		}
	}

	sum := determineScore(winner)

	return sum
}

func recursiveWar(player1 []int, player2 []int, game int) int {
	player1Hands := [][]int{}

	var winner []int
	for true {
		if contains(player1Hands, player1) {
			// fmt.Println("loop detected, player1 wins")
			if game != 1 {
				return 1
			}
			winner = player1
			break
		}
		player1Hands = append(player1Hands, player1)

		card1 := player1[0]
		card2 := player2[0]
		player1 = player1[1:]
		player2 = player2[1:]

		if len(player1) >= card1 && len(player2) >= card2 {
			// recurse to new game
			hand1 := []int{}
			hand2 := []int{}
			for i := 0; i < card1; i++ {
				hand1 = append(hand1, player1[i])
			}
			for i := 0; i < card2; i++ {
				hand2 = append(hand2, player2[i])
			}
			win := recursiveWar(hand1, hand2, game+1)
			// fmt.Println("\nWinner of recursive round: Player", win)
			if win == 1 {
				player1 = append(player1, card1, card2)
			} else {
				player2 = append(player2, card2, card1)
			}
		} else if card1 > card2 {
			player1 = append(player1, card1, card2)
		} else {
			player2 = append(player2, card2, card1)
		}
		// fmt.Println(card1, card2, player1, player2)

		if len(player1) == 0 {
			if game != 1 {
				return 2
			}
			winner = player2
			break
		}
		if len(player2) == 0 {
			if game != 1 {
				return 1
			}
			winner = player1
			break
		}
	}

	sum := determineScore(winner)
	return sum
}

func main() {
	file := utils.ReadFile(os.Args[1])

	playerHands := [2][]int{}
	player := 0
	for _, line := range file {
		if strings.Contains(line, "Player 1:") {
			player = 0
			continue
		}
		if strings.Contains(line, "Player 2:") {
			player = 1
			continue
		}
		if line == "" {
			continue
		}
		num, _ := strconv.Atoi(line)
		playerHands[player] = append(playerHands[player], num)
	}

	part1 := playWar(playerHands[0], playerHands[1])
	fmt.Println("Part 1:", part1)

	part2 := recursiveWar(playerHands[0], playerHands[1], 1)
	fmt.Println("Part 2:", part2)
}
