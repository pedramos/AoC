package main

import "testing"

func TestPart1(t *testing.T) {
	r := Part1("t1.txt")
	if r != 114 {
		t.Errorf("Expected 114, got %d\n", r)
	}

}

func TestPart2(t *testing.T) {
	r := Part2("t1.txt")
	if r != 2 {
		t.Errorf("Expected 2, got %d\n", r)
	}
}
