package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
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
func findMax(values []float64) float64 {
	sort.Float64s(values)
	return values[len(values)-1]
}
func findMin(values []float64) float64 {
	sort.Float64s(values)
	return values[0]
}
func findVariance(values []float64) float64 {
	mean := findMean(values)
	var variance float64
	var count int
	for _, value := range values {
		if !math.IsNaN(value) {
			variance += (value - mean) * (value - mean)
			count++
		}
	}

	if count > 1 {
		variance /= float64(count - 1)
	}

	return variance
}
func findStandardDeviance(val float64) float64 {
	val = math.Sqrt(val)
	shift := math.Pow(10, float64(2))
	return math.Round(val*shift) / shift
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
		max := findMax(columnValues)
		fmt.Printf("Max : %.2f\n", max)
		min := findMin(columnValues)
		fmt.Printf("Min : %.2f\n", min)
		rangenum := max - min
		fmt.Printf("Range : %.2f\n", rangenum)
		mean := findMean(columnValues)
		fmt.Printf("Mean : %.2f\n", mean)
		mode := findMode(columnValues)
		fmt.Printf("Mode : %.2f\n", mode)
		median := findMedian(columnValues)
		fmt.Printf("Median : %.2f\n", median)
		variance := findVariance(columnValues)
		fmt.Printf("Variance : %.2f\n", variance)
		std := findStandardDeviance(variance)
		fmt.Printf("Standard Deviance : %.2f\n", std)
	}
}
