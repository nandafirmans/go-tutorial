package main

import "fmt"

type Person struct {
	Name string
}

// kalau tidak menggunakan tipe data pointer
// maka struct function (method) dibawah akan mempasing value baru
// sehingga ketika kita tidak dapat memutasi datanya langsung

// disarankan selalu menggunakan Pointer untuk struct function (method)
// karena kita dapat langsung memutasi data didalamnya
// dan selain itu lebih hemat memory karena kita tidak membuat value baru setiap method ini diesekusi
func (person *Person) Married() {
	person.Name = "Mr. " + person.Name
}

func main() {
	person1 := Person{"Afif"}
	person1.Married()

	fmt.Println(person1)
}
