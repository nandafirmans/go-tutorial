package goroutine

import (
	"fmt"
	"testing"
	"time"
)

// untuk semua programing langguage yang memiliki fitur parallel programing-
// atau concurrency sudah pasti masalah race condition ini ada.

// race condition adalah kondisi dimana ada banyak goroutine yang memutasi-
// shared variable disaat yang bersamaan.
func TestRaceCondition(t *testing.T) {
	number := 0

	for i := 0; i < 1000; i++ {
		// karena ada kemungkinan dari 1000 go routine ini jalan bersamaan sehingga
		// value akhir dari variable "number" bukan 10000
		go func() {
			for j := 0; j < 100; j++ {
				number += 1
			}
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Number: ", number)
}
