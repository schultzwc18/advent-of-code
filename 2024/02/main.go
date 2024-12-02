package main

import (
	"log"
	"math"
	"os"
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

func checkReportSafety(report []string) bool {
	isAscending := false
	isReportSafe := true

	// handles any empty lines within input file
	if len(report) < 2 {
		return false
	}

	for index, value := range report {
		if index+1 >= len(report) {
			break
		}

		valueInt, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}

		nextValueInt, err := strconv.Atoi(report[index+1])
		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}

		valueDifference := int(math.Abs(float64(nextValueInt - valueInt)))

		if index == 0 && valueInt < nextValueInt {
			isAscending = true
		}

		if valueDifference < 1 || valueDifference > 3 {
			isReportSafe = false
		}

		if isAscending && valueInt > nextValueInt {
			isReportSafe = false
		} else if !isAscending && valueInt < nextValueInt {
			isReportSafe = false
		}
	}

	return isReportSafe
}

func getSafeReports(reports []string) int {
	safeReportCount := 0

	for _, report := range reports {

		reportSplit := strings.Fields(report)
		reportSafety := checkReportSafety(reportSplit)

		if reportSafety {
			safeReportCount++
		}
	}

	return safeReportCount
}

func main() {
	lines, err := readFile("input/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	safeReports := getSafeReports(lines)

	log.Printf("Safe Reports: %v", safeReports)
}
