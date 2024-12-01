package main

import (
	"fmt"
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

func main() {
	lines, err := readFile("input/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	listOne, listTwo := getLists(lines)
	totalDistance := 0

	for index, value := range listOne {
		totalDistance += int(math.Abs(float64(listTwo[index] - value)))
	}

	fmt.Println(totalDistance)
}
