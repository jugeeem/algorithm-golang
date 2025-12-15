package main

import (
	"fmt"

	"github.com/jugeeem/algorithm-golang/src/pkg"
)

// inOrder は整数配列が昇順に並んでいるかを確認します。
func inOrder(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			return false
		}
	}

	return true
}

// bogoSort はボゴソートアルゴリズムを実装します。
func bogoSort(numbers []int) []int {
	for !inOrder(numbers) {
		numbers = pkg.GenerateRandIntArray()
	}

	return numbers
}

// execute はボゴソートを実行し、結果を表示します。
func execute() {
	numbers := pkg.GenerateRandIntArray()
	sortedNumbers := bogoSort(numbers)

	fmt.Println("Sorted Numbers:", sortedNumbers)
}

func main() {
	pkg.MeasureWithWait("Bogo Sort", execute)
}
