package main

import "testing"

func TestPart1(t *testing.T) {
	r := Part1("t1.txt")
	if r != 2 {
		t.Errorf("Expected 2, got %d\n", r)
	}
	r = Part1("t2.txt")
	if r != 6 {
		t.Errorf("Expected 6, got %d\n", r)
	}
	r = Part1("p1.txt")
	if r != 11309 {
		t.Errorf("Expected 11309, got %d\n", r)
	}

}

func TestPart2(t *testing.T) {
	r := Part2("t1.txt")
	if r != 2 {
		t.Errorf("Expected 2, got %d\n", r)
	}
	r = Part2("t2.txt")
	if r != 6 {
		t.Errorf("Expected 6, got %d\n", r)
	}
	r = Part2("t3.txt")
	if r != 6 {
		t.Errorf("Expected 6, got %d\n", r)
	}
}
