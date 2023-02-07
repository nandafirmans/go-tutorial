package goroutine

import (
	"fmt"
	"strconv"
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

func TestBufferedChannel(t *testing.T) {
	// argumen kedua adalah jumlah buffer (buffer capacity) yang tersedia pada channel.
	// buffered channel ini sangat cocok untuk case yg apabila goroutine penerima lebih lambat-
	// dari pada si pengirim. (jadi data yg dikirim/diterima akan diantrikan sesuai jumlah buffernya).
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		// ketika channel memiliki buffer maka tidak akan terjadi blocking-
		// saat kita mengirim value pada channel walaupun tidak ada yg menerima value channel-nya
		channel <- "Nanda"
		channel <- "Firmansyah"
		channel <- "Afif"
	}()

	go func() {
		// begitu juga saat menerima data, di channel tidak akan terjadi deadlock walaupun-
		// tidak ada data yang dikirim. tapi ini tergantung kapasitas buffer-
		// apabila kapasitas buffer 3 namun kita menerima data ke 4 dan posisinya tidak ada data yg dikirim-
		// kedalam channel maka akan terjadi deadlock saat mengambil data ke 4
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)

		// code dibawah akan menyebabkan deadlock karena data yg dikirim hanya 3 sedangkan kita
		// coba menerima data ke 4
		// fmt.Println(<-channel)
	}()

	fmt.Println("Done Executing Buffered Channel")
	time.Sleep(2 * time.Second)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 1; i <= 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Done Executing TestRangeChannel")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Channel 1", data)
			counter++

		case data := <-channel2:
			fmt.Println("Channel 2", data)
			counter++

		// kalau tidak ada default select maka-
		// akan terjadi deadlock error karena data-
		// diambil tapi belum ada yg dikirim.
		// jadi akan terus looping tanpa harus deadlock sampai kedua datanya diterima
		default:
			fmt.Println("Menuggu Data")
		}

		if counter == 2 {
			break
		}
	}

}
