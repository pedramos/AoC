package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

var converter = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"zero":  0,
}

func Calc(r io.Reader) (total int, values []int) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		var (
			b        = s.Bytes()
			b0       = 0
			b1       = 1
			lineNums []int
		)
		for b1 <= len(b) {
			//log.Println(string(b[b0:b1]))
			for k, v := range converter {
				if string(b[b0:b1]) == k {
					b0++
					lineNums = append(lineNums, v)
				}
				if r, _ := utf8.DecodeRune(b[b0:b1]); unicode.IsDigit(r) {
					b0 = b1
					lineNums = append(lineNums, int(r-'0'))
				}
			}
			b1++
			if b1 > len(b) || b1-b0 > 5 {
				b0++
				b1 = b0
			}
		}
		if len(lineNums) == 0 {
			log.Fatalf("Not enough numbers on this line: %s", s.Text())
		}
		finalLineNum := lineNums[0]*10 + lineNums[len(lineNums)-1]
		values = append(values, finalLineNum)
	}
	total = 0
	for i := range values {
		total += values[i]
	}
	return
}

func main() {

	// starters := []byte{'o', 't', 'f','s', 'e', 'n', 'z'}

	for _, path := range os.Args[1:] {
		f, err := os.Open(path)
		if err != nil {
			log.Fatal("Failed to open file\n")
		}
		total, values := Calc(f)
		fmt.Println(values)
		fmt.Println(total)
	}

}
