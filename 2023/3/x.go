package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

var PartFlag = flag.Int("p", 1, "Which part should I run, 1 or 2")

func usage() {
	log.Fatal("usage: x -p N INPUT\n")
}

type Symbol struct {
	x, y int
	r    rune
}

type Part struct {
	x0, len int
	y       int
	id      int
}

func (p Part) String() string {
	return fmt.Sprintf("(id=%d x0=%d y=%d len=%d)", p.id, p.x0, p.y, p.len)
}

type Grid struct {
	linesize int
	grid     [][]*Part
	syms     []Symbol
	isInit   bool
}

func NewGrid(r io.Reader) (*Grid, error) {
	var (
		grid *Grid
		x    int
		y    int
	)
	s := bufio.NewScanner(r)
	for s.Scan() {
		grid = initGrid(grid, s.Bytes())
		grid.AddLine()
		for _, r := range s.Text() {
			switch {
			case r >= '0' && r <= '9':
				grid.AddPart(x, y, int(r)-'0')
			case r == '.':
			default:
				grid.AddSymbol(x, y, r)
			}
			x++
		}
		y++
		x = 0
	}
	return grid, nil
}

func initGrid(g *Grid, line []byte) *Grid {
	if g == nil {
		g = &Grid{}
	}
	if g.isInit == false {
		g.linesize = utf8.RuneCount(line)
		g.isInit = true
		return g
	}
	return g
}

func (g *Grid) AddLine() { g.grid = append(g.grid, make([]*Part, g.linesize)) }

func (g *Grid) AddSymbol(x, y int, r rune) {
	g.syms = append(g.syms, Symbol{x, y, r})
}

func (g *Grid) AddPart(x, y, digit int) {

	if x != 0 && g.grid[y][x-1] != nil {
		g.grid[y][x] = g.grid[y][x-1]
		g.grid[y][x].len++
		g.grid[y][x].id = g.grid[y][x].id*10 + digit
	} else {
		g.grid[y][x] = &Part{
			x0:  x,
			y:   y,
			len: 1,
			id:  digit,
		}
	}
}

func (g Grid) Part(x, y int) *Part {
	return g.grid[y][x]
}

func (g Grid) String() string {
	var sb strings.Builder

	for _, line := range g.grid {
		for _, part := range line {
			sb.WriteString(fmt.Sprintf("%s ", part))
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

type Register struct {
	index map[*Part]int // just to make easy to find duplicates
	parts []*Part
}

func (r *Register) Add(p *Part) {
	if _, found := r.index[p]; !found {
		r.index[p] = len(r.parts)
		r.parts = append(r.parts, p)
	}
}

func FindNeighbours(g *Grid) Register {
	var reg Register = Register{
		index: make(map[*Part]int, 100),
		parts: make([]*Part, 0, 100),
	}

	for _, sym := range g.syms {
		xdelta := []int{sym.x - 1, sym.x, sym.x + 1}
		ydelta := []int{sym.y - 1, sym.y, sym.y + 1}

		for _, x := range xdelta {
			for _, y := range ydelta {
				p := g.Part(x, y)
				if p == nil {
					continue
				}
				reg.Add(p)
			}
		}
	}
	return reg
}

func FindGears(g *Grid) []int {
	var ratios []int
	for _, sym := range g.syms {
		if sym.r != '*' {
			continue
		}

		xdelta := []int{sym.x - 1, sym.x, sym.x + 1}
		ydelta := []int{sym.y - 1, sym.y, sym.y + 1}

		var reg Register = Register{
			index: make(map[*Part]int, 8),
			parts: make([]*Part, 0, 8),
		}

		for _, x := range xdelta {
			for _, y := range ydelta {
				p := g.Part(x, y)
				if p == nil {
					continue
				}
				reg.Add(p)
			}
		}
		if len(reg.parts) == 2 {
			ratios = append(ratios, reg.parts[0].id*reg.parts[1].id)
		}

	}
	return ratios
}

func (r Register) String() string {
	return fmt.Sprintf("%v", r.parts)
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
		log.Fatalf("Could not open file %s: %v", path, err)
	}
	grid, err := NewGrid(f)
	if err != nil {
		log.Fatalf("Error parsing grid: %v\n", err)
	}
	// fmt.Println(grid)
	reg := FindNeighbours(grid)
	// fmt.Println(reg)

	count := 0
	for _, p := range reg.parts {
		count += p.id
	}
	// fmt.Printf("Total: %d\n", count)
	return count

}

func Part2(path string) int {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open file %s: %v", path, err)
	}
	grid, err := NewGrid(f)
	if err != nil {
		log.Fatalf("Error parsing grid: %v\n", err)
	}
	// fmt.Println(grid)
	ratios := FindGears(grid)
	// fmt.Println(reg)

	count := 0
	for _, r := range ratios {
		count += r
	}
	// fmt.Printf("Total: %d\n", count)
	return count
}
