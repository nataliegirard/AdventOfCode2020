package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	description string
	contains    string
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

func findBags(rules []Bag, bagName string) map[string]int {
	bagsContaining := make(map[string]int)
	for _, rule := range rules {
		if rule.contains == "no other bags" {
			continue
		}

		bagsWithin := strings.Split(rule.contains, ", ")
		r2 := regexp.MustCompile(`([0-9]) ([a-z]+ [a-z]+) bag[s]?`)
		for _, subBag := range bagsWithin {
			result := r2.FindSubmatch([]byte(subBag))
			ref := string(result[2])
			if ref == bagName {
				bagsContaining[rule.description] = 1
				newBags := findBags(rules, rule.description)
				for k, v := range newBags {
					bagsContaining[k] = v
				}
				break
			}
		}
	}
	return bagsContaining
}

func countBags(rules []Bag, bagName string) int {
	count := 0
	var bag Bag
	for _, b := range rules {
		if b.description == bagName {
			bag = b
			break
		}
	}

	if bag.contains == "no other bags" {
		return 0
	}

	bagsWithin := strings.Split(bag.contains, ", ")
	r2 := regexp.MustCompile(`([0-9]) ([a-z]+ [a-z]+) bag[s]?`)
	for _, subBag := range bagsWithin {
		result := r2.FindSubmatch([]byte(subBag))
		quantity, _ := strconv.Atoi(string(result[1]))
		name := string(result[2])
		count = count + quantity + quantity*countBags(rules, name)
	}

	return count
}

func main() {
	file := readFile(os.Args[1])

	rules := make([]Bag, 0)
	for _, line := range file {
		r1 := regexp.MustCompile(`^([a-z]+ [a-z]+) bags contain (.*).$`)
		breakdown := r1.FindSubmatch([]byte(line))

		newBag := Bag{
			description: string(breakdown[1]),
			contains:    string(breakdown[2]),
		}
		rules = append(rules, newBag)
	}

	bagsContaining := findBags(rules, "shiny gold")
	fmt.Println("Part 1:", len(bagsContaining))

	totalBags := countBags(rules, "shiny gold")
	fmt.Println("Part 2:", totalBags)
}
