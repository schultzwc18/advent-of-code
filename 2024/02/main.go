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

func concatArray(indexToSkip int, reportArray []string) []string {
	var revisedArray []string

	for index, _ := range reportArray {
		if indexToSkip != index {
			revisedArray = append(revisedArray, reportArray[index])
		}

	}

	return revisedArray
}

func checkReportSafety(report []string, includeToleration bool) bool {
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

	if includeToleration && !isReportSafe {
		for index, _ := range report {
			tolerationReport := concatArray(index, report)
			tolerationReportSafety := checkReportSafety(tolerationReport, false)
			tolerationReport = nil

			if tolerationReportSafety {
				isReportSafe = true
				return isReportSafe
			}
		}
	}

	return isReportSafe
}

func getSafeReports(reports []string) (int, int) {
	safeReportCount := 0
	safeTolerationReportCount := 0

	for _, report := range reports {
		reportSplit := strings.Fields(report)
		reportSafety := checkReportSafety(reportSplit, false)
		reportTolerationSafety := checkReportSafety(reportSplit, true)

		if reportSafety {
			safeReportCount++
		}

		if reportTolerationSafety {
			safeTolerationReportCount++
		}
	}

	return safeReportCount, safeTolerationReportCount
}

func main() {
	lines, err := readFile("input/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	safeReports, safeTolerationReports := getSafeReports(lines)

	log.Printf("Safe Reports: %v", safeReports)
	log.Printf("Safe Toleration Reports: %v", safeTolerationReports)
}
