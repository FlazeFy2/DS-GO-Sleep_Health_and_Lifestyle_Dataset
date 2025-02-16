package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// Descriptive Statistic - Find Mean, Modus, Median
func findMean(values []float64) float64 {
	var sum float64
	for _, value := range values {
		sum += value
	}

	return sum / float64(len(values))
}
func findMode(values []float64) []float64 {
	frequencyMap := make(map[float64]int)
	var maxCount int

	for _, value := range values {
		frequencyMap[value]++
		if frequencyMap[value] > maxCount {
			maxCount = frequencyMap[value]
		}
	}

	var modes []float64
	for value, count := range frequencyMap {
		if count == maxCount {
			modes = append(modes, value)
		}
	}

	return modes
}
func findMedian(values []float64) float64 {
	sort.Float64s(values)
	n := len(values)
	if n%2 == 0 {
		return (values[n/2-1] + values[n/2]) / 2
	}
	return values[n/2]
}

func main() {
	// Read CSV
	file, err := os.Open("Sleep_health_and_lifestyle_dataset.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	headers := records[0]

	// Descriptive Statistic - Find Mean, Modus, Median
	target_col := []string{"Age", "Sleep Duration", "Quality of Sleep", "Physical Activity Level", "Stress Level", "Heart Rate", "Daily Steps"}
	for _, colName := range target_col {
		columnIndex := -1
		for i, header := range headers {
			if header == colName {
				columnIndex = i
				break
			}
		}

		if columnIndex == -1 {
			fmt.Printf("Column '%s' not found in the CSV file\n", colName)
			continue
		}

		var columnValues []float64
		for _, record := range records[1:] {
			if len(record) > columnIndex {
				value := record[columnIndex]
				if v, err := strconv.ParseFloat(value, 64); err == nil {
					columnValues = append(columnValues, v)
				}
			}
		}

		fmt.Printf("\nFor column '%s'\n", colName)
		mean := findMean(columnValues)
		fmt.Printf("Mean : %.2f\n", mean)
		mode := findMode(columnValues)
		fmt.Printf("Mode : %.2f\n", mode)
		median := findMedian(columnValues)
		fmt.Printf("Median : %.2f\n", median)
	}
}
