package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsyncronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	// wait group digunakan untuk menjalankan proses asyncronous
	// dengan cara menambahkan group.Add(number int) dan di set group.Done() untuk mengetahui bahwa goroutine
	// sudah selesai di esekusi
	// sehingga kita tidak pelu melakukan time.sleep untuk menunggu proses goroutine selesai
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsyncronous(group)
	}

	group.Wait()
}
