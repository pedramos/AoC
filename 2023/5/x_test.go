package main

import "testing"

func TestPart1(t *testing.T) {
	r := Part1("t1.txt")
	if r != 35 {
		t.Errorf("Expected 35, got %d\n", r)
	}
}

func TestPart2(t *testing.T) {
	r := Part2("t1.txt")
	if r != 46 {
		t.Errorf("Expected 46, got %d\n", r)
	}
}
