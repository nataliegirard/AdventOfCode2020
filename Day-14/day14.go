package main

import (
	"advent/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func applyMask(mask string, val int, version int) string {
	bin := strconv.FormatInt(int64(val), 2)
	value := string(bin)
	pos := len(value) - 1
	newVal := ""

	for i := len(mask) - 1; i >= 0; i-- {
		if version == 1 {
			if mask[i] == '1' {
				newVal = "1" + newVal
			} else if mask[i] == '0' {
				newVal = "0" + newVal
			} else if pos >= 0 {
				newVal = string(value[pos]) + newVal
			} else {
				newVal = "0" + newVal
			}
		} else {
			if mask[i] == '1' {
				newVal = "1" + newVal
			} else if mask[i] == 'X' {
				newVal = "X" + newVal
			} else if pos >= 0 {
				newVal = string(value[pos]) + newVal
			} else {
				newVal = "0" + newVal
			}
		}
		pos--
	}

	return newVal
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func saveValues(mem map[int64]int64, val string) map[int64]int64 {
	var newval string
	var newValue int64

	index := strings.Index(val, "X")
	if index > -1 {
		newval = replaceAtIndex(val, '0', index)
		mem = saveValues(mem, newval)

		newval = replaceAtIndex(val, '1', index)
		mem = saveValues(mem, newval)
		return mem
	}

	newValue, _ = strconv.ParseInt(val, 2, 64)
	mem[newValue] = newValue

	return mem
}

func main() {
	file := utils.ReadFile(os.Args[1])

	mem := make(map[int]int)
	addr2 := make(map[int64]int64)
	mem2 := make(map[int64]int64)
	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	for _, line := range file {
		if strings.Contains(line, "mask") {
			l := strings.Split(line, " = ")
			mask = l[1]
		} else {
			re := regexp.MustCompile("^mem\\[([0-9]+)\\] = ([0-9]+)$")
			result := re.FindSubmatch([]byte(line))
			addr, _ := strconv.Atoi(string(result[1]))
			value, _ := strconv.Atoi(string(result[2]))

			newVal := applyMask(mask, value, 1)
			newValue, _ := strconv.ParseInt(newVal, 2, 64)
			mem[addr] = int(newValue)

			newVal2 := applyMask(mask, addr, 2)
			addr2 = saveValues(addr2, newVal2)
			for val := range addr2 {
				mem2[val] = int64(value)
			}
			addr2 = make(map[int64]int64)
		}
	}

	sum := 0
	for _, val := range mem {
		sum = sum + val
	}
	fmt.Println("Part 1:", sum)

	var total int64
	total = 0
	for _, val := range mem2 {
		total = total + val
	}
	fmt.Println("Part 2:", total)
}
