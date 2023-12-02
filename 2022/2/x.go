package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
Rock 	// A // X // 1
Paper 	// B/// Y // 2
Scissors	// C // Z // 3
*/

func main() {

	total1 := 0
	total2 := 0

	f, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	iPlayed := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	resultPoints := map[string]int{
		"A X": 3,
		"B Y": 3,
		"C Z": 3,
		"C X": 6,
		"A Y": 6,
		"B Z": 6,
	}

	// for part2
	gestureToPlayDecode := map[string]map[string]string{
		"A": map[string]string{
			"X": "Z",
			"Y": "X",
			"Z": "Y",
		},
		"B": map[string]string{
			"X": "X",
			"Y": "Y",
			"Z": "Z",
		},
		"C": map[string]string{
			"X": "Y",
			"Y": "Z",
			"Z": "X",
		},
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// Part1
		total1 = total1 + resultPoints[scanner.Text()] + iPlayed[strings.Fields(scanner.Text())[1]]
		// fmt.Printf("%s: %d\n", scanner.Text(), resultPoints[scanner.Text()] + iPlayed[strings.Fields(scanner.Text())[1]])
		
		// Part2
		oponent := strings.Fields(scanner.Text())[0]
		result := strings.Fields(scanner.Text())[1]	
	
		played := gestureToPlayDecode[oponent][result]
		game := fmt.Sprintf("%s %s", oponent, played)

		total2 = total2 + resultPoints[game] + iPlayed[played]
		// fmt.Printf("%s %s: %s %d\n", oponent, result, played, resultPoints[game] + iPlayed[played])
	}

	fmt.Printf("Part 1: %d\n", total1)
	fmt.Printf("Part 2: %d\n", total2)

}
