package main

import (
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFile(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(content), "\n"), nil
}

func getLists(lines []string) ([]int, []int) {
	var listOne []int
	var listTwo []int

	for _, line := range lines {
		lineSplit := strings.Fields(line)

		if len(lineSplit) != 2 {
			log.Printf("Line %v is not length 2", line)
			continue
		}

		numOne, err := strconv.Atoi(lineSplit[0])
		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}
		listOne = append(listOne, numOne)

		numTwo, err := strconv.Atoi(lineSplit[1])
		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}
		listTwo = append(listTwo, numTwo)
	}

	sort.Ints(listOne)
	sort.Ints(listTwo)

	return listOne, listTwo
}

func getTotalDistance(listOne []int, listTwo []int) int {
	totalDistance := 0

	for index, value := range listOne {
		totalDistance += int(math.Abs(float64(listTwo[index] - value)))
	}

	return totalDistance
}

func getTotalOccurancesScore(num int, list []int) int {
	occurances := 0

	for _, value := range list {
		if value == num {
			occurances++
		}
	}

	return num * occurances
}

func getSimilarityScore(listOne []int, listTwo []int) int {
	similarityScore := 0
	similarityOccurances := make(map[int]int)

	for _, value := range listOne {
		val, ok := similarityOccurances[value]
		if ok {
			similarityScore += val
		} else {
			similarityOccurances[value] = getTotalOccurancesScore(value, listTwo)
			similarityScore += similarityOccurances[value]
		}
	}

	return similarityScore
}

func main() {
	lines, err := readFile("input/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	listOne, listTwo := getLists(lines)
	totalDistance := getTotalDistance(listOne, listTwo)
	similarityScore := getSimilarityScore(listOne, listTwo)

	log.Printf("Total distance: %v", totalDistance)
	log.Printf("Similarity score: %v", similarityScore)
}
