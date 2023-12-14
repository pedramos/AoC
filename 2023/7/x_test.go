package main

import "testing"

func TestPart1(t *testing.T) {
	r := Part1("t1.txt")
	if r != 6440 {
		t.Errorf("Expected 6440, got %d\n", r)
	}
}

func TestPart2(t *testing.T) {
	r := Part2("t1.txt")
	if r != 5905 {
		t.Errorf("Expected 5905, got %d\n", r)
	}
}
