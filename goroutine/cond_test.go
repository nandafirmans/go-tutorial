package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Cond adalah locking dengan kondisi, jadi setelah goroutine lock
// maka akan menunggu ".Wait()" apakan akan lanjut esekusi atau tidak
// untuk melanjutkan maka bisa mengirim ".Signal()" untuk lanjut esekusi satu/satu goroutine nya
// dan ".Broadcast()" berarti tidak perlu wait lagi untuk tiap lockingnya

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	//go func() {
	//time.Sleep(1 * time.Second)
	//cond.Broadcast()
	//}()

	group.Wait()
}
