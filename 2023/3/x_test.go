package main

import "testing"

func TestPart1(t *testing.T) {
	r := Part1("t1.txt")
	if r != 4361 {
		t.Errorf("Expected 4361, got %d\n", r)
	}
}

func TestPart2(t *testing.T) {
	r := Part2("t1.txt")
	if r != 467835 {
		t.Errorf("Expected 4361, got %d\n", r)
	}
}
