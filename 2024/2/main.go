package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	BadResult int = iota
	IncreasingSlope
	DecreasingSlope
)

func slope(x []string) int {
	x0, _ := strconv.Atoi(x[0])
	x1, _ := strconv.Atoi(x[1])
	delta := x1 - x0

	var s int
	if delta > 0 && delta <= 3 {
		s = IncreasingSlope
	}
	if delta < 0 && delta >= -3 {
		s = DecreasingSlope
	}
	if s == BadResult {
		return s
	}

	for i := range x {
		if i == len(x)-1 {
			break
		}
		x0, _ := strconv.Atoi(x[i])
		x1, _ := strconv.Atoi(x[i+1])
		delta := x1 - x0
		if delta < 0 && s == IncreasingSlope ||
			delta > 0 && s == DecreasingSlope ||
			s == IncreasingSlope && delta > 3 ||
			s == DecreasingSlope && delta < -3 ||
			delta == 0 {
			return BadResult
		}
	}
	return s
}

func solve1(args []string) {
	f, err := os.Open(args[1])
	if err != nil {
		log.Fatal(err)
	}
	safe := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		l := strings.Split(s.Text(), " ")
		// fmt.Println(l)
		s := slope(l)
		if s != BadResult {
			safe++
			// fmt.Println(s)
			continue
		}
		// fmt.Println(s)
	}
	fmt.Println(safe)
}

func solve2(args []string) {
	f, err := os.Open(args[1])
	if err != nil {
		log.Fatal(err)
	}

	safe := 0

	s := bufio.NewScanner(f)

	for s.Scan() {
		l := strings.Split(s.Text(), " ")
		// fmt.Println(l)
		s := slope(l)
		if s != BadResult {
			safe++
			// fmt.Println(s)
			continue
		}
		for i := range l {
			a := slices.Clone(l)
			a = append(a[:i], a[i+1:]...)
			// fmt.Printf("	%d\n", i)
			// fmt.Printf("	%v\n", a)
			s = slope(a)
			if s != BadResult {
				safe++
				break
			}
		}
		// fmt.Println(s)
	}
	fmt.Println(safe)
}

func main() {
	solve1(os.Args)
	solve2(os.Args)
}
