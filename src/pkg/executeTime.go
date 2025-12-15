package pkg

import (
	"fmt"
	"time"
)

func Measure(name string, f func()) {
	start := time.Now()
	f()
	end := time.Now()
	fmt.Printf("%s execution time: %v\n", name, end.Sub(start))
}
