package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(3 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

func TestTimerAfter(t *testing.T) {
	// bedanya langsung return channelnya
	channel := time.After(3 * time.Second)
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)
}

func TestTimerAfterFunc(t *testing.T) {
	// kalau butuh fungsinya saja bisa gunakan AfterFunc, mirip2 setTimeout di JS
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(3*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})

	fmt.Println(time.Now())
	group.Wait()
}
