package goroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGoMaxProcs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU ", totalCpu)

	// cara mengubah total threads:
	// runtime.GOMAXPROCS(20)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread ", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine ", totalGoroutine)
}
