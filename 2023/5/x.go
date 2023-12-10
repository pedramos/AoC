package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
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

type Range struct {
	p0, p1, v0 int
}

func NewRange(src, dst, delta int) Range {
	p0 := src
	p1 := src + delta
	v0 := dst
	return Range{p0, p1, v0}
}

type Mapp []Range

func (m Mapp) Value(input int) int {
	for _, r := range m {
		if input >= r.p0 && input <= r.p1 {
			return r.v0 + (input - r.p0)
		}
	}
	return input
}

type Converter []Mapp

func NewMap(s *bufio.Scanner, c Converter) Converter {
	var m Mapp
	for s.Scan() && len(s.Bytes()) > 0 {
		var (
			input    = strings.Fields(s.Text())
			dst, _   = strconv.Atoi(input[0])
			src, _   = strconv.Atoi(input[1])
			delta, _ = strconv.Atoi(input[2])
		)
		m = append(m, NewRange(src, dst, delta))
	}
	c = append(c, m)
	return c
}

func (c Converter) Convert(seeds ...int) []int {
	var result = make([]int, 0, len(seeds))
	for _, s := range seeds {
		for _, m := range c {
			s = m.Value(s)
		}
		result = append(result, s)
	}
	return result
}
func (c Converter) ConvertDebug(s int) {

	for _, m := range c {
		fmt.Printf("v=%d\n", s)
		s = m.Value(s)
	}
	fmt.Printf("final=%d\n", s)
}

func (c Converter) String() string {
	var sb strings.Builder
	for _, m := range c {
		sb.WriteString(fmt.Sprintf("%v\n", m))
	}
	return sb.String()
}

func Part1(path string) int {
	var (
		seeds []int
		c     Converter
	)

	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("opening file %s: %v", path, err)
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		input := strings.Fields(s.Text())
		switch {
		case len(input) == 0:
			continue
		case input[0] == "seeds:":
			for _, seed := range input[1:] {
				v, err := strconv.Atoi(seed)
				if err != nil {
					log.Fatalf("converting %s to int: %v", seed, err)
				}
				seeds = append(seeds, v)
			}
		case input[1] == "map:":
			c = NewMap(s, c)
		default:
		}

	}
	result := -1
	for _, s := range c.Convert(seeds...) {
		if s < result || result == -1 {
			result = s
		}
	}
	return result
}

func Part2(path string) int {
	var (
		seeds []int
		c     Converter
	)

	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("opening file %s: %v", path, err)
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		input := strings.Fields(s.Text())
		switch {
		case len(input) == 0:
			continue
		case input[0] == "seeds:":
			for _, seed := range input[1:] {
				v, err := strconv.Atoi(seed)
				if err != nil {
					log.Fatalf("converting %s to int: %v", seed, err)
				}
				seeds = append(seeds, v)
			}
		case input[1] == "map:":
			c = NewMap(s, c)
		default:
		}

	}
	seeds = expandSeeds(seeds)
	result := -1
	for _, s := range c.Convert(seeds...) {
		if s < result || result == -1 {
			result = s
		}
	}
	return result
}

func expandSeeds(seeds []int) []int {
	var s []int
	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j < seeds[i+1]+seeds[i]; j++ {
			s = append(s, j)
		}
	}
	return s
}
