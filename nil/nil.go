package main

import "fmt"

// Nil adalah data kosong
// Biasanya pada bahasa pemrograman lain object yg belum di inisialisasi maka otomatis nilainya adalah null/nil.
// Berbeda dengan Go saat kita membuat tipe data tertentu (int, string, dll) maka akan dibuatkan default value tertentu
// Nil sendiri hanya dapat digunakan di beberapa tipe data seperti interface, function, map, slice, pointer & channel

func main() {
	var foo string
	fmt.Println(foo)
	// fmt.Println(foo == nil) // <== error

	var obj map[string]string
	fmt.Println(obj)
	fmt.Println(obj == nil)
}
