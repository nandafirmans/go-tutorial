package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	// ticker ini mirip setInterval di javascript
	// channel akan dikirim berdasarkan waktu yang ditentukan (1 detik dalam fungsi ini)

	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for time := range ticker.C {
		fmt.Println(time)
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
}
