package main

import (
	"log"
	"testing"
)

func TestMemory(t *testing.T) {
	want := 184576302
	lines, err := readFile("input/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	value := scanMemory(lines)

	if value != want {
		t.Errorf("Got: %v; Want: %v", value, want)
	}
}

func TestSafeToleratedReports(t *testing.T) {
	want := 118173507
	lines, err := readFile("input/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	value := scanEnabledMemory(lines)

	if value != want {
		t.Errorf("Got: %v; Want: %v", value, want)
	}
}
