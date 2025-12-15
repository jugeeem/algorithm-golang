package pkg

import "math/rand"

// GenerateRandIntArray はランダムな整数配列を生成して返します。
func GenerateRandIntArray() []int {
	arraySize := 100
	maxRange := 100000

	// ランダムな整数配列を生成
	array := make([]int, arraySize)
	for i := range arraySize {
		array[i] = rand.Intn(maxRange)
	}

	return array
}
