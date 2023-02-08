package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// untuk menghandle kasus race condition maka kita bisa menggunakan Mutex (Mutual Exclution)
// dengan membuat seperti antrian saat sebelum kita memutasi shared variable di goroutine
func TestMutex(t *testing.T) {
	number := 0
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				// hanya satu goroutine yang dapat melakukan lock & unlock-
				// sehingga goroutine yang lain harus menunggu-
				// dengan begitu goroutine secara bergantian memutasi variable "number"
				mutex.Lock()
				number += 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Number: ", number)

}
