package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		// once digunakan untuk menjalankan fungsi hanya sekali saja
		// walaupun di loop ia akan tetap running 1 kali
		once.Do(OnlyOnce)
		group.Done()
	}

	group.Wait()
	fmt.Println("Counter", counter)

}
