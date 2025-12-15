package pkg

import (
	"fmt"
	"sync"
	"time"
)

// Measure は関数の実行時間を計測します（同期版）
func Measure(name string, f func()) {
	start := time.Now()
	f()
	end := time.Now()
	fmt.Printf("%s execution time: %v\n", name, end.Sub(start))
}

// MeasureAsync は関数の実行時間をgoroutineで非同期計測します
// 完了を待つにはsync.WaitGroupなどを外部で管理してください
func MeasureAsync(name string, f func(), wg *sync.WaitGroup) {
	if wg != nil {
		wg.Add(1)
	}

	go func() {
		if wg != nil {
			defer wg.Done()
		}
		start := time.Now()
		f()
		end := time.Now()
		fmt.Printf("%s execution time: %v\n", name, end.Sub(start))
	}()
}

// MeasureWithWait は関数の実行時間をgoroutineで計測し、完了を待ちます
func MeasureWithWait(name string, f func()) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		start := time.Now()
		f()
		end := time.Now()
		fmt.Printf("%s execution time: %v\n", name, end.Sub(start))
	}()

	wg.Wait()
}

// MeasureConcurrent は複数の関数を並行実行し、それぞれの実行時間を計測します
func MeasureConcurrent(tasks map[string]func()) {
	var wg sync.WaitGroup

	for name, f := range tasks {
		wg.Add(1)
		go func(taskName string, taskFunc func()) {
			defer wg.Done()
			start := time.Now()
			taskFunc()
			end := time.Now()
			fmt.Printf("%s execution time: %v\n", taskName, end.Sub(start))
		}(name, f)
	}

	wg.Wait()
}
