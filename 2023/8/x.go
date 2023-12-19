package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/exp/slices"
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

type Direction int

const (
	Left  Direction = 0
	Right           = 1
)

func (d Direction) Index() int { return int(d) }

type Position struct {
	name   string
	routes [2]string
}

type Mapp map[string]Position

func NewMapp() Mapp { return make(map[string]Position) }

func (m Mapp) GoTo(from Position, d Direction) Position {
	return m[m[from.name].routes[d.Index()]]
}

func parse(input []byte) Position {
	p := Position{}

	p.name = string(input[:3])
	p.routes = [2]string{string(input[7:10]), string(input[12:15])}

	return p
}

func Part1(path string) int {

	var (
		Walk = func(m Mapp, d []Direction) int {
			curr := m["AAA"]
			steps := 0
			i := 0
			for ; curr.name != "ZZZ" && i < len(d); steps++ {
				i = steps % len(d)
				curr = m.GoTo(curr, d[i])
			}
			return steps
		}

		input      = NewMapp()
		directions []Direction
	)
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("opening file %s; %v\n", path, err)
	}
	s := bufio.NewScanner(f)
	s.Scan()

	// scan list of directions
	for _, r := range s.Text() {
		switch r {
		case 'L':
			directions = append(directions, Left)
		case 'R':
			directions = append(directions, Right)
		default:
			log.Fatalf("Got '%c' on directions list\n", r)
		}
	}

	s.Scan() // scan an empty line

	for s.Scan() {
		pos := parse(s.Bytes())
		input[pos.name] = pos
	}

	return Walk(input, directions)
	// return 0
}

func Part2(path string) int {
	var (
		CalcLCM = func(values []int) int {

			x0 := slices.Clone(values)
			for {
				// are we done?
				if len(x0) == 1 {
					return values[0]
				}
				for i := 1; i < len(values); i++ {
					if values[i] != values[i-1] {
						break
					}
					if i == len(values)-1 {
						return values[0]
					}
				}

				// increase minimum value
				min := -1
				minidx := -1
				for i, v := range values {
					if min == -1 {
						min = v
						minidx = i
					}
					if v < min {
						min = v
						minidx = i
					}
				}
				values[minidx] += x0[minidx]
			}
		}

		Walk = func(m Mapp, dir []Direction) int {
			var currents []Position
			for _, pos := range m {
				if pos.name[2] == 'A' {
					currents = append(currents, pos)
				}
			}

			var lcm []int
			for _, curr := range currents {
				didx := 0
				steps := 0
				for curr.name[2] != 'Z' && didx < len(dir) {
					didx = steps % len(dir)
					curr = m.GoTo(curr, dir[didx])
					steps++
				}
				lcm = append(lcm, steps)
				fmt.Println(steps)
			}
			return CalcLCM(lcm)
		}

		input      = NewMapp()
		directions []Direction
	)

	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("opening file %s; %v\n", path, err)
	}
	s := bufio.NewScanner(f)
	s.Scan()

	// scan list of directions
	for _, r := range s.Text() {
		switch r {
		case 'L':
			directions = append(directions, Left)
		case 'R':
			directions = append(directions, Right)
		default:
			log.Fatalf("Got '%c' on directions list\n", r)
		}
	}

	s.Scan() // scan an empty line

	for s.Scan() {
		pos := parse(s.Bytes())
		input[pos.name] = pos
	}

	return Walk(input, directions)
	// return 0
}
