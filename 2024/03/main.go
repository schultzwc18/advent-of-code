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

func getEnabledMemory(line string) []string {
	re := regexp.MustCompile(`do\(\).*?don\'t\(\)`)
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

func consolidateInput(input []string) string {
	var inputString strings.Builder
	for _, memoryString := range input {
		inputString.WriteString(memoryString)
	}

	return inputString.String()
}

func scanMemory(memory []string) int {
	inputString := consolidateInput(memory)
	instructions := getInstructions(inputString)
	value := calculateResult(instructions)

	return value

}

func scanEnabledMemory(memory []string) int {
	var inputString strings.Builder

	inputString.WriteString("do()")
	inputString.WriteString(consolidateInput(memory))
	inputString.WriteString("don't()")

	enabledMemory := getEnabledMemory(inputString.String())
	enabledMemoryString := consolidateInput(enabledMemory)
	instructions := getInstructions(enabledMemoryString)
	value := calculateResult(instructions)

	return value
}

func main() {
	lines, err := readFile("input/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	value := scanMemory(lines)
	enabledValue := scanEnabledMemory(lines)

	log.Printf("Value: %v", value)
	log.Printf("Enabled Value: %v", enabledValue)
}
