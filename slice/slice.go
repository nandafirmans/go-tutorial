package main

import "fmt"

func main() {
	// slice potongan array
	months := [...]string{"Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agt", "Sep", "Oct", "Nov", "Dec"}
	slice1 := months[5:8]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// akan merubah data slice
	// months[5] = "Jun Diubah"
	// fmt.Println(slice1)

	// akan merubah data diarray aslinya
	// slice1[1] = "Jul Diubah"
	// fmt.Println(months)

	slice2 := months[10:]
	fmt.Println(slice2)

	// append kalau kapasitasnya gak muat maka akan buat array baru tp kalau muat akan merefer ke array aslinya
	slice3 := append(slice2, "New Month")
	fmt.Println(slice3)
	slice3[1] = "DESEMBERRRR"
	fmt.Println(slice3)
	fmt.Println(slice2)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Nanda"
	newSlice[1] = "Firmansyah"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
