package main

import (
	"fmt"

	"github.com/jugeeem/algorithm-golang/src/pkg"
)

func cocktailSort(numbers []int) []int {
	lenNumbers := len(numbers)
	swapped := true
	start := 0
	end := lenNumbers - 1

	for swapped {
		swapped = false

		for i := start; i < end; i++ {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}

		swapped = false
		end--

		for i := end - 1; i > start-1; i-- {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				swapped = true
			}
		}

		start++
	}

	return numbers
}

func execute() {
	numbers := pkg.GenerateRandIntArray()
	sortedNumbers := cocktailSort(numbers)
	fmt.Println("Sorted Numbers:", sortedNumbers)
}

func main() {
	pkg.MeasureWithWait("Cocktail Sort", execute)
}
