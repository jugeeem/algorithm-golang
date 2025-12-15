package main

import (
	"fmt"

	"github.com/jugeeem/algorithm-golang/src/pkg"
)

func bubbleSort(numbers []int) []int {
	n := len(numbers)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	return numbers
}

func execute() {
	numbers := pkg.GenerateRandIntArray()
	sortedNumbers := bubbleSort(numbers)
	fmt.Println("Sorted Numbers:", sortedNumbers)
}

func main() {
	pkg.MeasureWithWait("Bubble Sort", execute)
}
