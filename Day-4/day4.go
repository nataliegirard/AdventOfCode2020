package main

import (
	"advent/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func validatePassportSimple(paper map[string]string) bool {
	count := 0
	_, byr := paper["byr"]
	if byr {
		count++
	}
	_, iyr := paper["iyr"]
	if iyr {
		count++
	}
	_, eyr := paper["eyr"]
	if eyr {
		count++
	}
	_, hgt := paper["hgt"]
	if hgt {
		count++
	}
	_, hcl := paper["hcl"]
	if hcl {
		count++
	}
	_, ecl := paper["ecl"]
	if ecl {
		count++
	}
	_, pid := paper["pid"]
	if pid {
		count++
	}
	return count == 7
}

func validatePassportComplex(paper map[string]string) bool {
	count := 0
	byr, _ := paper["byr"]
	byrYear, _ := strconv.Atoi(byr)
	if byrYear >= 1920 && byrYear <= 2002 { // four digits; at least 1920 and at most 2002
		count++
	}

	iyr, _ := paper["iyr"]
	iyrYear, _ := strconv.Atoi(iyr)
	if iyrYear >= 2010 && iyrYear <= 2020 { // four digits; at least 2010 and at most 2020
		count++
	}

	eyr, _ := paper["eyr"]
	eyrYear, _ := strconv.Atoi(eyr)
	if eyrYear >= 2020 && eyrYear <= 2030 { // four digits; at least 2020 and at most 2030
		count++
	}

	/* a number followed by either cm or in:
	If cm, the number must be at least 150 and at most 193.
	If in, the number must be at least 59 and at most 76
	*/
	hgt, _ := paper["hgt"]
	hgtRE := regexp.MustCompile(`^([0-9]+)(cm|in)$`)
	hgtResult := hgtRE.FindSubmatch([]byte(hgt))
	if len(hgtResult) != 0 {
		hgtValue, _ := strconv.Atoi(string(hgtResult[1]))
		if string(hgtResult[2]) == "cm" && hgtValue >= 150 && hgtValue <= 193 {
			count++
		} else if string(hgtResult[2]) == "in" && hgtValue >= 59 && hgtValue <= 76 {
			count++
		}
	}

	hcl, _ := paper["hcl"]
	hclRE := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	hclResult := hclRE.Match([]byte(hcl))
	if hclResult { // a # followed by exactly six characters 0-9 or a-f
		count++
	}

	ecl, _ := paper["ecl"]
	eclRE := regexp.MustCompile(`^amb|blu|brn|gry|grn|hzl|oth$`)
	eclResult := eclRE.Match([]byte(ecl))
	if eclResult { // exactly one of: amb blu brn gry grn hzl oth
		count++
	}

	pid, _ := paper["pid"]
	pidRE := regexp.MustCompile(`^[0-9]{9}$`)
	pidResult := pidRE.Match([]byte(pid))
	if pidResult { // a nine-digit number, including leading zeroes
		count++
	}

	return count == 7
}

func main() {
	file := utils.ReadFile(os.Args[1])

	countSimple := 0
	countComplex := 0
	paper := make(map[string]string)

	for _, line := range file {
		if line == "" {
			valid := validatePassportSimple(paper)
			if valid {
				countSimple++
				valid = validatePassportComplex(paper)

				if valid {
					countComplex++
				}
			}

			paper = make(map[string]string)
			continue
		}
		s := strings.Fields(line)
		for _, entry := range s {
			field := strings.Split(entry, ":")
			paper[field[0]] = field[1]
		}
	}
	valid := validatePassportSimple(paper)

	if valid {
		countSimple++
		valid = validatePassportComplex(paper)

		if valid {
			countComplex++
		}
	}
	fmt.Println("Part 1:", countSimple)
	fmt.Println("Part 2:", countComplex)
}
