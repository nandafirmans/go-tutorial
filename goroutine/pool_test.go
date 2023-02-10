package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	// Pool adalah implementasi design pattern object pool pattern.
	// sederhananya design pattern pool ini digunakan untuk menyimpan data.
	// selanjutnya untuk menggunakan datanya kita dapat mengambilnya dari pool-
	// dan setelah selesai menggunakan datanya kita dapat menyimpan kembali Poolnya.

	pool := sync.Pool{}
	group := sync.WaitGroup{}

	pool.Put("Nanda")
	pool.Put("Firmansyah")
	pool.Put("Afif")

	for i := 0; i < 10; i++ {

		go func() {
			group.Add(1)
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
			group.Done()
		}()
	}

	group.Wait()
}
