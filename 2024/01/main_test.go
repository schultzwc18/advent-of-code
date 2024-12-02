package main

import (
	"log"
	"testing"
)

func TestTotalDistance(t *testing.T) {
	want := 2375403
	lines, err := readFile("input/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	listOne, listTwo := getLists(lines)
	totalDistance := getTotalDistance(listOne, listTwo)

	if totalDistance != want {
		t.Errorf("Got: %v; Want: %v", totalDistance, want)
	}
}

func TestSimilarityScore(t *testing.T) {
	want := 23082277
	lines, err := readFile("input/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	listOne, listTwo := getLists(lines)
	similarityScore := getSimilarityScore(listOne, listTwo)

	if similarityScore != want {
		t.Errorf("Got: %v; Want: %v", similarityScore, want)
	}
}
