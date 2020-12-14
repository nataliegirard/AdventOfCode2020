package main

import (
	"advent/utils"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type schedule struct {
	busNumber     int
	nextDeparture int
}

/* Chinese Remainder Theorem
Code taken from https://rosettacode.org/wiki/Chinese_remainder_theorem#Go
*/
var one = big.NewInt(1)

func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}

func main() {
	file := utils.ReadFile(os.Args[1])

	arrivalTime, _ := strconv.Atoi(file[0])
	buses := strings.Split(file[1], ",")
	busSchedule := []schedule{}
	nextDeparture := 100000000000
	nextBus := -1

	for _, num := range buses {
		if num == "x" {
			continue
		}
		n, _ := strconv.Atoi(num)
		s := schedule{}
		last := arrivalTime / n
		s.busNumber = n
		s.nextDeparture = (last + 1) * n
		busSchedule = append(busSchedule, s)
		if s.nextDeparture < nextDeparture {
			nextDeparture = s.nextDeparture
			nextBus = n
		}
	}
	fmt.Println("Part 1:", nextBus*(nextDeparture-arrivalTime))

	n := []*big.Int{}
	a := []*big.Int{}
	for i, num := range buses {
		if num == "x" {
			continue
		}
		number, _ := strconv.ParseInt(num, 10, 64)
		n = append(n, big.NewInt(number))
		a = append(a, big.NewInt(int64(-i)))
	}
	part2, _ := crt(a, n)
	fmt.Println("Part 2:", part2)
}
