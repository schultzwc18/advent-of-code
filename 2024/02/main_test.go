package main

import (
	"log"
	"testing"
)

func TestSafeReports(t *testing.T) {
	want := 421
	lines, err := readFile("input/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	safeReports, _ := getSafeReports(lines)

	if safeReports != want {
		t.Errorf("Got: %v; Want: %v", safeReports, want)
	}
}

func TestSafeToleratedReports(t *testing.T) {
	want := 476
	lines, err := readFile("input/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	_, safeReports := getSafeReports(lines)

	if safeReports != want {
		t.Errorf("Got: %v; Want: %v", safeReports, want)
	}
}
