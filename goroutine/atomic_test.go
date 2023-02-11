package goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	// atomic adalah sebuah package yg digunakan untuk menggunakan data primitive secara aman pada proses concurrent
	// sehingga kita tidak pelu melakukan proses locking

	group := sync.WaitGroup{}
	var counter int64 = 0

	for i := 0; i < 100; i++ {
		go func() {
			defer group.Done()
			group.Add(1)

			for j := 0; j < 100; j++ {
				// tidak perlu pakai Mutex
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	group.Wait()
	fmt.Println("Counter ", counter)

}
