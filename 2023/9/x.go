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

func IsZero(input []int) bool {
	for _, i := range input {
		if i != 0 {
			return false
		}
	}
	return true
}

func Part1(path string) int {
	// var predictions []int
	var result int
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("opening file %s: %v", path, err)
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		var curr []int
		for _, num := range strings.Fields(s.Text()) {
			n, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("converting %s to int: %v", num, err)
			}
			curr = append(curr, n)
		}
		//-- expand line to find sequence zero --//
		var expanded [][]int
		expanded = append(expanded, make([]int, len(curr)))
		copy(expanded[0], curr)
		for seq := 0; !IsZero(expanded[len(expanded)-1]); seq++ {
			expanded = append(expanded, make([]int, len(expanded[seq])-1))
			for i := 0; i+1 < len(expanded[seq]); i++ {
				expanded[seq+1][i] = expanded[seq][i+1] - expanded[seq][i]
			}
		}
		// fmt.Println(expanded)

		//-- add last elements to find prediction --//

		for seq := len(expanded) - 2; seq > 0; seq-- {
			expanded[seq-1] = append(expanded[seq-1],
				expanded[seq][len(expanded[seq])-1]+expanded[seq-1][len(expanded[seq-1])-1])
		}
		// fmt.Println(expanded)
		result += expanded[0][len(expanded[0])-1]

	}

	return result

}

func Part2(path string) int {
	// var predictions []int
	var result int
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("opening file %s: %v", path, err)
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		var curr []int
		for _, num := range strings.Fields(s.Text()) {
			n, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("converting %s to int: %v", num, err)
			}
			curr = append(curr, n)
		}
		//-- expand line to find sequence zero --//
		var expanded [][]int
		expanded = append(expanded, make([]int, len(curr)))
		copy(expanded[0], curr)
		for seq := 0; !IsZero(expanded[len(expanded)-1]); seq++ {
			expanded = append(expanded, make([]int, len(expanded[seq])-1))
			for i := 0; i+1 < len(expanded[seq]); i++ {
				expanded[seq+1][i] = expanded[seq][i+1] - expanded[seq][i]
			}
		}
		// fmt.Println(expanded)

		//-- add last elements to find prediction --//

		for seq := len(expanded) - 2; seq > 0; seq-- {
			expanded[seq-1] = append([]int{expanded[seq-1][0] - expanded[seq][0]}, expanded[seq-1]...)
		}
		// fmt.Println(expanded)
		result += expanded[0][0]

	}

	return result

}
