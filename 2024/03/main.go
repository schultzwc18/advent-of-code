package main

import (
	"log"
	"os"
	"regexp"
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

func getInstructions(line string) []string {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	instructions := re.FindAllString(line, -1)

	return instructions
}

func calculateResult(instructions []string) int {
	value := 0
	prefix := "mul("
	suffix := ")"
	delimiter := ","

	for _, instruction := range instructions {
		numberPair := strings.Replace(strings.Replace(instruction, prefix, "", -1), suffix, "", -1)
		numberList := strings.Split(numberPair, delimiter)

		numOne, err := strconv.Atoi(numberList[0])
		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}

		numTwo, err := strconv.Atoi(numberList[1])
		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}

		value += numOne * numTwo
	}

	return value
}

func scanMemory(memory []string) int {
	var instructions []string
	for _, memoryString := range memory {
		currentInstructions := getInstructions(memoryString)
		instructions = append(instructions, currentInstructions...)
	}

	value := calculateResult(instructions)

	return value

}

func main() {
	lines, err := readFile("input/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	value := scanMemory(lines)

	log.Printf("Value: %v", value)
}
