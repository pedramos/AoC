package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
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

type Hand struct {
	cards    [5]int
	rank     int
	bid      int
	cardsStr string // for troubleshooting
}

type HandRankCheck func(h Hand) bool

func NewHand(CardToValue map[rune]int, handRankers []HandRankCheck, cards string, bid int) Hand {
	h := Hand{bid: bid, cardsStr: cards}
	for i, c := range cards {
		h.cards[i] = CardToValue[c]
	}
	h.Classify(handRankers)
	return h
}

func (h *Hand) Classify(handRankers []HandRankCheck) {
	for n, isRank := range handRankers {
		if isRank(*h) {
			h.rank = n
			return
		}
	}
}

// Untie returns a pointer to the winner
func UntieFirstWon(h1 Hand, h2 Hand) bool {
	for i := range h1.cards {
		switch {
		case h1.cards[i] == h2.cards[i]:
			continue
		case h1.cards[i] > h2.cards[i]:
			return false
		case h1.cards[i] < h2.cards[i]:
			return true
		}
	}
	return true
}

type Hands []Hand

func (h Hands) Len() int      { return len(h) }
func (h Hands) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h Hands) Less(i, j int) bool {
	switch {
	case h[i].rank > h[j].rank:
		return true
	case h[i].rank < h[j].rank:
		return false
	case h[i].rank == h[j].rank:
		return UntieFirstWon(h[i], h[j])
	}
	return false // it should be unreacheble
}

func Part1(path string) int {

	var CardValue map[rune]int = map[rune]int{
		'2': 0,
		'3': 1,
		'4': 2,
		'5': 3,
		'6': 4,
		'7': 5,
		'8': 6,
		'9': 7,
		'T': 8,
		'J': 9,
		'Q': 10,
		'K': 11,
		'A': 12,
	}

	FiveOfAKind := func(h Hand) bool {
		cards := h.cards
		for i := 1; i < len(cards); i++ {
			if cards[i] != cards[i-1] {
				return false
			}
		}
		return true
	}

	FourOfAKind := func(h Hand) bool {
		found := make(map[int]int, 4)
		for _, i := range h.cards {
			found[i]++
		}
		for _, f := range found {
			if f == 4 {
				return true
			}
		}
		return false
	}

	FullHouse := func(h Hand) bool {
		found := make(map[int]int, 4)
		for _, i := range h.cards {
			found[i]++
		}
		found3 := false
		found2 := false
		for _, f := range found {
			if f == 3 {
				found3 = true
			}
			if f == 2 {
				found2 = true
			}
		}
		if found3 && found2 {
			return true
		}
		return false
	}

	ThreeOfAKind := func(h Hand) bool {
		found := make(map[int]int, 4)
		for _, i := range h.cards {
			found[i]++
		}
		for _, f := range found {
			if f == 3 {
				return true
			}
		}
		return false
	}

	TwoPair := func(h Hand) bool {
		found := make(map[int]int, 4)
		pairs := 0
		for _, i := range h.cards {
			found[i]++
			if len(found) > 3 {
				return false
			}
		}

		for _, f := range found {
			if f == 2 {
				pairs++
			}
		}

		if pairs == 2 {
			return true
		}
		return false
	}

	OnePair := func(h Hand) bool {
		found := make(map[int]int, 4)
		pairs := 0
		for _, i := range h.cards {
			found[i]++
		}
		for _, f := range found {
			if f == 2 {
				pairs++
			}
		}
		if pairs == 1 {
			return true
		}
		return false
	}

	HighCard := func(h Hand) bool { return true }

	// hand types sorted by value, highest first
	var HandTypeRank []HandRankCheck = []HandRankCheck{
		FiveOfAKind,
		FourOfAKind,
		FullHouse,
		ThreeOfAKind,
		TwoPair,
		OnePair,
		HighCard,
	}

	var plays []Hand

	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("opening file %s, %v", path, err)
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.Fields(s.Text())
		b, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatalf("converting %s to int: %v", line[1], err)
		}
		plays = append(plays, NewHand(CardValue, HandTypeRank, line[0], b))
	}
	// 	for i := range plays {
	// 		fmt.Printf("rank=%d bid=%d cards=%s\n", plays[i].rank, plays[i].bid, plays[i].cardsStr)
	// 	}
	sort.Sort(Hands(plays))
	result := 0
	for i := range plays {
		result += (i + 1) * plays[i].bid
		// fmt.Printf("rank=%d bid=%d\n", plays[i].rank, plays[i].bid)
	}
	return result
}

