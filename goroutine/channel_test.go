package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(3 * time.Second)

		// channel akan memblock proses apabila tidak ada yg-
		// mengambil valuenya.
		// dan sebaliknya apabila tidak ada data yg dikirim-
		// ke channel tapi ada yg mengambil maka akan terjadi deadlock.
		channel <- "Hello World"

		fmt.Println("Selesai kirim data")
	}()

	data := <-channel
	fmt.Println(data)
}

// by default go lang selalu pass by value
// tapi ada pengecualian untuk tipe data channel
// sehingga kita tidak perlu membuat pointer channel
// khusus channel otomatis pass by reference
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Herro Warudo"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	result := <-channel
	fmt.Println(result)
}

// secara default pada channel kita bisa mengirim atau mengambil data secara bebas
// tapi dengan tanda seperti ini "chan<-" (in) menandakan kita hanya bisa mengirim-
// value di dalam fungsi ini, apabila kita melakukan pengambilan value maka akan error
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello World"

	// error, karena tidak boleh mengambil data
	// fmt.Println(<-channel)
}

// sebaliknya tanda ini "<-chan" (out) menandakan kita hanya dapat mengambil value
// pada fungsi ini dan apabila kita mengirim value maka akan terjadi error
func OnlyOut(channel <-chan string) {
	result := <-channel
	fmt.Println(result)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
}
