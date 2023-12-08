package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var PartFlag = flag.Int("p", 1, "Which part should I run, 1 or 2")

func usage() {
	log.Fatal("usage: x -p N INPUT\n")
}

func main() {
	flag.Usage = usage
	if len(os.Args) < 2 {
		flag.Usage()
	}
	flag.Parse()
	files := flag.Args()
	switch *PartFlag {
	case 1:
		fmt.Printf("Result: %d\n", Part1(files[0]))
	case 2:
		fmt.Printf("Result: %d\n", Part2(files[0]))
	default:
		flag.Usage()
	}

}

func Part1(path string) int {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("opening input file: %v", err)
	}

	result := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		var winNums []int
		sep := strings.IndexRune(s.Text(), '|')
		start := strings.IndexRune(s.Text(), ':')
		towin := strings.Fields(s.Text()[start+1 : sep-1])
		got := strings.Fields(s.Text()[sep+1:])
		for _, w := range towin {
			for _, g := range got {
				wi, err := strconv.Atoi(w)
				if err != nil {
					log.Fatalf("converting %s into int: %v\n", w, err)
				}
				gi, _ := strconv.Atoi(g)
				if err != nil {
					log.Fatalf("converting %s into int: %v\n", g, err)
				}
				if wi == gi {
					winNums = append(winNums, wi)
				}
			}
		}
		if len(winNums) >= 1 {
			result = result + int(math.Pow(2, float64(len(winNums)-1)))
		}
	}
	return result
}

type Card struct {
	id  int
	win []int
}

func Part2(path string) int {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("opening input file: %v", err)
	}

	var pile []*Card
	var lineno int
	s := bufio.NewScanner(f)
	for s.Scan() {
		lineno++
		var winNums []int
		sep := strings.IndexRune(s.Text(), '|')
		start := strings.IndexRune(s.Text(), ':')
		towin := strings.Fields(s.Text()[start+1 : sep-1])
		got := strings.Fields(s.Text()[sep+1:])
		for _, w := range towin {
			for _, g := range got {
				wi, err := strconv.Atoi(w)
				if err != nil {
					log.Fatalf("converting %s into int: %v\n", w, err)
				}
				gi, _ := strconv.Atoi(g)
				if err != nil {
					log.Fatalf("converting %s into int: %v\n", g, err)
				}
				if wi == gi {
					winNums = append(winNums, wi)
				}
			}
		}
		pile = append(pile, &Card{lineno, winNums})
	}
	for i := 0; i < len(pile); i++ {
		for j := 0; j < len(pile[i].win); j++ {
			pile = append(pile, pile[pile[i].id+j])
		}
	}
	return len(pile)
}