func Part2(path string) int {

	var CardValue map[rune]int = map[rune]int{
		'J': 0,
		'2': 1,
		'3': 2,
		'4': 3,
		'5': 4,
		'6': 5,
		'7': 6,
		'8': 7,
		'9': 8,
		'T': 9,
		'Q': 10,
		'K': 11,
		'A': 12,
	}

	FiveOfAKind := func(h Hand) bool {
		found := make(map[int]int, 4)
		for _, i := range h.cards {
			found[i]++
		}

		if found[0] >= 4 {
			return true
		}

		for c, f := range found {
			if c == 0 {
				continue
			}
			if f+found[0] >= 5 {
				return true
			}
		}
		return false
	}

	FourOfAKind := func(h Hand) bool {
		found := make(map[int]int, 4)
		for _, i := range h.cards {
			found[i]++
		}
		if found[0] >= 3 {
			return true
		}
		for c, f := range found {
			if c == 0 {
				continue
			}
			if f+found[0] >= 4 {
				return true
			}
		}
		return false
	}

	FullHouse := func(h Hand) bool {
		found := make(map[int]int, 4)
		for _, i := range h.cards {
			found[i]++
		}

		found2 := false
		found3 := false

		for c, f := range found {
			if c == 0 {
				continue
			}
			if f == 3 {
				found3 = true
			}
			if f == 2 && found2 == false {
				found2 = true
				continue
			}
			if f+found[0] >= 3 && found2 == true {
				found3 = true
			}
		}
		if found3 && found2 {
			return true
		}
		return false
	}

	ThreeOfAKind := func(h Hand) bool {
		found := make(map[int]int, 4)
		for _, i := range h.cards {
			found[i]++
		}
		if found[0] >= 2 {
			return true
		}
		for c, f := range found {
			if c == 0 {
				continue
			}
			if f+found[0] >= 3 {
				return true
			}
		}
		return false
	}

	TwoPair := func(h Hand) bool {
		found := make(map[int]int, 4)
		pairs := 0
		for _, i := range h.cards {
			found[i]++
			if len(found) > 3 {
				return false
			}
		}

		if found[0] >= 2 {
			return true
		}
		for c, f := range found {
			if c == 0 {
				continue
			}
			if f+found[0] >= 2 {
				pairs++
				if found[0] > 0 {
					found[0]--
				}
			}
		}

		if pairs == 2 {
			return true
		}
		return false
	}

	OnePair := func(h Hand) bool {
		found := make(map[int]int, 4)
		pairs := 0
		for _, i := range h.cards {
			found[i]++
		}
		if found[0] == 1 {
			return true
		}
		for _, f := range found {
			if f == 2 {
				pairs++
			}
		}
		if pairs == 1 {
			return true
		}
		return false
	}

	HighCard := func(h Hand) bool { return true }

	// hand types sorted by value, highest first
	var HandTypeRank []HandRankCheck = []HandRankCheck{
		FiveOfAKind,
		FourOfAKind,
		FullHouse,
		ThreeOfAKind,
		TwoPair,
		OnePair,
		HighCard,
	}

	var plays []Hand

	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("opening file %s, %v", path, err)
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.Fields(s.Text())
		b, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatalf("converting %s to int: %v", line[1], err)
		}
		plays = append(plays, NewHand(CardValue, HandTypeRank, line[0], b))
	}
	// 	for i := range plays {
	// 		fmt.Printf("rank=%d bid=%d cards=%s\n", plays[i].rank, plays[i].bid, plays[i].cardsStr)
	// 	}

	sort.Sort(Hands(plays))
	result := 0
	for i := range plays {
		result += (i + 1) * plays[i].bid
		// fmt.Printf("rank=%d bid=%d\n", plays[i].rank, plays[i].bid)
	}
	return result
}
