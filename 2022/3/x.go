package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {

	result := 0

	f, err := os.Open("input_p2")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)

	result = part2(s)

	fmt.Printf("Total: %d\n", result)

}

func part1(s *bufio.Scanner) int {
	result := 0
	for s.Scan() {
		line := s.Bytes()
		viewed := map[rune]int{}
		badges := []rune{}
		inHalf := 1
		for i := range line {
			if i == len(line)/2 {
				inHalf = 2
			}
			item := rune(line[i])
			switch viewed[item] {
			case 0:
				viewed[item] = inHalf
			case inHalf:
				continue
			default:
				badges = append(badges, item)
				viewed[item] = inHalf
			}

		}
		// fmt.Printf("%c\n", badges)
		for _, r := range badges {
			if r >= 'a' && r <= 'z' {
				r = r - 'a' + 1
				result += int(r)
			} else if r >= 'A' && r <= 'Z' {
				r = r - 'A' + 27
				result += int(r)
			}
			// fmt.Println(int(r))
		}
	}
	return result
}

func part2(s *bufio.Scanner) int {

	// This Split is for part 2
	s.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.Count(data, []byte{'\n'}); i > 3 {
			elfGroup := bytes.Split(data, []byte{'\n'})[0:3]
			line := bytes.Join(elfGroup, []byte{'@'})
			return len(line) + 1, line, nil
		}
		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			return len(data), bytes.ReplaceAll(data, []byte{'\n'}, []byte{'@'}), nil
		}
		// Request more data.
		return 0, nil, nil
	})

	type elfid byte
	const (
		E1 elfid = 1 << iota
		E2
		E3
	)

	result := 0
	badges := []rune{}

	for s.Scan() {
		line := s.Bytes()
		viewed := map[rune]elfid{}
		elf := E1
		for i := range line {
			if line[i] == '@' {
				// fmt.Printf("Switch to elf: %b\n", elf<<1)
				elf = elf << 1
				continue
			}
			item := rune(line[i])

			if viewed[item] == 0 {
				viewed[item] = elfid(elf)
			} else if viewed[item]&elf == elf {
				continue
			} else {
				viewed[item] = elf | viewed[item]
				if viewed[item] == E1|E2|E3 {
					// fmt.Printf("found item in all: %c\n", item)
					badges = append(badges, item)
				}
			}

		}

	}
	for _, r := range badges {
		if r >= 'a' && r <= 'z' {
			prio := r - 'a' + 1
			result += int(prio)
			// fmt.Printf("%c : %d\n", r, prio)
		} else if r >= 'A' && r <= 'Z' {
			prio := r - 'A' + 27
			result += int(prio)
			// fmt.Printf("%c : %d\n", r, prio)
		}
	}
	return result
}