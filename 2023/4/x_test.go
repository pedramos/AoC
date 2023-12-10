package main

import "testing"

func TestPart1(t *testing.T) {
	r := Part1("t1.txt")
	if r != 13 {
		t.Errorf("Expected 13, got %d\n", r)
	}
}

func TestPart2(t *testing.T) {
	r := Part2("t1.txt")
	if r != 30 {
		t.Errorf("Expected 30, got %d\n", r)
	}
}
