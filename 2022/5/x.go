package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type CrateStack [][]rune

func (cs *CrateStack) Pop(crate int) rune {
	s := (*cs)[crate]
	last := len(s) - 1

	r := s[last]
	s = s[:last]
	(*cs)[crate] = s
	return r
}

func (cs *CrateStack) Push(crate int, r rune) {
	//DEBUG fmt.Printf("Adding %c to crane %d\n", r, crate+1)
	(*cs)[crate] = append((*cs)[crate], r)
}

func (cs *CrateStack) BottomInsert(crate int, r rune) {
	//DEBUG fmt.Printf("Adding %c to crane %d\n", r, crate+1)
	tmp := (*cs)[crate]
	(*cs)[crate] = []rune{}
	(*cs)[crate] = append((*cs)[crate], r)
	(*cs)[crate] = append((*cs)[crate], tmp...)
}

func (cs CrateStack) Top() string {
	var b strings.Builder
	for i := range cs {
		b.WriteRune(cs[i][len(cs[i])-1])
	}
	return b.String()
}

func (cs CrateStack) Print() {
	for i := range cs {
		fmt.Printf("%d:	", i+1)
		for j := range cs[i] {
			fmt.Printf("%c ", cs[i][j])
		}
		fmt.Println("")
	}
}

func (cs *CrateStack) Move(src, dst, n int) {
	moving := []rune{}
	srcSz := len((*cs)[src])

	moving = append(moving, (*cs)[src][srcSz-n:]...)

	(*cs)[dst] = append((*cs)[dst], moving...)
	(*cs)[src] = (*cs)[src][:srcSz-n]
}

func parseCratePos(buf []byte, nth int) rune {
	i := nth*4 + 1
	return rune(buf[i])
}

func main() {

	f, _ := os.Open("p1")
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	nCrates := (len(scanner.Text()) + 1) / 4

	var stacks CrateStack
	stacks = make([][]rune, nCrates)

	line := scanner.Bytes()
	for i := 0; i < nCrates; i++ {
		if r := parseCratePos(line, i); unicode.IsLetter(r) {
			stacks.BottomInsert(i, r)
		}
	}
	for scanner.Scan() {
		line = scanner.Bytes()
		if rune(line[1]) == '1' {
			break
		}
		for i := 0; i < nCrates; i++ {
			if r := parseCratePos(line, i); unicode.IsLetter(r) {
				stacks.BottomInsert(i, r)
			}
		}
	}
	scanner.Scan()
	stacks.Print()
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		fmt.Println(line)

		n, _ := strconv.Atoi(line[1])
		src, _ := strconv.Atoi(line[3])
		dst, _ := strconv.Atoi(line[5])

		src--
		dst--

		// for i := 0; i < n; i++ {
		// 	part1(stacks, src, dst)
		// }
		part2(stacks, src, dst, n)
		stacks.Print()
	}
	fmt.Println(stacks.Top())

}

func part1(stacks CrateStack, src, dst int) {
	r := stacks.Pop(src)
	stacks.Push(dst, r)
}

func part2(stacks CrateStack, src, dst, n int) {
	stacks.Move(src, dst, n)
}