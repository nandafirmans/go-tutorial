package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// secara default di golang semua variable itu di passing by value, bukan by reference
func samplePassByValue() {
	address1 := Address{"Tangerang", "Banten", "Indonesia"}
	address2 := address1

	address2.City = "Serang"

	fmt.Println(address1)
	fmt.Println(address2)

	// address1 {Tangerang Banten Indonesia}
	// address2 {Serang Banten Indonesia}
	// hasilnya address1 datanya tidak berubah padahal data city di address2 diubah
	// ini karena ketika membuat address2 value dari address1 di copy baru bukan di reference
}

func samplePassByReference() {
	var address1 Address = Address{"Tangerang", "Banten", "Indonesia"}
	// address2 menunjuk ke address1 (pointer)
	var address2 *Address = &address1

	address2.City = "Serang"

	// {Serang Banten Indonesia}                                                                                                                                                  // &{Serang Banten Indonesia}
	// ketika address2 diubah maka address1 juga akan berubah valuenya
	// karena value address2 me-reference ke address1 (address2 adalah pointer address1)
	fmt.Println(address1)
	fmt.Println(address2)

	// ketikan kita mengubah address2 value baru seperti contoh dibawah
	// maka address1 tidak akan berubah valuenya, melainkan address2 akan menunjuk ke value yg baru
	address2 = &Address{"Malang", "Jawa Timur", "Indonesia"}

	// untuk memaksa agar pointer variable merubah value aslinya (dan semua yg me-reference padanya)
	// kita dapat menggunakan operator * contoh:
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	// membuat pointer baru dengan data kosong, dapat menggunakan funsi "new"
	var address3 *Address = new(Address)
	address3.City = "Jakarta"

	fmt.Println(address3)
}

func _ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func pointerInFunc() {
	alamat := Address{
		City:     "Jakarta Barat",
		Province: "Jakarta",
	}
	var alamatPointer *Address = &alamat

	_ChangeCountryToIndonesia(alamatPointer)

	fmt.Println(alamat)
}

func main() {
	samplePassByValue()
	fmt.Println("----")
	samplePassByReference()

	fmt.Println("")
	fmt.Println("----")
	fmt.Println("")

	pointerInFunc()

	fmt.Println("")
	fmt.Println("----")
	fmt.Println("")

}
